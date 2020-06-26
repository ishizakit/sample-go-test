package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/TechLoCo/sample-go-test/adapter/mock"
	"github.com/TechLoCo/sample-go-test/model"
)

func TestUserBatchGet(t *testing.T) {
	testCases := userBatchGetCases()

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			// モック作成
			mockUser := mock.NewUser()
			// モックの値を設定
			for _, userGetExpect := range c.userGetExpect {
				mockUser.
					On("Get", userGetExpect.Input.ID).
					Return(userGetExpect.Output.User, userGetExpect.Output.Err)
			}

			// サービス作成
			service := NewUserBatchGet(mockUser)

			// テストしたい関数を実行
			users, err := service.Run(c.input.ids)

			// 戻り値を検証
			assert.ElementsMatch(t, c.expect.users, users)
			if c.expect.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			// モックが呼び出されたことを検証
			mockUser.AssertNumberOfCalls(t, "Get", len(c.userGetExpect))
			for _, userGetExpect := range c.userGetExpect {
				mockUser.AssertCalled(t, "Get", userGetExpect.Input.ID)
			}
		})
	}
}

type userBatchGetCase struct {
	name          string
	input         userBatchGetInput
	expect        userBatchGetOutput
	userGetExpect []mock.UserGetIO
}

func userBatchGetCases() []userBatchGetCase {
	active := &model.User{
		ID:         1,
		Name:       "ishizakit",
		Email:      "example001@example.com",
		LastActive: time.Now(),
	}
	nonActive := &model.User{
		ID:         2,
		Name:       "tishizaki",
		Email:      "example002@example.com",
		LastActive: time.Now().AddDate(-2, 0, 0),
	}
	return []userBatchGetCase{
		{
			name: "正常系",
			input: userBatchGetInput{
				ids: []int{active.ID, nonActive.ID},
			},
			expect: userBatchGetOutput{
				users: []*model.User{active},
				err:   nil,
			},
			userGetExpect: []mock.UserGetIO{
				{
					Input: mock.UserGetInput{
						ID: active.ID,
					},
					Output: mock.UserGetOutput{
						User: active,
						Err:  nil,
					},
				},
				{
					Input: mock.UserGetInput{
						ID: nonActive.ID,
					},
					Output: mock.UserGetOutput{
						User: nonActive,
						Err:  nil,
					},
				},
			},
		},
	}
}
