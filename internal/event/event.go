package event

import (
	"fmt"

	"github.com/steugen/trigger/internal/config"
)
// events:
// pod_restart
// cron 
// resource_consumption 
// crashloopbackoff
// 
// Placeholder for now
func Start(cfg *config.Config) error {
	fmt.Println("Event listener started with configuration:", cfg)
	return nil
}
