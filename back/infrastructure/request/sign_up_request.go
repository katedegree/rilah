package request

import "back/infrastructure"

type SignUpRequest struct {
	Name        string `validate:"required|max=255"`
	AccountCode string `validate:"required|max=255"`
	Password    string `validate:"required|max=255"`
}

func (r *SignUpRequest) Validate() ([]string, bool) {
	rules := map[string]map[string]string{
		"Name": {
			"required": "名前を入力してください",
			"max":      "名前は255文字以内で入力してください",
		},
		"AccountCode": {
			"required": "アカウントコードを入力してください",
			"max":      "アカウントコードは255文字以内で入力してください",
		},
		"Password": {
			"required": "パスワードを入力してください",
			"max":      "パスワードは255文字以内で入力してください",
		},
	}

	return infrastructure.Validate(r, rules)
}
