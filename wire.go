// In this layer, we build the dependencies that are required to run the server.
// We start by building the infrastructures (db, cache etc), then the
// middleware, then the usecases, and finally the handlers.

//go:build wireinject

//go:generate wire .
package main

import (
	"context"
	"net/http"

	"github.com/alextanhongpin/errcodes/stacktrace"
	"github.com/alextanhongpin/go-api-test/config"
	"github.com/alextanhongpin/go-api-test/rest"
	"github.com/alextanhongpin/go-api-test/rest/api"
	v1 "github.com/alextanhongpin/go-api-test/rest/api/v1"
	"github.com/alextanhongpin/go-api-test/rest/middleware"
	"github.com/alextanhongpin/go-api-test/rest/security"
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
		api.NewHealthController,
	)

	authControllerSet = wire.NewSet(
		provideTokenSigner,
		wire.Struct(new(api.AuthController), "*"),
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
		authControllerSet,
		wire.Struct(new(api.API), "*"),
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
		provideAuthMiddleware,

		// APIs.
		rootSet,
		v1Set,

		// Controller.
		provideRouter,
	))
}

type productUsecase struct{}

func (uc *productUsecase) Find(ctx context.Context, id string) (*v1.Product, error) {
	return nil, stacktrace.Wrap(stacktrace.New("bad product"), id)
	//return &v1.Product{
	//Name: "red socks",
	//}, nil
}

func (uc *productUsecase) List(ctx context.Context) ([]v1.Product, error) {
	return []v1.Product{
		{Name: "red socks"},
		{Name: "green socks"},
		{Name: "blue socks"},
	}, nil
}

func provideTokenSigner(cfg *config.Config) *security.TokenSigner {
	return security.NewTokenSigner([]byte(cfg.JWT.Secret))
}

func provideAuthMiddleware(cfg *config.Config) middleware.Middleware {
	return middleware.RequireAuth([]byte(cfg.JWT.Secret))
}

func provideRouter(
	root *api.API,
	v1 *v1.API,
) http.Handler {
	r := rest.New()

	// Register routes.
	root.Register(r)
	v1.Register(r)

	return r
}
