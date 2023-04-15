// In this layer, we build the dependencies that are required to run the server.
// We start by building the infrastructures (db, cache etc), then the
// middlewares, then the usecases, and finally the handlers.

//go:build wireinject

//go:generate wire .
package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alextanhongpin/go-api-test/rest/apis"
	v1 "github.com/alextanhongpin/go-api-test/rest/apis/v1"
	httpmiddleware "github.com/alextanhongpin/go-core-microservice/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"
)

var (
	// Usecases set.
	productUsecaseSet = wire.NewSet(
		wire.Struct(new(productUsecase)),
		wire.Bind(new(v1.ProductUsecase), new(*productUsecase)),
	)

	// Handlers set.
	healthHandlerSet = wire.NewSet(
		provideHealthHandlerConfig,
		apis.NewHealthHandler,
	)

	categoryHandlerSet = wire.NewSet(
		wire.Struct(new(v1.CategoryHandler)),
	)

	productHandlerSet = wire.NewSet(
		productUsecaseSet,
		v1.NewProductHandler,
	)

	// APIs set.
	rootSet = wire.NewSet(
		healthHandlerSet,
		wire.Struct(new(apis.API), "*"),
	)

	v1Set = wire.NewSet(
		categoryHandlerSet,
		productHandlerSet,
		wire.Struct(new(v1.API), "*"),
	)
)

func New() http.Handler {
	panic(wire.Build(
		// Middlewares.
		provideBearerMiddleware,

		// APIs.
		rootSet,
		v1Set,

		// Handler.
		provide,
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

func provideBearerMiddleware() httpmiddleware.Middleware {
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Fatal(`"JWT_SECRET is required"`)
	}

	return httpmiddleware.RequireAuth([]byte(secret))
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
	root *apis.API,
	v1 *v1.API,
) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Register routes.
	root.Register(r)
	v1.Register(r)

	return r
}
