package provider

func New(p Provider) {
	p.PreCreate()
}
