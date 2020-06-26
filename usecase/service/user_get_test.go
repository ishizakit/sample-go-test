package service

import (
	"testing"
	"time"

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
			mockUser.AssertNumberOfCalls(t, "Get", 1)
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
	activeUser := model.User{
		ID:         1,
		Name:       "ishizakit",
		Email:      "example001@example.com",
		LastActive: time.Now(),
	}
	active := mock.NewNormalUserGetIO(activeUser)

	nonActiveUser := model.User{
		ID:         2,
		Name:       "tishizaki",
		Email:      "example002@example.com",
		LastActive: time.Now().AddDate(-2, 0, 0),
	}
	nonActive := mock.NewNormalUserGetIO(nonActiveUser)

	abnormalUser := model.User{
		ID:         0,
		Name:       "abnormal",
		Email:      "example000@example.com",
		LastActive: time.Now(),
	}
	abnormal := mock.NewAbnormalUserGetIO(abnormalUser)

	return []userGetCase{
		{
			name: "[正常系] アクティブユーザー",
			input: userGetInput{
				id: activeUser.ID,
			},
			expect: userGetOutput{
				user: &activeUser,
				err:  nil,
			},
			userGetExpect: active,
		},
		{
			name: "[正常系] 非アクティブユーザー",
			input: userGetInput{
				id: nonActiveUser.ID,
			},
			expect: userGetOutput{
				user: nil,
				err:  nil,
			},
			userGetExpect: nonActive,
		},
		{
			name: "[異常系] userMockでエラー発生",
			input: userGetInput{
				id: abnormalUser.ID,
			},
			expect: userGetOutput{
				user: nil,
				err:  abnormal.Output.Err,
			},
			userGetExpect: abnormal,
		},
	}
}
