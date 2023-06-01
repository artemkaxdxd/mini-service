package main

import (
	"net/http"

	"github.com/artemkaxdxd/mini-service/api"
	"github.com/artemkaxdxd/mini-service/app/image"
	"github.com/artemkaxdxd/mini-service/app/user"
	"github.com/artemkaxdxd/mini-service/repo"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := repo.New()
	defer db.Close()

	imageRepo := repo.NewImageStore(db)
	userRepo := repo.NewUserStore(db)

	imageService := image.NewImageService(imageRepo)
	userService := user.NewUserService(userRepo)

	router := api.InitWeb(imageService, userService)

	http.ListenAndServe(":3000", router)
}
