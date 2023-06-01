package api

import (
	"database/sql"

	"github.com/artemkaxdxd/mini-service/api/controller"
	"github.com/artemkaxdxd/mini-service/api/middleware"
	"github.com/go-chi/chi/v5"
)

func InitWeb(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	user := controller.NewUserController(db)
	image := controller.NewImageController(db)

	router.Post("/login", user.Login)

	router.Route("/", func(r chi.Router) {
		r.Use(middleware.ValidateToken)
		r.Post("/upload-picture", image.Upload)
		r.Get("/images", image.Get)
	})

	return router
}
