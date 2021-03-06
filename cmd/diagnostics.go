package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var diagnosticsCmd = &cobra.Command{
	Use:   "diagnostics",
	Short: "produce a report of the state of PCF",
	Run: func(cmd *cobra.Command, args []string) {
		client := setupOpsmanClient()
		report, err := client.Get("/api/v0/diagnostic_report", 10*time.Minute)
		if err != nil {
			rp.Fail(err)
		}
		rp.PrintReport(string(report))
	},
}
