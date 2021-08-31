package pkg

type Provider interface {
	Name() string
	PreCreate() error
}
