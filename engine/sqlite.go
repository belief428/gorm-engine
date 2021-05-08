package engine

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MSqlite struct {
	Path string
	Name string
}

func (this *MSqlite) DSN() gorm.Dialector {
	return sqlite.Open(this.Path + "/" + this.Name)
}
