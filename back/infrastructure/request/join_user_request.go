package request

import "back/infrastructure"

type JoinUserRequest struct {
	GroupID uint32 `validate:"required"`
	UserID  uint32 `validate:"required"`
}

func (r *JoinUserRequest) Validate() ([]string, bool) {
	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDは必須です",
		},
		"UserID": {
			"required": "ユーザーIDは必須です",
		},
	}

	return infrastructure.Validate(r, rules)
}
