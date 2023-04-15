package main

import (
	_ "embed"
	"os"

	"github.com/alextanhongpin/go-api-test/rest"
	"github.com/alextanhongpin/go-core-microservice/http/server"
	"golang.org/x/exp/slog"
)

func main() {
	h := rest.New()
	textHandler := slog.NewTextHandler(os.Stdout)
	logger := slog.New(textHandler)
	server.New(logger, h, 8080)
}
