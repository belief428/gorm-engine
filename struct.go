package gorm_engine

type EngineMode int

const (
	EngineModeForMysql EngineMode = iota + 1
	EngineModeForSqlite
)
