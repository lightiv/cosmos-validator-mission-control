package targets

import (
	"chainflow-vitwit/config"
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"os/exec"
)

func CheckGaiad(_ HTTPOptions, cfg *config.Config, c client.Client) {
	bp, err := createBatchPoints(cfg.InfluxDB.Database)
	if err != nil {
		return
	}

	cmd := exec.Command("bash", "-c", "</dev/tcp/0.0.0.0/26656 &>/dev/null")
	out, err := cmd.CombinedOutput()
	if err != nil {
		_ = SendTelegramAlert("Gaiad is not running", cfg)
		_ = SendEmailAlert("Gaiad is not running", cfg)
		_ = writeToInfluxDb(c, bp, "vcf_gaiad_status", map[string]string{}, map[string]interface{}{"status": "down"})
		return
	}

	resp := string(out)
	if resp != "" {
		_ = SendTelegramAlert(fmt.Sprintf("Gaiad is not running: \n%v", resp), cfg)
		_ = SendEmailAlert(fmt.Sprintf("Gaiad is not running: \n%v", resp), cfg)
		_ = writeToInfluxDb(c, bp, "vcf_gaiad_status", map[string]string{}, map[string]interface{}{"status": "down"})
		return
	}

	_ = writeToInfluxDb(c, bp, "vcf_gaiad_status", map[string]string{}, map[string]interface{}{"status": "up"})
}
