package request

import "back/infrastructure"

type LoginRequest struct {
	AccountCode string `validate:"required"`
	Password    string `validate:"required"`
}

func (r *LoginRequest) Validate() ([]string, bool) {
	rules := map[string]map[string]string{
		"AccountCode": {
			"required": "アカウントコードは必須です",
		},
		"Password": {
			"required": "パスワードは必須です",
		},
	}

	return infrastructure.Validate(r, rules)
}
