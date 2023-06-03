package main

import (
	"os"

	"github.com/alextanhongpin/core/http/server"
	"golang.org/x/exp/slog"
)

func main() {
	router := newRouter()

	textHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(textHandler)
	slog.SetDefault(logger)

	server.ListenAndServe(":8080", router)
}
