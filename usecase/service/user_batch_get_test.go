package service

import (
	"errors"
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

	return []userBatchGetCase{
		{
			name: "[正常系] アクティブ・非アクティブ",
			input: userBatchGetInput{
				ids: []int{activeUser.ID, nonActiveUser.ID},
			},
			expect: userBatchGetOutput{
				users: []*model.User{&activeUser},
				err:   nil,
			},
			userGetExpect: []mock.UserGetIO{active, nonActive},
		},
		{
			name: "[異常系] Mockでエラー発生",
			input: userBatchGetInput{
				ids: []int{abnormalUser.ID},
			},
			expect: userBatchGetOutput{
				users: nil,
				err:   errors.New(""),
			},
			userGetExpect: []mock.UserGetIO{abnormal},
		},
	}
}
