package provider

type Provider interface {
	Name() string
	PreCreate() error
}
