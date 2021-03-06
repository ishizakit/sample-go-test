package mock

import (
	"errors"

	"github.com/TechLoCo/sample-go-test/model"
	"github.com/stretchr/testify/mock"
)

type User struct {
	mock.Mock
}

// NewUser repositoryではなくadapterを返す
func NewUser() *User {
	return &User{}
}

// Get 引数に対応した戻り値を返す
func (u *User) Get(id int) (*model.User, error) {
	// 引数に対応した戻り値を返す
	arg := u.Called(id)
	// 各戻り値の型を変換して返す
	return arg.Get(0).(*model.User), arg.Error(1)
}

// UserGetIO User.Getのインプット・アウトプット
type UserGetIO struct {
	Input  UserGetInput  // 引数
	Output UserGetOutput // 戻り値
}

// NewNormalUserGetIO model.Userを元に正常な戻り値のUserGetIOを作成する
func NewNormalUserGetIO(user model.User) UserGetIO {
	return UserGetIO{
		Input: UserGetInput{
			user.ID,
		},
		Output: UserGetOutput{
			&user,
			nil,
		},
	}
}

// NewNormalUserGetIO model.Userを元に異常な戻り値のUserGetIOを作成する
func NewAbnormalUserGetIO(user model.User) UserGetIO {
	return UserGetIO{
		Input: UserGetInput{
			user.ID,
		},
		Output: UserGetOutput{
			nil,
			errors.New(""),
		},
	}
}

// UserGetInput User.Getの引数を表したstruct
type UserGetInput struct {
	ID int
}

// UserGetOutput User.Getの戻り値を表したstruct
type UserGetOutput struct {
	User *model.User
	Err  error
}
