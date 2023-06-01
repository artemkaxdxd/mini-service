package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	utils "github.com/artemkaxdxd/mini-service"
	"github.com/artemkaxdxd/mini-service/entity"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserController struct {
	db *sql.DB
}

func NewUserController(database *sql.DB) *UserController {
	return &UserController{db: database}
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var data *LoginBody

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	sql := "SELECT * FROM users WHERE username=?"
	res, err := u.db.Query(sql, data.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	defer res.Close()

	var user entity.User
	if res.Next() {
		err := res.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("error: no user found"))
	}

	err = utils.CheckPassword(user.Password, data.Password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)
}
