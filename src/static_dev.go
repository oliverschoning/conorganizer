//go:build dev
// +build dev

package main

import (
	"log/slog"
	"net/http"
	"os"
)

func static(logger *slog.Logger) http.Handler {
	logger.Info("Static assets are being served from public folder")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.FileServerFS(os.DirFS("public")).ServeHTTP(w, r)
	})
}
