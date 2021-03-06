package cmd

import (
	"errors"
	"fmt"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cloudops/omen/internal/errands"
	"github.com/pivotal-cloudops/omen/internal/tile"
	"github.com/spf13/cobra"
)

var errandProductSlugs []string

var errandsCmd = &cobra.Command{
	Use:   "errands",
	Short: "list the errands and their state",
	Long:  "Display a list of errands, optionally filtered by the product name",
	Run:   errandsFunc,
}

func init() {
	errandsCmd.Flags().StringSliceVar(&errandProductSlugs, "products", []string{},
		`(Optional) A comma-delimited list of products for errand updates. When omitted, all products will be affected.`)
}

var errandsFunc = func(*cobra.Command, []string) {
	c := setupOpsmanClient()
	api := api.New(api.ApiInput{
		Client: c,
	})
	et := errands.NewErrandReporter(api, tr)
	tl := tile.NewTilesLoader(c)

	if len(errandProductSlugs) > 0 {
		guids, err := mapGuid(tl, errandProductSlugs)

		if err != nil {
			rp.Fail(err)
		}

		err = et.Execute(guids)

		if err != nil {
			rp.Fail(err)
		}
	} else {
		reportAllErrands(tl, et)
	}
}

func mapGuid(tl tile.Loader, productSlugs []string) ([]string, error) {
	var guids []string
	deployedProducts, err := tl.LoadDeployed(false)
	if err != nil {
		return nil, err
	}

	for _, productSlug := range productSlugs {
		_tile, err := deployedProducts.FindBySlug(productSlug)
		if err != nil {
			return nil, err
		}
		guids = append(guids, _tile.GUID)
	}
	return guids, err
}

func reportAllErrands(tl tile.Loader, er errands.ErrandReporter) {
	deployedProducts, err := tl.LoadDeployed(false)
	if err != nil {
		rp.Fail(errors.New(fmt.Sprintf("Unable to fetch deployed products:\n%#v", err)))
	}
	for _, product := range deployedProducts.Data {
		err := er.Execute([]string{product.GUID})
		if err != nil {
			rp.Fail(err)
		}
	}
}
