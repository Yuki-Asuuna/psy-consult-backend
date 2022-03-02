package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var client *gorm.DB

const (
	mysql_server_ip = "150.158.159.26"
	port            = 3306
	db_username     = "root"
	db_password     = "ecnu0006"
	db_name         = "psy_hotline"
	charset         = "utf8mb4"
)

func MysqlInit() error {
	var err error
	//Example:
	//root:%ecnu#0006$@(150.158.159.26:3306)/graduate_exemption?charset=utf8mb4&parseTime=True&loc=Local
	target := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", db_username, db_password, mysql_server_ip, port, db_name, charset)
	client, err = gorm.Open("mysql", target)
	if err != nil {
		return err
	}
	return nil
}

func GetMySQLClient() *gorm.DB {
	return client
}
