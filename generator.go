package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
)

var Data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: "github.com/diamondburned/gotk4-gnome-desktop/pkg",
		Packages: []genmain.Package{
			{Name: "gnome-desktop-4"},
			{Name: "gnome-desktop-3.0", Namespaces: []string{"GnomeDesktop-3"}},
		},
		PkgGenerated: []string{
			"gnomedesktop",
			"gnomerr",
			"gnomebg",
		},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
		},
	},
)

func main() {
	genmain.Run(Data)
}
