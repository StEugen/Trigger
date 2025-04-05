package event

import (
	"fmt"

	"github.com/steugen/trigger/internal/config"
)

// Placeholder for now
func Start(cfg *config.Config) error {
	fmt.Println("Event listener started with configuration:", cfg)
	return nil
}
