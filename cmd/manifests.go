package cmd

import (
	"encoding/json"

	"github.com/pivotal-cloudops/omen/internal/manifest"
	"github.com/pivotal-cloudops/omen/internal/tile"
	"github.com/pivotal-cloudops/omen/internal/userio"
	"github.com/spf13/cobra"
)

var manifestsCmd = &cobra.Command{
	Use:   "manifests",
	Short: "get the manifests of all deployments and cloud-config",
	Run: func(cmd *cobra.Command, args []string) {
		client := getOpsmanClient()
		tileLoader := tile.NewTilesLoader(client)
		manifestLoader := manifest.NewManifestsLoader(client, tileLoader)

		manifests, err := manifestLoader.LoadDeployed()
		if err != nil {
			userio.PrintReport("", err)
		}

		bytes, err := json.MarshalIndent(manifests, "", " ")
		if err != nil {
			userio.PrintReport("", err)
		}

		userio.PrintReport(string(bytes), nil)
	},
}
