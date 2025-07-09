package request

import "back/infrastructure"

type joinUserRequest struct {
	GroupID uint32 `validate:"required"`
	UserID  uint32 `validate:"required"`
}

func NewJoinUserRequest(GroupID uint32, UserID uint32) Request {
	return &joinUserRequest{
		GroupID: GroupID,
		UserID:  UserID,
	}
}
func (r *joinUserRequest) Validate(v infrastructure.IValidate) ([]string, bool) {
	rules := map[string]map[string]string{
		"GroupID": {
			"required": "グループIDは必須です",
		},
		"UserID": {
			"required": "ユーザーIDは必須です",
		},
	}

	return v.Execute(r, rules)
}
