package controller

import (
	"encoding/json"
	"net/http"

	utils "github.com/artemkaxdxd/mini-service"
	"github.com/artemkaxdxd/mini-service/entity"
)

type UserService interface {
	Get(username string) (*entity.User, error)
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserController struct {
	user UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{user: service}
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

	user, err := u.user.Get(data.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
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
