package gorm_engine

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/gorm/schema"

	"github.com/belief428/gorm-engine/logic"
	"gorm.io/gorm"
)

type Engine struct {
	Mode   EngineMode
	Config *EngineConfig
}

type EngineConfig struct {
	Debug                                   bool
	TablePrefix                             string
	Complex                                 bool
	MaxIdleConns, MaxOpenConns, MaxLifetime int
}

type EngineHandle func(kind EngineMode, config *EngineConfig) *Engine

func (this *Engine) Start(engine logic.IEngine) *gorm.DB {
	option := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   this.Config.TablePrefix,
			SingularTable: !this.Config.Complex,
		},
	}
	if this.Config.Debug {
		option.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Silent,
				Colorful:      false,
			},
		)
	}
	db, err := gorm.Open(engine.DSN(), option)

	if err != nil {
		panic("Orm Open Error：" + err.Error())
	}
	_db, _ := db.DB()
	_db.SetMaxIdleConns(this.Config.MaxIdleConns)
	_db.SetMaxOpenConns(this.Config.MaxOpenConns)
	_db.SetConnMaxLifetime(time.Duration(this.Config.MaxLifetime) * time.Second)
	return db
}

func NewEngine() EngineHandle {
	return func(mode EngineMode, config *EngineConfig) *Engine {
		return &Engine{
			Mode: mode, Config: config,
		}
	}
}
