package request

import "back/infrastructure"

type Request interface {
	Validate(v infrastructure.IValidate) ([]string, bool)
}
