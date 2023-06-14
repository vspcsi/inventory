package providers

import (
	"encoding/json"
	"errors"
	"github.com/vspcsi/inventory/internal"
	"os"
)

type Entry struct {
	Address string   `json:"address"`
	Label   string   `json:"label"`
	Tags    []string `json:"tags"`
}

type Local struct{}

func NewLocal() *Local {
	return &Local{}
}

func (provider *Local) Initialize() error {
	return nil
}

func (provider *Local) Fetch(cache internal.Cache) error {
	filename := "tags.json"
	var data map[string]Entry

	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	file, err := os.ReadFile("tags.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	for address, entry := range data {
		cache.Create(address, entry.Label, entry.Tags)
	}
	return nil
}

func (provider *Local) FetchOne(cache internal.Cache, key string) (bool, error) {
	return false, errors.New("not implemented")
}
