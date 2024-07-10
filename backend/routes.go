package backend

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func AddRoutes(app *fiber.App) {

	// context timeouts
	shortT := 2 * time.Second
	// longT := 10 * time.Second

	// PAGES
	ip := NewIndexPage("test-state")
	getWelcome := timeout.NewWithContext(ip.GetWelcome, shortT)
	sp := NewScanPage("blabla")
	// postResult := timeout.NewWithContext(sp.PostScan, longT)

	// routes for static files (e.g. css, js, favicons, etc...)
	app.Static("/assets", "assets")

	// main page
	app.Get("/", getWelcome)
	app.Post("/scan", sp.PostScan)

	// page not found
	app.Use(NotFoundMiddleware)
}
