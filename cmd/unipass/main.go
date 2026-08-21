package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"unipass/pkg/config"
	"unipass/pkg/health"
	"unipass/pkg/locationiq"
	"unipass/pkg/ui"
)

var version = "1.0.0"

func main() {
	// Load environment variables
	_ = config.LoadEnv(".env")

	healthFlag := flag.Bool("health", false, "Run system-wide environment diagnostics")
	jsonFlag := flag.Bool("json", false, "Output health diagnostics as raw JSON")
	flag.Parse()

	apiKey := os.Getenv("LOCATIONIQ_API_KEY")
	var locClient *locationiq.Client
	if apiKey != "" {
		locClient = locationiq.NewClient(apiKey)
		// Location Client is ready to use
		_ = locClient
	}

	if *healthFlag {
		hc := health.NewHealthController()
		report := hc.Diagnose()

		if *jsonFlag {
			data, err := json.MarshalIndent(report, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error serializing health report: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(data))
			return
		}

		ui.RenderHeader("UniPass Health Diagnosis", version)
		fmt.Printf("  • Operating System : %s (%s)\n", report.OS, report.Arch)
		fmt.Printf("  • Environment      : %s\n", report.Environment)
		fmt.Printf("  • Interactive TTY  : %t\n", report.HasTTY)
		fmt.Printf("  • Color Support    : %t\n", report.HasColorSupport)
		fmt.Printf("  • Location/GPS     : %t\n", report.HasGPS)
		fmt.Println("================================================================================")
		return
	}

	ui.RenderHeader("UniPass CLI", version)
	fmt.Println("🚀 UniPass system initialized and ready.")
}
