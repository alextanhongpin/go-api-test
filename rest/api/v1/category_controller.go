package v1

import (
	"errors"
	"net/http"

	"github.com/alextanhongpin/core/http/response"
	"github.com/alextanhongpin/go-api-test/rest/gate"
	"golang.org/x/exp/slog"
)

type CategoryController struct{}

func (h *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Is there a reason to use `gate` instead of calling directly say
	// `userID := UserIdFromContext(ctx)`?
	// Perhaps there are additional information that we want to get from the
	// database for the particular user, so using gate abstracts the fetching of
	// the user from the database.
	// For such scenario, instead of calling `gate.New`, we can pass down a gate
	// factory using dependency injection.
	// If user id is all you need, just call UserIdFromContext(ctx).
	g, err := gate.New(ctx)
	if err != nil {
		response.JSONError(w, err)
		return
	}
	userID := g.User().ID()

	slog.Info("got user id", slog.String("userID", userID.String()))

	if !g.Allow(&gate.CategoryCreator{}) {
		slog.Error("not allowed to create category", slog.String("userID", userID.String()))
		response.JSONError(w, response.ErrForbidden)
		return
	}

	response.JSONError(w, response.ErrUnknown)
}

func (h *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, response.OK(&Category{Name: "Toys"}), http.StatusOK)
}

func (h *CategoryController) List(w http.ResponseWriter, r *http.Request) {
	response.JSONError(w, errors.New("not implemented"))
}

func (h *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	response.JSONError(w, errors.New("not implemented"))
}

func (h *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	response.JSONError(w, errors.New("not implemented"))
}
