//go:build dev
// +build dev

package main

import (
	"log/slog"
	"net/http"
)

func static(logger *slog.Logger) http.Handler {
	logger.Info("Static assets are being served from web/static/")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.FileServer(http.Dir("/static")).ServeHTTP(w, r)
	})
}
