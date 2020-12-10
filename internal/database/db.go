package database

import (
	"github.com/jinzhu/gorm"
	"github.com/murilosrg/go-pay-me/config"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(config.Config().DB.Driver, config.Config().DB.Addr)

	if err != nil {
		logrus.Fatal(err)
	}

	DB.Callback().Create().Remove("gorm:update_time_stamp")
	DB.Callback().Update().Remove("gorm:update_time_stamp")
}
