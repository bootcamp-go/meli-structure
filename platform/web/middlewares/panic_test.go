package middlewares_test

import (
	"bootcamp-web/platform/web/middlewares"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPanic(t *testing.T) {
	panicMiddleware := middlewares.NewPanic()
	err := panicMiddleware(func(w http.ResponseWriter, r *http.Request) error {
		panic("any panic")
	})(nil, nil)

	require.ErrorContains(t, err, "panic [any panic] trace [goroutine")
}
