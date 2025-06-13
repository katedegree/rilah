package request

import "back/infrastructure"

type CreateGroupRequest struct {
	Name string `validate:"required|max=255"`
}

func (r *CreateGroupRequest) Validate() ([]string, bool) {
	rules := map[string]map[string]string{
		"Name": {
			"required": "グループ名を入力してください。",
			"max":      "グループ名は255文字以内で入力してください。",
		},
	}

	return infrastructure.Validate(r, rules)
}
