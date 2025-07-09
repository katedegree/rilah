package infrastructure

import (
	"github.com/go-playground/validator/v10"
)

type IValidate interface {
	Execute(input any, rules map[string]map[string]string) ([]string, bool)
}
type Validate struct{}

func NewValidate() IValidate {
	return &Validate{}
}

func (v *Validate) Execute(input any, rules map[string]map[string]string) ([]string, bool) {
	var vld = validator.New()

	err := vld.Struct(input)
	if err == nil {
		return nil, true
	}

	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{"エラーが発生しました。"}, false
	}

	var msgs []string
	for _, fe := range ves {
		if tags, ok := rules[fe.Field()]; ok {
			if msg, ok := tags[fe.Tag()]; ok {
				msgs = append(msgs, msg)
			}
		}
	}

	return msgs, false
}
