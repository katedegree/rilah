package request

import (
	"back/infrastructure"
	"io"
)

type updateUserRequest struct {
	Name        string        `validate:"max=255"`
	AccountCode string        `validate:"max=255"`
	Password    string        `validate:"max=255"`
	File        io.ReadSeeker `validate:""`
	ContentType string        `validate:""`
}

func NewUpdateUserRequest(name *string, accountCode *string, password *string, file io.ReadSeeker, contentType string) Request {
	req := &updateUserRequest{}
	if name != nil {
		req.Name = *name
	}
	if accountCode != nil {
		req.AccountCode = *accountCode
	}
	if password != nil {
		req.Password = *password
	}
	if file != nil && contentType != "" {
		req.File = file
		req.ContentType = contentType
	}

	return req
}
func (r *updateUserRequest) Validate(v infrastructure.IValidate) ([]string, bool) {
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

	return v.Execute(r, rules)
}
