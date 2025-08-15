package internal

import (
	"github.com/lmb1113/qh-gin-api/config"
	"github.com/lmb1113/qh-gin-api/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	var general config.GeneralDB
	switch global.QGA_CONFIG.System.DbType {
	case "mysql":
		general = global.QGA_CONFIG.Mysql.GeneralDB
	case "pgsql":
		general = global.QGA_CONFIG.Pgsql.GeneralDB
	case "oracle":
		general = global.QGA_CONFIG.Oracle.GeneralDB
	case "sqlite":
		general = global.QGA_CONFIG.Sqlite.GeneralDB
	case "mssql":
		general = global.QGA_CONFIG.Mssql.GeneralDB
	default:
		general = global.QGA_CONFIG.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
