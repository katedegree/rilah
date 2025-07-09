package request

import "back/infrastructure"

type updateGroupRequest struct {
	GroupID uint32 `validate:"required|gt=0"`
	Name    string `validate:"required|max=255"`
}

func NewUpdateGroupRequest(GroupID uint32, Name string) Request {
	return &updateGroupRequest{
		GroupID: GroupID,
		Name:    Name,
	}
}
func (r *updateGroupRequest) Validate(v infrastructure.IValidate) ([]string, bool) {
	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDを入力してください。",
			"gt":       "グループIDは0より大きい必要があります",
		},
		"Name": {
			"required": "グループ名を入力してください。",
			"max":      "グループ名は255文字以内で入力してください。",
		},
	}

	return v.Execute(r, rules)
}
