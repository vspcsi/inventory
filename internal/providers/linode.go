package providers

import (
	"context"
	"errors"
	"github.com/vspcsi/inventory/internal"
	"github.com/linode/linodego"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

type Linode struct {
	client linodego.Client
}

func NewLinode() *Linode {
	return &Linode{}
}

func (provider *Linode) Initialize() error {
	apiKey, ok := os.LookupEnv("LINODE_TOKEN")
	if !ok {
		return errors.New("could not find LINODE_TOKEN, please assert it is set")
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})

	oauth2Client := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
		},
	}

	provider.client = linodego.NewClient(oauth2Client)
	return nil
}

func (provider *Linode) Fetch(cache internal.Cache) error {
	instances, err := provider.client.ListInstances(context.Background(), nil)
	if err != nil {
		return errors.New("could not fetch list of instances")
	}

	for _, instance := range instances {
		for _, ip := range instance.IPv4 {
			if ip.IsPrivate() {
				continue
			}

			cache.Create(ip.String(), instance.Label, instance.Tags)
		}
	}
	return nil
}

func (provider *Linode) FetchOne(cache internal.Cache, key string) (bool, error) {
	return false, errors.New("not implemented")
}
