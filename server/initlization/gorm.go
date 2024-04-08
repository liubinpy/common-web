package initlization

import (
	"common-web/server/global"
	"common-web/server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMysql() *gorm.DB {

	mysqlConfig := mysql.Config{
		DSN:               global.CommonConfig.Mysql.Dsn(),
		DefaultStringSize: 200,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(global.CommonConfig.Mysql.MaxIdleConns))
	sqlDB.SetMaxOpenConns(global.CommonConfig.Mysql.MaxOpenConns)

	db.AutoMigrate(&model.User{})
	return db
}
