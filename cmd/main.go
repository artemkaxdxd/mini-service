package main

import (
	"net/http"

	"github.com/artemkaxdxd/mini-service/api"
	"github.com/artemkaxdxd/mini-service/repo"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := repo.New()
	defer db.Close()

	router := api.InitWeb(db)

	http.ListenAndServe(":3000", router)
}
