package service

import (
	"github.com/TechLoCo/sample-go-test/model"
	"github.com/TechLoCo/sample-go-test/usecase/repository"
)

type userBatchGet struct {
	userRepo repository.User
}

func NewUserBatchGet(userRepo repository.User) UserBatchGet {
	return &userBatchGet{
		userRepo: userRepo,
	}
}

// Run IDで指定した活動しているユーザーを複数取得
func (u *userBatchGet) Run(ids []int) ([]*model.User, error) {
	var users []*model.User
	for _, id := range ids {
		user, err := u.userRepo.Get(id)
		if err != nil {
			return nil, err
		}

		if !user.IsActive() {
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
