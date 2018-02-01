package tile

import (
	"encoding/json"
	"fmt"
	"time"
)

type omClient interface {
	Get(endpoint string, timeout time.Duration) ([]byte, error)
}

type Loader struct {
	client omClient
}

func NewTilesLoader(omClient omClient) Loader {
	return Loader{client: omClient}
}

func (l Loader) LoadStaged(fetchTileMetadata bool) (Tiles, error) {
	return l.load(fetchTileMetadata, "staged")
}

func (l Loader) LoadDeployed(fetchTileMetadata bool) (Tiles, error) {
	return l.load(fetchTileMetadata, "deployed")
}

func (l Loader) load(fetchTileMetadata bool, status string) (Tiles, error) {
	b, err := l.client.Get(fmt.Sprintf("/api/v0/%s/products", status), 10*time.Minute)

	if err != nil {
		return Tiles{}, err
	}

	var data []*Tile
	err = json.Unmarshal(b, &data)
	if err != nil {
		return Tiles{}, err
	}

	if fetchTileMetadata {
		for _, d := range data {
			err = l.loadTileMetadata(d)
			if err != nil {
				return Tiles{}, err
			}
		}
	}

	return Tiles{data}, nil

}

func (l Loader) loadTileMetadata(t *Tile) error {
	urlsToPointer := []struct {
		url     string
		pointer *map[string]interface{}
	}{
		{"/api/v0/staged/products/%s/networks_and_azs", &t.Networks},
		{"/api/v0/staged/products/%s/errands", &t.Errands},
		{"/api/v0/staged/products/%s/resources", &t.Resources},
		{"/api/v0/staged/products/%s/properties", &t.Properties},
	}

	for _, up := range urlsToPointer {
		data, err := l.client.Get(fmt.Sprintf(up.url, t.GUID), 10*time.Minute)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, up.pointer)
		if err != nil {
			return err
		}
	}

	return nil
}
