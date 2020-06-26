package service

import "github.com/TechLoCo/sample-go-test/model"

// UserGet userGetのインターフェース
type UserGet interface {
	Run(id int) (*model.User, error)
}

// userGetInput UserGet.Run()の引数を表すstruct
type userGetInput struct {
	id int
}

// userGetOutput UserGet.Run()の戻り値を表すstruct
type userGetOutput struct {
	user *model.User
	err  error
}
