package model

type MiddlewareContext interface {
	Set(key string, value any)
	Get(key string) any
	MustGet(key string) any
	Has(key string) bool
}

type MContext map[string]any

func (m MContext) Set(key string, value any) {
	m[key] = value
}

func (m MContext) Get(key string) any {
	return m[key]
}

func (m MContext) MustGet(key string) any {
	if value, ok := m[key]; ok {
		return value
	}
	panic("key not found: " + key)
}

func (m MContext) Has(key string) bool {
	_, ok := m[key]
	return ok
}
