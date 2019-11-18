package database

import (
	"fmt"
	"gin-webcore/migrations/adminaccesses"
	"gin-webcore/migrations/admingroups"
	"gin-webcore/migrations/administrators"
	"gin-webcore/migrations/adminlevels"
	"gin-webcore/migrations/menugroups"
	"gin-webcore/migrations/menusettings"
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
	runMigrate(DB)
}

// runMigrate .
func runMigrate(db *gorm.DB) {
	if !db.HasTable(&admingroups.AdminGroup{}) {
		db.AutoMigrate(&admingroups.AdminGroup{})
	}

	if !db.HasTable(&adminaccesses.AdminAccess{}) {
		db.AutoMigrate(&adminaccesses.AdminAccess{})
	}

	if !db.HasTable(&administrators.Administrator{}) {
		db.Debug().AutoMigrate(&administrators.Administrator{})
	}

	if !db.HasTable(&adminlevels.AdminLevel{}) {
		db.AutoMigrate(&adminlevels.AdminLevel{})
	}

	if !db.HasTable(&menugroups.MenuGroup{}) {
		db.AutoMigrate(&menugroups.MenuGroup{})
	}

	if !db.HasTable(&menusettings.MenuSetting{}) {
		db.AutoMigrate(&menusettings.MenuSetting{})
	}
}
