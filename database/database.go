package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB global .
var DB *gorm.DB
var err error

func init() {
	fmt.Println("DB Conn")
	DB, err = gorm.Open("mysql", "root:123456@tcp(mariadb:3306)/default_go?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(DB)
	if err != nil {
		log.Fatalln(err)
	}

	err = DB.DB().Ping()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Mysql connect")

}
