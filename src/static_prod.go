//go:build !dev
// +build !dev

package main

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"

	hashFS "github.com/benbjohnson/hashfs"
)

//go:embed public/*
var staticFS embed.FS
var staticRootFS, _ = fs.Sub(staticFS, "/public")

func static(logger *slog.Logger) http.Handler {
	logger.Debug("Static assets are embedded")
	return hashFS.FileServer(staticRootFS)
}
