package targets

import (
	"cosmos-validator-mission-control/config"
	"encoding/json"
	"log"

	client "github.com/influxdata/influxdb1-client/v2"
)

// GetSelfDelegation to get self delegation of a validator
func GetSelfDelegation(ops HTTPOptions, cfg *config.Config, c client.Client) {
	bp, err := createBatchPoints(cfg.InfluxDB.Database)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	var selfDelegationBal string
	var balanceDeamon string

	if cfg.DaemonName == "akashd" {
		var delegationResp AkashSelfDelegationBalance
		err = json.Unmarshal(resp.Body, &delegationResp)
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		selfDelegationBal = delegationResp.Result.Balance.Amount
		balanceDeamon = delegationResp.Result.Balance.Denom

	} else {

		var delegationResp SelfDelegation
		err = json.Unmarshal(resp.Body, &delegationResp)
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		selfDelegationBal = delegationResp.Result.Balance
		// balanceDeamon = "umuon"
	}

	if balanceDeamon == "" {
		balanceDeamon = cfg.StakingDemon
	}

	addressBalance := convertToCommaSeparated(selfDelegationBal) + balanceDeamon
	_ = writeToInfluxDb(c, bp, "vcf_self_delegation_balance", map[string]string{}, map[string]interface{}{"balance": addressBalance})
	log.Printf("Address Balance: %s", addressBalance)
}
