package resolver

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

type Resolver struct {
	Orm     *gorm.DB
	Storage *s3.Client
}
