package request

import "back/infrastructure"

type deleteGroupRequest struct {
	GroupID uint32 `validate:"required"`
}

func NewDeleteGroupRequest(GroupID uint32) Request {
	return &deleteGroupRequest{
		GroupID: GroupID,
	}
}
func (r *deleteGroupRequest) Validate(v infrastructure.IValidate) ([]string, bool) {
	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDを入力してください。",
		},
	}

	return v.Execute(r, rules)
}
