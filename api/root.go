package api

import (
	"github.com/artemkaxdxd/mini-service/api/controller"
	"github.com/artemkaxdxd/mini-service/api/middleware"
	"github.com/artemkaxdxd/mini-service/app/image"
	"github.com/artemkaxdxd/mini-service/app/user"
	"github.com/go-chi/chi/v5"
)

func InitWeb(image *image.ImageService, user *user.UserService) *chi.Mux {
	router := chi.NewRouter()

	imageController := controller.NewImageController(image)
	userController := controller.NewUserController(user)

	router.Post("/login", userController.Login)

	router.Route("/", func(r chi.Router) {
		r.Use(middleware.ValidateToken)
		r.Post("/upload-picture", imageController.Upload)
		r.Get("/images", imageController.Get)
	})

	return router
}
