package manifest_test

import (
	"errors"
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

		"github.com/pivotal-cloudops/omen/internal/fakes"
	"github.com/pivotal-cloudops/omen/internal/manifest"
	"github.com/pivotal-cloudops/omen/internal/tile"
)

var _ = Describe("Manifests loader", func() {
	DescribeTable("Load all manifests", func(status string) {
		fakeOMClient := fakes.FakeOMClient{
			GetFunc: func(endpoint string) ([]byte, error) {
				switch endpoint {
				case fmt.Sprintf("/api/v0/%s/products", status):
					return ioutil.ReadFile("testdata/tiles.json")
				case fmt.Sprintf("/api/v0/%s/products/guid/manifest", status):
					return ioutil.ReadFile(fmt.Sprintf("testdata/%s/manifest.json", status))
				case fmt.Sprintf("/api/v0/%s/director/manifest", status):
					return ioutil.ReadFile(fmt.Sprintf("testdata/%s/p-bosh-manifest.json", status))
				case fmt.Sprintf("/api/v0/%s/cloud_config", status):
					return ioutil.ReadFile("testdata/cloud_config.json")
				default:
					return nil, errors.New(fmt.Sprintf("invalid endpoint %v", endpoint))
				}
			},
		}

		tl := tile.NewTilesLoader(fakeOMClient)
		loader := manifest.NewManifestsLoader(fakeOMClient, tl)

		var (
			manifests manifest.Manifests
			err       error
		)

		if status == "deployed" {
			manifests, err = loader.LoadAllDeployed()
		} else {
			manifests, err = loader.LoadAllStaged()
		}

		Expect(err).ToNot(HaveOccurred())
		Expect(manifests.Data).To(HaveLen(2))

		manifest := manifests.Data[0]

		Expect(manifest.Name).To(Equal("guid"))
		Expect(manifest.Releases).To(HaveLen(1))
		Expect(manifest.InstanceGroups).To(HaveLen(1))
		Expect(manifest.Stemcells).To(HaveLen(1))
		Expect(manifest.Update).ToNot(BeEmpty())
		Expect(manifest.Variables).To(HaveLen(1))

		directorManifest := manifests.Data[1]

		Expect(directorManifest.Name).To(Equal("p-bosh"))
		Expect(directorManifest.Releases).To(HaveLen(1))
		Expect(directorManifest.InstanceGroups).To(HaveLen(1))
		Expect(directorManifest.Stemcells).To(HaveLen(1))
		Expect(directorManifest.Update).ToNot(BeEmpty())
		Expect(directorManifest.Variables).To(HaveLen(1))

		Expect(manifests.CloudConfig).ToNot(BeEmpty())
	},
		Entry("should load staged", "staged"),
		Entry("should load deployed", "deployed"))

	Describe("load", func() {
		It("fetches the staged manifests for the specific tile guids", func() {
			fakeOMClient := fakes.FakeOMClient{
				GetFunc: func(endpoint string) ([]byte, error) {
					switch endpoint {
					case "/api/v0/staged/products":
						return ioutil.ReadFile("testdata/tiles.json")
					case "/api/v0/staged/products/guid/manifest":
						return ioutil.ReadFile("testdata/staged/manifest.json")
					case "/api/v0/staged/cloud_config":
						return ioutil.ReadFile("testdata/cloud_config.json")
					default:
						return nil, errors.New(fmt.Sprintf("invalid endpoint %v", endpoint))
					}
				},
			}

			tl := tile.NewTilesLoader(fakeOMClient)
			loader := manifest.NewManifestsLoader(fakeOMClient, tl)

			manifests, err := loader.LoadStaged([]string{"guid"})
			Expect(err).ToNot(HaveOccurred())
			Expect(len(manifests.Data)).To(Equal(1))

			manifest := manifests.Data[0]
			Expect(manifest.Name).To(Equal("guid"))
		})
	})

	It("should report error if loadCloudConfig throws an error", func() {
		fakeOMClient := fakes.FakeOMClient{
			GetFunc: func(endpoint string) ([]byte, error) {
				switch endpoint {
				case "/api/v0/deployed/products":
					return ioutil.ReadFile("testdata/tiles.json")
				case "/api/v0/deployed/products/guid/manifest":
					return ioutil.ReadFile("testdata/deployed/manifest.json")
				case "/api/v0/deployed/director/manifest":
					return ioutil.ReadFile("testdata/deployed/p-bosh-manifest.json")
				case "/api/v0/deployed/cloud_config":
					return nil, errors.New("cloud config error")
				default:
					return nil, errors.New(fmt.Sprintf("invalid endpoint %v", endpoint))
				}
			},
		}

		tl := tile.NewTilesLoader(fakeOMClient)
		loader := manifest.NewManifestsLoader(fakeOMClient, tl)

		_, err := loader.LoadAllDeployed()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("cloud config error"))
	})

	It("should fail if the tile loader returns an error", func() {
		tl := fakes.FakeTilesLoader{
			DeployedResponseFunc: func(_ bool) (tile.Tiles, error) {
				return tile.Tiles{}, errors.New("oops")
			},
		}

		fakeOMClient := fakes.FakeOMClient{
			GetFunc: func(endpoint string) ([]byte, error) {
				return []byte{}, nil
			},
		}

		loader := manifest.NewManifestsLoader(fakeOMClient, tl)
		_, err := loader.LoadAllDeployed()
		Expect(err).To(HaveOccurred())
	})

	It("should fail if fetching from om returns an error", func() {

		fakeOMClient := fakes.FakeOMClient{
			GetFunc: func(endpoint string) ([]byte, error) {
				switch endpoint {
				case "/api/v0/deployed/products":
					return ioutil.ReadFile("testdata/tiles.json")
				default:
					return []byte{}, errors.New("oops")
				}
			},
		}

		tl := tile.NewTilesLoader(fakeOMClient)
		loader := manifest.NewManifestsLoader(fakeOMClient, tl)
		_, err := loader.LoadAllDeployed()
		Expect(err).To(HaveOccurred())

	})
})
