package main

import (
	"embed"
	"io/fs"

	"rootin.com/internal"
)

//go:embed www/*
var content embed.FS

func main() {
	content, err := fs.Sub(content, "www")
	if err != nil {
		return
	}
	internal.Run(content)
}
