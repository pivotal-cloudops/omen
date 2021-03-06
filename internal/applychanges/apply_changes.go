package applychanges

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pivotal-cloudops/omen/internal/diff"
	"github.com/pivotal-cloudops/omen/internal/manifest"
	"github.com/pivotal-cloudops/omen/internal/tile"
	"github.com/pivotal-cloudops/omen/internal/userio"
)

type applyChangesBody struct {
	IgnoreWarnings bool        `json:"ignore_warnings"`
	DeployProducts interface{} `json:"deploy_products"`
}

type ApplyChangesOptions struct {
	TileSlugs      []string
	NonInteractive bool
	DryRun         bool
	Quiet          bool
}

//go:generate counterfeiter . manifestsLoader
type manifestsLoader interface {
	LoadAllDeployed() (manifest.Manifests, error)
	LoadAllStaged() (manifest.Manifests, error)
	LoadDeployed(tileGuids []string) (manifest.Manifests, error)
	LoadStaged(tileGuids []string) (manifest.Manifests, error)
}

type tilesLoader interface {
	LoadStaged(bool) (tile.Tiles, error)
	LoadDeployed(bool) (tile.Tiles, error)
}

//go:generate counterfeiter . reportPrinter
type reportPrinter interface {
	PrintReport(string)
}

//go:generate counterfeiter . opsmanClient
type opsmanClient interface {
	Post(endpoint, data string, timeout time.Duration) ([]byte, error)
}

type ApplyChangesOp interface {
	Execute() error
}

type applyChangesOp struct {
	manifestsLoader manifestsLoader
	tilesLoader     tilesLoader
	opsmanClient    opsmanClient
	reportPrinter   reportPrinter
	options         ApplyChangesOptions
}

func NewApplyChangesOp(ml manifestsLoader, tl tilesLoader, c opsmanClient, rp reportPrinter, options ApplyChangesOptions) ApplyChangesOp {
	return &applyChangesOp{
		manifestsLoader: ml,
		tilesLoader:     tl,
		opsmanClient:    c,
		reportPrinter:   rp,
		options:         options,
	}
}

func (a *applyChangesOp) Execute() error {
	tileGuids, err := a.slugsToGuids()
	if err != nil {
		return err
	}

	if a.shouldPrintOutput() {
		manifestDiff, err := a.makeDiff(tileGuids)

		if err != nil {
			fmt.Println(err)
			return err
		}

		if len(manifestDiff) > 0 {
			a.reportPrinter.PrintReport(manifestDiff)
		} else if a.isNotADryRun() {
			a.reportPrinter.PrintReport("Warning: Opsman has detected no pending changes")
		}
	}

	if a.options.DryRun {
		return nil
	}

	if a.isInteractive() {
		proceed := userio.GetConfirmation("Do you wish to continue (y/n)?")

		if proceed == false {
			fmt.Println("Cancelled apply changes")
			return nil
		}

		fmt.Println("Applying changes")
	}

	return a.applyChanges(tileGuids)
}

func (a *applyChangesOp) isInteractive() bool {
	return a.options.NonInteractive == false
}

func (a *applyChangesOp) shouldPrintOutput() bool {
	return a.options.Quiet == false
}

func (a *applyChangesOp) isNotADryRun() bool {
	return a.options.DryRun == false
}

func (a *applyChangesOp) applyChanges(tileGuids []string) error {
	var guids interface{}
	if len(tileGuids) == 0 {
		guids = "all"
	} else {
		guids = tileGuids
	}

	bytes, err := json.Marshal(applyChangesBody{DeployProducts: guids, IgnoreWarnings: true})
	if err != nil {
		return err
	}
	body := string(bytes)

	resp, err := a.opsmanClient.Post("/api/v0/installations", body, 10*time.Minute)
	if err != nil {
		fmt.Printf("An error occurred applying changes: %v \n", err)
		return err
	}

	if a.options.Quiet {
		a.reportPrinter.PrintReport(string(resp))
	} else {
		a.reportPrinter.PrintReport(fmt.Sprintf("Successfully applied changes: %s \n", string(resp)))
	}
	return nil
}

func (a *applyChangesOp) slugsToGuids() ([]string, error) {
	if len(a.options.TileSlugs) == 0 {
		return []string{}, nil
	}

	tiles, err := a.tilesLoader.LoadStaged(false)
	if err != nil {
		return nil, err
	}
	var resp []string
	for _, s := range a.options.TileSlugs {
		t, err := tiles.FindBySlug(s)
		if err != nil {
			return nil, err
		}
		resp = append(resp, t.GUID)
	}
	return resp, nil
}

func (a *applyChangesOp) makeDiff(tileGuids []string) (string, error) {
	var (
		manifestA manifest.Manifests
		manifestB manifest.Manifests
		err       error
	)

	if len(tileGuids) == 0 {
		manifestA, err = a.manifestsLoader.LoadAllDeployed()
	} else {
		manifestA, err = a.manifestsLoader.LoadDeployed(tileGuids)
	}

	if err != nil {
		return "", err
	}

	if len(tileGuids) == 0 {
		manifestB, err = a.manifestsLoader.LoadAllStaged()
	} else {
		manifestB, err = a.manifestsLoader.LoadStaged(tileGuids)
	}

	if err != nil {
		return "", err
	}

	d, err := diff.FlatDiff(manifestA, manifestB)

	if err != nil {
		return "", err
	}

	return d, err
}
