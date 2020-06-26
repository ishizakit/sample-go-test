package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TechLoCo/sample-go-test/adapter/mock"
	"github.com/TechLoCo/sample-go-test/model"
)

func TestUserGet(t *testing.T) {
	testCases := userGetCases()

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			// モック作成
			mockUser := mock.NewUser()
			// モックの値を設定
			mockUser.
				On("Get", c.userGetExpect.Input.ID).
				Return(c.userGetExpect.Output.User, c.userGetExpect.Output.Err)

			// サービス作成
			service := NewUserGet(mockUser)

			// テストしたい関数を実行
			user, err := service.Run(c.input.id)

			// 戻り値を検証
			assert.Equal(t, c.expect.user, user)
			if c.expect.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			// モックが呼び出されたことを検証
			mockUser.AssertCalled(t, "Get", c.userGetExpect.Input.ID)
		})
	}
}

type userGetCase struct {
	name          string
	input         userGetInput
	expect        userGetOutput
	userGetExpect mock.UserGetIO
}

func userGetCases() []userGetCase {
	normal := &model.User{
		ID:    1,
		Name:  "ishizakit",
		Email: "example001@example.com",
	}
	return []userGetCase{
		{
			name: "正常系",
			input: userGetInput{
				id: normal.ID,
			},
			expect: userGetOutput{
				user: normal,
				err:  nil,
			},
			userGetExpect: mock.UserGetIO{
				Input: mock.UserGetInput{
					ID: normal.ID,
				},
				Output: mock.UserGetOutput{
					User: normal,
					Err:  nil,
				},
			},
		},
	}
}
