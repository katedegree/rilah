package request

import "back/infrastructure"

type loginRequest struct {
	AccountCode string `validate:"required"`
	Password    string `validate:"required"`
}

func NewLoginRequest(AccountCode string, Password string) Request {
	return &loginRequest{
		AccountCode: AccountCode,
		Password:    Password,
	}
}
func (r *loginRequest) Validate(v infrastructure.IValidate) ([]string, bool) {
	rules := map[string]map[string]string{
		"AccountCode": {
			"required": "アカウントコードは必須です",
		},
		"Password": {
			"required": "パスワードは必須です",
		},
	}

	return v.Execute(r, rules)
}
