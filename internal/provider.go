package internal

type Provider interface {
	Initialize() error
	Fetch(cache Cache) error
	FetchOne(cache Cache, key string) (bool, error)
}
