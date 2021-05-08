package engine

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	Path string
	Name string
}

func (this *Sqlite) DSN() gorm.Dialector {
	return sqlite.Open(this.Path + "/" + this.Name)
}
