package initialize

import (
	"gorm.io/driver/mysql"
	"os"
	"qh-gin-api/initialize/internal"

	"qh-gin-api/global"

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

func InitDb2() *gorm.DB {
	m := global.QGA_CONFIG.Mysql
	m.Dbname = "play"
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func RegisterTables() {
	db := global.QGA_DB
	err := db.AutoMigrate() // todo 自动注册表
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
