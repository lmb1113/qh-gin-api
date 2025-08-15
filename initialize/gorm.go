package initialize

import (
	"os"

	"github.com/lmb1113/qh-gin-api/global"
	"github.com/lmb1113/qh-gin-api/model/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.QGA_CONFIG.System.DbType {
	case "mysql":
		global.QGA_ACTIVE_DBNAME = &global.QGA_CONFIG.Mysql.Dbname
		return GormMysql()
	case "pgsql":
		global.QGA_ACTIVE_DBNAME = &global.QGA_CONFIG.Pgsql.Dbname
		return GormPgSql()
	case "oracle":
		global.QGA_ACTIVE_DBNAME = &global.QGA_CONFIG.Oracle.Dbname
		return GormOracle()
	case "mssql":
		global.QGA_ACTIVE_DBNAME = &global.QGA_CONFIG.Mssql.Dbname
		return GormMssql()
	case "sqlite":
		global.QGA_ACTIVE_DBNAME = &global.QGA_CONFIG.Sqlite.Dbname
		return GormSqlite()
	default:
		global.QGA_ACTIVE_DBNAME = &global.QGA_CONFIG.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.QGA_DB
	err := db.AutoMigrate(
		user.User{},
	)
	if err != nil {
		global.QGA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		global.QGA_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.QGA_LOG.Info("register table success")
}
