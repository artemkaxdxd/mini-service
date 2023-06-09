package user

import "github.com/artemkaxdxd/mini-service/entity"

type IUserRepo interface {
	Get(username string) (*entity.User, error)
}

type UserService struct {
	user IUserRepo
}

func NewUserService(repo IUserRepo) *UserService {
	return &UserService{user: repo}
}

func (service *UserService) Get(username string) (*entity.User, error) {
	user, err := service.user.Get(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
