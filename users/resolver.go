package users

import (
	"gorm.io/gorm"
	"open-btm.com/observe"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB     *gorm.DB
	Tracer *observe.RouteTracer
}
