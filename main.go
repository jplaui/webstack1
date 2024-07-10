package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/jplaui/templ-fiber/backend"
)

var (
	g errgroup.Group
)

func main() {

	// initialize web application
	app := fiber.New()

	// routes.go groups all handler strings in a single file
	backend.AddRoutes(app)

	// start the app
	g.Go(func() error {
		err := app.Listen(":3000")
		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("closing fiber server")
			return err
		}
		return err
	})

	err := g.Wait()
	if err != nil {
		log.Error().Err(err).Msg("web exit")
		return
	}
}
