// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rest

import (
	"context"
	"github.com/alextanhongpin/go-api-test/rest/apis"
	"github.com/alextanhongpin/go-api-test/rest/apis/v1"
	"github.com/alextanhongpin/go-core-microservice/http/middleware"
	"github.com/go-chi/chi/v5"
	middleware2 "github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"
	"log"
	"net/http"
	"os"
	"time"
)

// Injectors from wire.go:

func New() http.Handler {
	healthHandlerConfig := provideHealthHandlerConfig()
	healthHandler := apis.NewHealthHandler(healthHandlerConfig)
	middleware := provideBearerMiddleware()
	api := &apis.API{
		HealthHandler: healthHandler,
		BearerMW:      middleware,
	}
	categoryHandler := &v1.CategoryHandler{}
	restProductUsecase := &productUsecase{}
	productHandler := v1.NewProductHandler(restProductUsecase)
	v1API := &v1.API{
		CategoryHandler: categoryHandler,
		ProductHandler:  productHandler,
	}
	handler := provide(api, v1API)
	return handler
}

// wire.go:

var (
	// Usecases set.
	productUsecaseSet = wire.NewSet(wire.Struct(new(productUsecase)), wire.Bind(new(v1.ProductUsecase), new(*productUsecase)))

	// Handlers set.
	healthHandlerSet = wire.NewSet(
		provideHealthHandlerConfig, apis.NewHealthHandler,
	)

	categoryHandlerSet = wire.NewSet(wire.Struct(new(v1.CategoryHandler)))

	productHandlerSet = wire.NewSet(
		productUsecaseSet, v1.NewProductHandler,
	)

	// APIs set.
	rootSet = wire.NewSet(
		healthHandlerSet, wire.Struct(new(apis.API), "*"),
	)

	v1Set = wire.NewSet(
		categoryHandlerSet,
		productHandlerSet, wire.Struct(new(v1.API), "*"),
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

func provideBearerMiddleware() middleware.Middleware {
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Fatal(`"JWT_SECRET is required"`)
	}

	return middleware.RequireAuth([]byte(secret))
}

func provideHealthHandlerConfig() *apis.HealthHandlerConfig {
	startAt := time.Now()
	buildDate := os.Getenv("BUILD_DATE")
	buildAt, err := time.Parse(time.RFC3339, buildDate)
	if err != nil {
		log.Fatalf("failed to parse BUILD_DATE: %v", err)
	}

	return &apis.HealthHandlerConfig{
		Name:    os.Getenv("APP_NAME"),
		Version: os.Getenv("APP_VERSION"),
		BuildAt: buildAt,
		StartAt: startAt,
		VCSRef:  os.Getenv("VCS_REF"),
		VCSURL:  os.Getenv("VCS_URL"),
	}
}

func provide(
	root *apis.API, v1_2 *v1.API,
) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware2.RequestID)
	r.Use(middleware2.RealIP)
	r.Use(middleware2.Logger)
	r.Use(middleware2.Recoverer)

	r.Use(middleware2.Timeout(60 * time.Second))

	root.Register(r)
	v1_2.
		Register(r)

	return r
}
