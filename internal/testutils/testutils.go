package testutils

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
)

var (
	ContextType = mock.MatchedBy(func(ctx context.Context) bool {
		return true
	})
	StringType = mock.AnythingOfType("string")
)

var InternalServerErrorBytes = []byte(`{
	"error": {
		"code": "api.internal",
		"message": "Oops, something went wrong. Please try again later."
	}
}`)

func CmpJSON(t *testing.T, a, b []byte, opts ...cmp.Option) {
	var l, r map[string]any
	if err := json.Unmarshal(a, &l); err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(b, &r); err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(l, r, opts...); diff != "" {
		t.Errorf("want(+), got(-): %s", diff)
	}
}
