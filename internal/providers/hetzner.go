package providers

import (
	"context"
	"errors"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/vspcsi/inventory/internal"
	"os"
	"strings"
)

type Hetzner struct {
	client hcloud.Client
}

func NewHetzner() *Hetzner {
	return &Hetzner{}
}

func (provider *Hetzner) Initialize() error {
	key, ok := os.LookupEnv("HETZNER_TOKEN")
	if !ok {
		return errors.New("count not find HETZNER_TOKEN, please assert it is set")
	}

	provider.client = *hcloud.NewClient(hcloud.WithToken(key))
	return nil
}

func (provider *Hetzner) Fetch(cache internal.Cache) error {
	instances, err := provider.client.Server.All(context.Background())
	if err != nil {
		return errors.New("could not fetch list of instances")
	}

	for _, instance := range instances {
		tags := strings.Split(instance.Labels["tags"], ".")
		cache.Create(instance.PublicNet.IPv4.IP.String(), instance.Name, tags)
	}
	return nil
}

func (provider *Hetzner) FetchOne(cache internal.Cache, key string) (bool, error) {
	return false, errors.New("not implemented")
}
