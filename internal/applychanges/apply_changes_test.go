package applychanges_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"io/ioutil"
	"reflect"

	"github.com/pivotal-cloudops/omen/internal/applychanges"
	"github.com/pivotal-cloudops/omen/internal/applychanges/applychangesfakes"
	"github.com/pivotal-cloudops/omen/internal/common"
	"github.com/pivotal-cloudops/omen/internal/fakes"
	"github.com/pivotal-cloudops/omen/internal/manifest"
	"github.com/pivotal-cloudops/omen/internal/tile"
)

var _ = Describe("Apply Changes - Execute", func() {
	var mockClient *applychangesfakes.FakeOpsmanClient

	BeforeEach(func() {
		mockClient = &applychangesfakes.FakeOpsmanClient{}
	})

	defReportPrinter := &applychangesfakes.FakeReportPrinter{}

	It("Applies all changes by default", func() {
		manifests := manifest.Manifests{}

		mloader := &applychangesfakes.FakeManifestsLoader{
			LoadAllStub: func(status common.ProductStatus) (manifest.Manifests, error) {
				return manifests, nil
			},
		}

		tloader := fakes.FakeTilesLoader{}

		subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, defReportPrinter, applychanges.ApplyChangesOptions{[]string{}, true, false, false})
		subject.Execute()

		postedUrl, postedBody, _ := mockClient.PostArgsForCall(0)
		Expect(postedUrl).To(Equal("/api/v0/installations"))
		Expect(postedBody).To(ContainSubstring(`"deploy_products": "all"`))
	})

	It("Applies changes with no difference between staged and deployed", func() {
		manifests := manifest.Manifests{}

		mloader := &applychangesfakes.FakeManifestsLoader{
			LoadAllStub: func(status common.ProductStatus) (manifest.Manifests, error) {
				return manifests, nil
			},
		}

		tloader := fakes.FakeTilesLoader{}

		subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, defReportPrinter, applychanges.ApplyChangesOptions{[]string{}, true, false, false})
		subject.Execute()

		postedUrl, postedBody, _ := mockClient.PostArgsForCall(0)
		Expect(postedUrl).To(Equal("/api/v0/installations"))
		Expect(postedBody).To(ContainSubstring(`"deploy_products": "all"`))
	})

	It("Applies changes with difference between staged and deployed", func() {
		stagedManifests := manifest.Manifests{
			Data: []manifest.Manifest{
				{
					Name: "staged",
				},
			},
		}
		deployedManifests := manifest.Manifests{
			Data: []manifest.Manifest{
				{
					Name: "deployed",
				},
			},
		}

		tloader := fakes.FakeTilesLoader{}

		mloader := &applychangesfakes.FakeManifestsLoader{
			LoadAllStub: func(status common.ProductStatus) (manifest.Manifests, error) {
				if status == common.DEPLOYED {
					return deployedManifests, nil
				}
				return stagedManifests, nil
			},
		}

		subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, defReportPrinter, applychanges.ApplyChangesOptions{[]string{}, true, false, false})
		subject.Execute()

		postedUrl, postedBody, _ := mockClient.PostArgsForCall(0)
		Expect(postedUrl).To(Equal("/api/v0/installations"))
		Expect(postedBody).To(ContainSubstring(`"deploy_products": "all"`))
	})

	It("Prints out the the diff between all staged and deployed tiles", func() {
		stagedManifests := manifest.Manifests{
			Data: []manifest.Manifest{
				{
					Name: "staged",
				},
			},
		}
		deployedManifests := manifest.Manifests{
			Data: []manifest.Manifest{
				{
					Name: "deployed",
				},
			},
		}

		mloader := &applychangesfakes.FakeManifestsLoader{
			LoadAllStub: func(status common.ProductStatus) (manifest.Manifests, error) {
				if status == common.DEPLOYED {
					return deployedManifests, nil
				}
				return stagedManifests, nil
			},
		}

		tloader := fakes.FakeTilesLoader{}

		rp := &applychangesfakes.FakeReportPrinter{}
		subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, rp, applychanges.ApplyChangesOptions{[]string{}, true, false, false})
		subject.Execute()
		diff := rp.PrintReportArgsForCall(0)
		Expect(diff).To(Equal("-manifests.deployed.name=deployed\n+manifests.staged.name=staged\n"))
	})

	Describe("selective tile deployments", func() {
		It("applies changes to specified products", func() {
			fetchTileMetadata := true
			manifests := manifest.Manifests{}
			mloader := &applychangesfakes.FakeManifestsLoader{
				LoadStub: func(status common.ProductStatus, tileGuids []string) (manifest.Manifests, error) {
					return manifests, nil
				},
			}

			tloader := fakes.FakeTilesLoader{
				StagedResponseFunc: func(b bool) (tile.Tiles, error) {
					fetchTileMetadata = b
					return tile.Tiles{
						Data: []*tile.Tile{
							{
								GUID: "guid1",
								Type: "product1",
							},
							{
								GUID: "guid2",
								Type: "product2",
							},
						},
					}, nil
				},
			}

			subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, defReportPrinter, applychanges.ApplyChangesOptions{[]string{"product1", "product2"}, true, false, false})
			subject.Execute()

			Expect(fetchTileMetadata).To(BeFalse())

			postedUrl, postedBody, _ := mockClient.PostArgsForCall(0)
			Expect(postedUrl).To(Equal("/api/v0/installations"))
			Expect(postedBody).To(ContainSubstring(`"deploy_products": "guid1,guid2"`))
		})

		It("fails when slug not found", func() {
			mloader := &applychangesfakes.FakeManifestsLoader{
				LoadStub: func(status common.ProductStatus, tileGuids []string) (manifest.Manifests, error) {
					return manifest.Manifests{}, nil
				},
			}

			tloader := fakes.FakeTilesLoader{
				StagedResponseFunc: func(b bool) (tile.Tiles, error) {
					return tile.Tiles{
						Data: []*tile.Tile{
							{
								GUID: "guid1",
								Type: "product1",
							},
							{
								GUID: "guid2",
								Type: "product2",
							},
						},
					}, nil
				},
			}

			subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, defReportPrinter, applychanges.ApplyChangesOptions{[]string{"product3", "product2"}, true, false, false})
			err := subject.Execute()

			Expect(err).To(HaveOccurred())

			Expect(mockClient.PostCallCount()).To(BeZero())
		})

		It("fails when tile loading fails", func() {
			mloader := &applychangesfakes.FakeManifestsLoader{
				LoadStub: func(status common.ProductStatus, tileGuids []string) (manifest.Manifests, error) {
					return manifest.Manifests{}, nil
				},
			}

			tloader := fakes.FakeTilesLoader{
				StagedResponseFunc: func(b bool) (tile.Tiles, error) {
					return tile.Tiles{}, errors.New("can't load tiles")
				},
			}

			subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, defReportPrinter, applychanges.ApplyChangesOptions{[]string{"product3"}, true, false, false})
			err := subject.Execute()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("can't load tiles"))
			Expect(mockClient.PostCallCount()).To(BeZero())
		})

		It("prints out the diff for only the tiles being deployed", func() {

			mloader := &applychangesfakes.FakeManifestsLoader{

				LoadAllStub: func(status common.ProductStatus) (manifest.Manifests, error) {
					return manifest.Manifests{}, errors.New("loadAll should not be called")
				},

				LoadStub: func(status common.ProductStatus, tileGuids []string) (manifest.Manifests, error) {
					if status == common.DEPLOYED && reflect.DeepEqual(tileGuids, []string{"product1", "product2"}) {
						return manifest.Manifests{
							Data: []manifest.Manifest{
								{
									Name: "product1",
								},
								{
									Name: "product2",
								},
							},
						}, nil
					}

					if status == common.STAGED && reflect.DeepEqual(tileGuids, []string{"product1", "product2"}) {
						return manifest.Manifests{
							Data: []manifest.Manifest{
								{
									Name: "staged-product1",
								},
								{
									Name: "staged-product2",
								},
							},
						}, nil
					}

					return manifest.Manifests{}, errors.New("don't know how to load these manifests")
				},
			}

			tloader := fakes.FakeTilesLoader{
				StagedResponseFunc: func(b bool) (tile.Tiles, error) {
					return tile.Tiles{
						Data: []*tile.Tile{
							{
								Type: "product1",
								GUID: "product1",
							},
							{
								Type: "product2",
								GUID: "product2",
							},
						},
					}, nil
				},
			}

			rp := &applychangesfakes.FakeReportPrinter{}

			subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, rp, applychanges.ApplyChangesOptions{[]string{"product1", "product2"}, true, false, false})
			subject.Execute()
			diff := rp.PrintReportArgsForCall(0)

			expectedDiff, err := ioutil.ReadFile("testdata/diff.txt")
			Expect(err).ToNot(HaveOccurred())

			Expect(diff).To(Equal(string(expectedDiff)))

		})
	})

	Describe("Dry run", func() {
		It("it outputs the diff but does not apply changes", func() {
			stagedManifests := manifest.Manifests{
				Data: []manifest.Manifest{
					{
						Name: "staged",
					},
				},
			}
			deployedManifests := manifest.Manifests{
				Data: []manifest.Manifest{
					{
						Name: "deployed",
					},
				},
			}

			mloader := &applychangesfakes.FakeManifestsLoader{
				LoadAllStub: func(status common.ProductStatus) (manifest.Manifests, error) {
					if status == common.DEPLOYED {
						return deployedManifests, nil
					}
					return stagedManifests, nil
				},
			}

			tloader := fakes.FakeTilesLoader{}

			rp := &applychangesfakes.FakeReportPrinter{}
			subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, rp, applychanges.ApplyChangesOptions{[]string{}, true, true, false})
			subject.Execute()
			diff := rp.PrintReportArgsForCall(0)
			Expect(diff).To(Equal("-manifests.deployed.name=deployed\n+manifests.staged.name=staged\n"))

			Expect(mockClient.PostCallCount()).To(BeZero())
		})
	})

	Describe("Quiet run", func() {
		It("only outputs the return of ops manager apply changes", func() {
			applyChangesReply := `{"install":{"id": 303}}`

			tloader := fakes.FakeTilesLoader{}
			rp := &applychangesfakes.FakeReportPrinter{}
			mloader := &applychangesfakes.FakeManifestsLoader{}

			mockClient.PostReturns([]byte(applyChangesReply), nil)

			subject := applychanges.NewApplyChangesOp(mloader, tloader, mockClient, rp, applychanges.ApplyChangesOptions{[]string{}, true, false, true})
			subject.Execute()
			Expect(rp.PrintReportCallCount()).To(Equal(1))

			Expect(rp.PrintReportArgsForCall(0)).To(MatchJSON(applyChangesReply))
		})
	})

})
