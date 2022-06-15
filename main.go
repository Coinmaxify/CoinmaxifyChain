package main

import (
	_ "embed"
	"github.com/Coinmaxify/Coinmaxify/command/root"
	"github.com/Coinmaxify/Coinmaxify/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
