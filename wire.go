// In this layer, we build the dependencies that are required to run the server.
// We start by building the infrastructures (db, cache etc), then the
// middlewares, then the usecases, and finally the handlers.

//go:build wireinject

//go:generate wire .
package main

import (
	"context"
	"net/http"

	"github.com/alextanhongpin/go-api-test/config"
	"github.com/alextanhongpin/go-api-test/rest"
	"github.com/alextanhongpin/go-api-test/rest/apis"
	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
	httpmiddleware "github.com/alextanhongpin/go-core-microservice/http/middleware"
	"github.com/google/wire"
)

var (
	// Usecases set.
	productUsecaseSet = wire.NewSet(
		wire.Struct(new(productUsecase)),
		wire.Bind(new(v1.ProductUsecase), new(*productUsecase)),
	)

	// Controllers set.
	healthControllerSet = wire.NewSet(
		apis.NewHealthController,
	)

	categoryControllerSet = wire.NewSet(
		wire.Struct(new(v1.CategoryController)),
	)

	productControllerSet = wire.NewSet(
		productUsecaseSet,
		v1.NewProductController,
	)

	// APIs set.
	rootSet = wire.NewSet(
		healthControllerSet,
		wire.Struct(new(apis.API), "*"),
	)

	v1Set = wire.NewSet(
		categoryControllerSet,
		productControllerSet,
		wire.Struct(new(v1.API), "*"),
	)
)

func newRouter() http.Handler {
	panic(wire.Build(
		// Provide config.
		config.New,

		// Middlewares.
		provideBearerMiddleware,

		// APIs.
		rootSet,
		v1Set,

		// Controller.
		provideRouter,
	))
}

type productUsecase struct{}

func (uc *productUsecase) Find(ctx context.Context, id string) (*v1.Product, error) {
	return &v1.Product{
		Name: "red socks",
	}, nil
}

func (uc *productUsecase) List(ctx context.Context) ([]v1.Product, error) {
	return []v1.Product{
		{Name: "red socks"},
		{Name: "green socks"},
		{Name: "blue socks"},
	}, nil
}

func provideBearerMiddleware(cfg *config.Config) httpmiddleware.Middleware {
	return httpmiddleware.RequireAuth([]byte(cfg.JWT.Secret))
}

func provideRouter(
	root *apis.API,
	v1 *v1.API,
) http.Handler {
	r := rest.New()

	// Register routes.
	root.Register(r)
	v1.Register(r)

	return r
}
