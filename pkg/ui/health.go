package ui

import (
	"fmt"
	"time"
	"unipass/pkg/config"
)

func ProcessHealthCheckStatus(lastCheck time.Time) {
	if config.EvaluateWeeklyHealthCheck(lastCheck) {
		fmt.Println("[DIAGNOSTIC] Weekly system health evaluation required.")
	} else {
		fmt.Println("[DIAGNOSTIC] System health status optimal.")
	}
}
