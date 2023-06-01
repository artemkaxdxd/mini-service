package repo

import (
	"database/sql"
	"errors"

	"github.com/artemkaxdxd/mini-service/entity"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password_hash"`
}

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{DB: db}
}

func (store *UserStore) Get(username string) (*entity.User, error) {
	sql := "SELECT * FROM users WHERE username=?"
	res, err := store.DB.Query(sql, username)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	var user *entity.User
	if res.Next() {
		err := res.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("not found")
	}

	return user, nil
}
