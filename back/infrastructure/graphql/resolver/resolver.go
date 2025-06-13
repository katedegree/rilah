package resolver

import "gorm.io/gorm"

type Resolver struct {
	Orm *gorm.DB
}
