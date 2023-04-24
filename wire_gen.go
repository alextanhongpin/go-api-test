// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/alextanhongpin/go-api-test/config"
	"github.com/alextanhongpin/go-api-test/rest"
	"github.com/alextanhongpin/go-api-test/rest/apis"
	"github.com/alextanhongpin/go-api-test/rest/apis/v1"
	"github.com/alextanhongpin/go-core-microservice/http/middleware"
	"github.com/google/wire"
	"net/http"
)

// Injectors from wire.go:

func newRouter() http.Handler {
	configConfig := config.New()
	middleware := provideBearerMiddleware(configConfig)
	healthController := apis.NewHealthController(configConfig)
	api := &apis.API{
		BearerAuth:       middleware,
		HealthController: healthController,
	}
	categoryController := &v1.CategoryController{}
	mainProductUsecase := &productUsecase{}
	productController := v1.NewProductController(mainProductUsecase)
	v1API := &v1.API{
		BearerAuth:         middleware,
		CategoryController: categoryController,
		ProductController:  productController,
	}
	handler := provideRouter(api, v1API)
	return handler
}

// wire.go:

var (
	// Usecases set.
	productUsecaseSet = wire.NewSet(wire.Struct(new(productUsecase)), wire.Bind(new(v1.ProductUsecase), new(*productUsecase)))

	// Controllers set.
	healthControllerSet = wire.NewSet(apis.NewHealthController)

	categoryControllerSet = wire.NewSet(wire.Struct(new(v1.CategoryController)))

	productControllerSet = wire.NewSet(
		productUsecaseSet, v1.NewProductController,
	)

	// APIs set.
	rootSet = wire.NewSet(
		healthControllerSet, wire.Struct(new(apis.API), "*"),
	)

	v1Set = wire.NewSet(
		categoryControllerSet,
		productControllerSet, wire.Struct(new(v1.API), "*"),
	)
)

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

func provideBearerMiddleware(cfg *config.Config) middleware.Middleware {
	return middleware.RequireAuth([]byte(cfg.JWT.Secret))
}

func provideRouter(
	root *apis.API, v1_2 *v1.API,
) http.Handler {
	r := rest.New()

	root.Register(r)
	v1_2.
		Register(r)

	return r
}
