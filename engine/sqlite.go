package engine

import (
	"github.com/belief428/gorm-engine/tools"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MSqlite struct {
	Path string
	Name string
}

func (this *MSqlite) DSN() gorm.Dialector {
	if isExist, _ := tools.PathExists(this.Path); !isExist {
		_ = tools.MkdirAll(this.Path)
	}
	return sqlite.Open(this.Path + "/" + this.Name)
}
