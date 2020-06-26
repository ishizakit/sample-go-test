package service

import (
	"github.com/TechLoCo/sample-go-test/model"
	"github.com/TechLoCo/sample-go-test/usecase/repository"
)

type userGet struct {
	userRepo repository.User
}

func NewUserGet(userRepo repository.User) UserGet {
	return &userGet{
		userRepo: userRepo,
	}
}

// Run IDで指定した活動しているユーザーを取得
func (u *userGet) Run(id int) (*model.User, error) {
	user, err := u.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	if !user.IsActive() {
		return nil, nil
	}
	return user, nil
}
