package targets

import (
	"chainflow-vitwit/config"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"os/exec"
	"regexp"
)

func GaiadVersion(_ HTTPOptions, cfg *config.Config, c client.Client) {
	cmd := exec.Command("gaiad", "version", "--long")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		return
	}

	resp := string(out)

	r := regexp.MustCompile(`version: ([0-9]{1}.[0-9]{1}.[0-9]{1})`)
	matches := r.FindAllStringSubmatch(resp, -1)
	if len(matches) == 0 {
		return
	}
	if len(matches[0]) != 2 {
		return
	}
	log.Printf("Version: %s", matches[0][1])
}
