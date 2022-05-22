package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/test_server/internal/infra/http/controllers"
)

func Router(eventController *controllers.EventController) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/grass", func(apiRouter chi.Router) {
			AddEventRoutes(&apiRouter, eventController)

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/FindAllUsers",
			eventController.FindAll(),
		)
		apiRouter.Get(
			"/FindByName/{name}",
			eventController.FindByName(),
		)

		apiRouter.Get(
			"/Create/{name}/{age}/{city}/{country}",
			eventController.CreateUser(),
		)
		apiRouter.Get(
			"/UpdateData/{id}/{name}/{age}/{city}/{country}",
			eventController.UpdateById(),
		)
	})
}
