package engine

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MConfig struct {
	User, Password, Host string
	Port                 int
	DBName, Parameters   string
}

func (this *MConfig) DSN() gorm.Dialector {
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		this.User, this.Password, this.Host, this.Port, this.DBName, this.Parameters,
	))
}
