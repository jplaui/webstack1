package backend

import (
	"context"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// takes fiber context, templ component and additional component handlers
func Render(c *fiber.Ctx, componentHandler http.Handler) error {
	return adaptor.HTTPHandler(componentHandler)(c)
}

// used to pass state variables in the context
func SetRequestCtx(next http.Handler, ctxKey, ctxValue any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxKey, ctxValue)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func HandleComponent(component templ.Component, options ...func(*templ.ComponentHandler)) http.Handler {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return componentHandler
}
