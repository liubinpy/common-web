package initlization

import (
	"common-web/server/global"
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

	return db
}
