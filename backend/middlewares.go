package backend

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"

	"github.com/jplaui/templ-fiber/frontend"
)

func NotFoundMiddleware(c *fiber.Ctx) error {
	ch := HandleComponent(frontend.NotFound(), templ.WithStatus(http.StatusNotFound))
	return Render(c, ch)
}
