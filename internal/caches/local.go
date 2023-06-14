package caches

import (
	"github.com/vspcsi/inventory/internal"
	"log"
)

type Data struct {
	Label string
	Tags  []string
}
type Local struct {
	data      map[string]*Data
	providers []internal.Provider
}

func NewLocal(providers []internal.Provider) *Local {
	self := &Local{
		providers: providers,
		data:      map[string]*Data{},
	}

	for _, provider := range self.providers {
		if err := provider.Initialize(); err != nil {
			log.Fatalln(err)
		}

		if err := provider.Fetch(self); err != nil {
			log.Fatalln(err)
		}
	}
	return self
}

func (receiver *Local) Create(key string, label string, values []string) {
	receiver.data[key] = &Data{
		Label: label,
		Tags:  values,
	}
}

func (receiver *Local) Delete(key string) {
	delete(receiver.data, key)
}

func (receiver *Local) GetTags(key string) []string {
	if data, ok := receiver.data[key]; ok {
		return data.Tags
	}

	for _, provider := range receiver.providers {
		ok, err := provider.FetchOne(receiver, key)

		if err != nil {
			log.Fatalln(err)
		}

		if ok {
			return receiver.data[key].Tags
		}
	}
	return nil
}

func (receiver *Local) GetLabel(key string) string {
	if data, ok := receiver.data[key]; ok {
		return data.Label
	}

	for _, provider := range receiver.providers {
		ok, err := provider.FetchOne(receiver, key)

		if err != nil {
			log.Fatalln(err)
		}

		if ok {
			return receiver.data[key].Label
		}
	}
	return ""
}
