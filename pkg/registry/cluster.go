package registry

import "github.com/pavlelee/vscode-test-module/pkg/provider"

func New(p provider.Provider) {
	p.PreCreate()
}
