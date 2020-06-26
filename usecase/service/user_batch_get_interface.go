package service

import "github.com/TechLoCo/sample-go-test/model"

// UserBatchGet userBatchGetのインターフェース
type UserBatchGet interface {
	Run(ids []int) ([]*model.User, error)
}

// userBatchGetInput UserBatchGet.Run()の引数を表すstruct
type userBatchGetInput struct {
	ids []int
}

// userBatchGetOutput UserBatchGet.Run()の戻り値を表すstruct
type userBatchGetOutput struct {
	users []*model.User
	err   error
}
