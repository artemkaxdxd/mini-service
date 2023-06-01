package api

import (
	"github.com/artemkaxdxd/mini-service/api/middleware"
	"github.com/artemkaxdxd/mini-service/app/image"
	"github.com/go-chi/chi/v5"
)

func InitWeb(image *image.ImageService, user *user.UserService) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/login", user.Login)

	router.Route("/", func(r chi.Router) {
		r.Use(middleware.ValidateToken)
		r.Post("/upload-picture", image.Upload)
		r.Get("/images", image.Get)
	})

	return router
}
