package internal

type Cache interface {
	Create(key string, label string, values []string)
	Delete(key string)

	GetLabel(key string) string
	GetTags(key string) []string
}
