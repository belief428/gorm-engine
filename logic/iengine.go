package logic

import "gorm.io/gorm"

type IEngine interface {
	DSN() gorm.Dialector
}
