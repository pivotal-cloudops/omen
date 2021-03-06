package cmd

import (
	"encoding/json"

	"github.com/pivotal-cloudops/omen/internal/manifest"
	"github.com/pivotal-cloudops/omen/internal/tile"
	"github.com/spf13/cobra"
)

var manifestsCmd = &cobra.Command{
	Use:   "manifests",
	Short: "get the manifests of all deployments and cloud-config",
	Run: func(cmd *cobra.Command, args []string) {
		client := setupOpsmanClient()
		tileLoader := tile.NewTilesLoader(client)
		manifestLoader := manifest.NewManifestsLoader(client, tileLoader)

		manifests, err := manifestLoader.LoadAllDeployed()
		if err != nil {
			rp.Fail(err)
		}

		bytes, err := json.MarshalIndent(manifests, "", " ")
		if err != nil {
			rp.Fail(err)
		}

		rp.PrintReport(string(bytes))
	},
}
