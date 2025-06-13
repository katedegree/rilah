package request

import "back/infrastructure"

type DeleteGroupRequest struct {
	GroupID uint32 `validate:"required"`
}

func (r *DeleteGroupRequest) Validate() ([]string, bool) {
	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDを入力してください。",
		},
	}

	return infrastructure.Validate(r, rules)
}
