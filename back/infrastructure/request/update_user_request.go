package request

import (
	"back/infrastructure"
	"io"
)

type UpdateUserRequest struct {
	Name        string        `validate:"max=255"`
	AccountCode string        `validate:"max=255"`
	Password    string        `validate:"max=255"`
	File        io.ReadSeeker `validate:""`
	ContentType string        `validate:""`
}

func (r *UpdateUserRequest) Validate() ([]string, bool) {
	rules := map[string]map[string]string{
		"Name": {
			"max": "名前は255文字以内で入力してください。",
		},
		"AccountCode": {
			"max": "アカウントコードは255文字以内で入力してください。",
		},
		"Password": {
			"max": "パスワードは255文字以内で入力してください。",
		},
	}

	return infrastructure.Validate(r, rules)
}
