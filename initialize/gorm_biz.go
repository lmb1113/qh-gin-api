package initialize

import (
	"github.com/lmb1113/qh-gin-api/global"
)

func bizModel() error {
	db := global.QGA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
