package cmd

import (
	"fmt"
	"github.com/pivotal-cloudops/omen/internal/tile"
	"github.com/pivotal-cloudops/omen/internal/tileguid"
	"github.com/spf13/cobra"
	"os"
)

var guidCmd = &cobra.Command{
	Use:   "tile-guid <product slug>",
	Short: "Displays the guid for a product",
	Long:  "Displays the guid for the installed product based on it's slug",
	Run:   tileGuid,
}

func tileGuid(_ *cobra.Command, args []string) {
	client := setupOpsmanClient()
	tileLoader := tile.NewTilesLoader(client)

	guid, err := tileguid.FindGuid(tileLoader, args[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(guid)
}
