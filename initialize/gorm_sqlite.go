package initialize

import (
	"github.com/glebarez/sqlite"
	"github.com/lmb1113/qh-gin-api/config"
	"github.com/lmb1113/qh-gin-api/global"
	"github.com/lmb1113/qh-gin-api/initialize/internal"
	"gorm.io/gorm"
)

// GormSqlite 初始化Sqlite数据库
func GormSqlite() *gorm.DB {
	s := global.QGA_CONFIG.Sqlite
	return initSqliteDatabase(s)
}

// GormSqliteByConfig 初始化Sqlite数据库用过传入配置
func GormSqliteByConfig(s config.Sqlite) *gorm.DB {
	return initSqliteDatabase(s)
}

// initSqliteDatabase 初始化Sqlite数据库辅助函数
func initSqliteDatabase(s config.Sqlite) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		sqlDB.Exec("PRAGMA journal_mode=WAL;")
		return db
	}
}
