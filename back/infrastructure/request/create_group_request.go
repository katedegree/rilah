package request

import "back/infrastructure"

type createGroupRequest struct {
	Name string `validate:"required|max=255"`
}

func NewCreateGroupRequest(name string) Request {
	return &createGroupRequest{Name: name}
}
func (r *createGroupRequest) Validate(v infrastructure.IValidate) ([]string, bool) {
	rules := map[string]map[string]string{
		"Name": {
			"required": "グループ名を入力してください。",
			"max":      "グループ名は255文字以内で入力してください。",
		},
	}

	return v.Execute(r, rules)
}
