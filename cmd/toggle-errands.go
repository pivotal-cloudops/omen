package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cloudops/omen/internal/errands"
	"github.com/pivotal-cloudops/omen/internal/tile"
	"github.com/spf13/cobra"
)

var (
	errandAction         string
	errandType           string
	toggleErrandProducts []string

	actionEnable  = "enable"
	actionDisable = "disable"
	actionDefault = "default"
)

var toggleErrandsCmd = &cobra.Command{
	Use:   "toggle-errands",
	Short: "toggle the errand state for products",
	Long:  "Set the errand state for a list of products or all products",
	Run:   toggleErrandsFunc,
}

func init() {
	toggleErrandsCmd.Flags().StringVarP(&errandAction, "action", "a", "",
		`Set the toggle errand action. Valid values are: enable, disable, default`)

	toggleErrandsCmd.Flags().StringVar(&errandType, "errand-type", "post-deploy",
		`(Optional) Set to the errand type that you want to update. Only supported value is "post-deploy"`)

	toggleErrandsCmd.Flags().StringSliceVar(&toggleErrandProducts, "products", []string{},
		`(Optional) A comma-delimited list of product guids or slugs (e.g. p-redis) for errand updates. When omitted, all products will be affected.`)
}

var toggleErrandsFunc = func(*cobra.Command, []string) {
	validateFlags()
	c := setupOpsmanClient()
	es := api.New(api.ApiInput{
		Client: c,
	})
	et := newErrandToggler(es)

	products := "all"
	if len(toggleErrandProducts) > 0 {
		products = strings.Join(toggleErrandProducts, ",")
	}

	rep := fmt.Sprintf("Action: %s, Errand-Type: %s, Products: %s", errandAction, errandType, products)
	rp.PrintReport(rep)
	tl := tile.NewTilesLoader(c)

	if len(toggleErrandProducts) > 0 {
		toggleErrandsForProducts(tl, et, toggleErrandProducts)
	} else {
		toggleAllErrands(tl, et)
	}
}

func toggleErrandsForProducts(tl tile.Loader, et errands.ErrandToggler, products []string) {
	tiles, err := tl.LoadDeployed(false)
	if err != nil {
		rp.Fail(errors.New(fmt.Sprintf("Unable to fetch deployed products:\n%#v", err)))
	}

	tileGUIDs := mapProductNamesOrGUIDsToGUIDs(tiles, products)

	err = et.Execute(tileGUIDs)

	if err != nil {
		rp.Fail(err)
	}
}

func mapProductNamesOrGUIDsToGUIDs(tiles tile.Tiles, products []string) []string {
	foundTiles, err := tiles.FindBySlugsOrGUIDs(products)
	if err != nil {
		rp.Fail(errors.New(fmt.Sprintf("Unable to find products:\n%#v", err)))
	}
	tileGUIDs := make([]string, 0)
	for _, foundTile := range foundTiles {
		tileGUIDs = append(tileGUIDs, foundTile.GUID)
	}
	return tileGUIDs
}

func newErrandToggler(api api.Api) errands.ErrandToggler {
	et := errands.NewErrandToggler(api, rp)
	if errandAction == actionEnable {
		return et.Enable()
	} else if errandAction == actionDefault {
		return et.Default()
	}
	return et
}

func toggleAllErrands(tl tile.Loader, et errands.ErrandToggler) {
	deployedProducts, err := tl.LoadDeployed(false)
	if err != nil {
		rp.Fail(errors.New(fmt.Sprintf("Unable to fetch deployed products:\n%#v", err)))
	}
	for _, product := range deployedProducts.Data {
		err := et.Execute([]string{product.GUID})
		if err != nil {
			rp.Fail(err)
		}
	}
}

func validateFlags() {
	if !isErrandActionValid(errandAction) {
		rp.Fail(errors.New("invalid value specified for mandatory flag 'action'"))
	}

	if errandType != "post-deploy" {
		rp.Fail(errors.New("invalid value specified for mandatory flag 'errand-type'"))
	}
}

func isErrandActionValid(action string) bool {
	_, ok := map[string]interface{}{
		actionEnable:  nil,
		actionDisable: nil,
		actionDefault: nil,
	}[action]
	return ok
}
