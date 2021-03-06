package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pivotal-cloudops/omen/internal/applychanges"
	"github.com/pivotal-cloudops/omen/internal/manifest"
	"github.com/pivotal-cloudops/omen/internal/tile"
	"github.com/spf13/cobra"
)

var nonInteractive bool
var products string
var dryRun bool
var quiet bool

var applyChangesCmd = &cobra.Command{
	Use:   "apply-changes",
	Short: "apply any staged changes",
	Long:  "Produces a diff of staged versus deployed changes and then applies those staged changes",
	Run:   applyChangesFunc,
}

func init() {
	applyChangesCmd.Flags().StringVarP(&products, "products", "P", "",
		`Optional flag to set the products to apply changes for (e.g. "product-1" or "product-1,product-2")`)

	applyChangesCmd.Flags().BoolVarP(&nonInteractive, "non-interactive", "n", false,
		"Set this flag to skip user confirmation for apply change")

	applyChangesCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false,
		"Set this flag to display the diff only and skip applying the changes")

	applyChangesCmd.Flags().BoolVarP(&quiet, "quiet", "q", false,
		"Set this flag to suppress the diff output for apply changes")
}

var applyChangesFunc = func(cmd *cobra.Command, args []string) {
	c := setupOpsmanClient()
	tl := tile.NewTilesLoader(c)
	ml := manifest.NewManifestsLoader(c, tl)

	var slugs []string
	if len(products) == 0 {
		printMessage("Applying changes to all products")
	} else {
		printMessage("Applying changes to these products:", products)
		products = strings.TrimSpace(products)
		for _, s := range strings.Split(products, ",") {
			slugs = append(slugs, strings.TrimSpace(s))
		}
	}

	options := applychanges.ApplyChangesOptions{
		TileSlugs:      slugs,
		NonInteractive: nonInteractive,
		DryRun:         dryRun,
		Quiet:          quiet,
	}
	op := applychanges.NewApplyChangesOp(ml, tl, c, rp, options)

	err := op.Execute()

	if err != nil {
		rp.Fail(err)
	}
}

func printMessage(message ... string) {
	if quiet {
		fmt.Fprintln(os.Stderr, message)
	} else {
		fmt.Println(message)
	}
}
