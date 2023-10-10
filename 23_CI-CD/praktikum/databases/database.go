package databases

import (
	"docker/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(conf *configs.Config) *gorm.DB{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USER, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil{
		panic(err)
	}
	return db
}