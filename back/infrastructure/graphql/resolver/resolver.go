package resolver

import (
	"back/domain/entity"
	"back/infrastructure"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

type Resolver struct {
	Orm                *gorm.DB
	Storage            *s3.Client
	Validator          infrastructure.IValidate
	AuthUserContext    infrastructure.IContext[*entity.UserEntity]
	HttpRequestContext infrastructure.IContext[*http.Request]
}
