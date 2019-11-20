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

		adminGroup := admingroups.AdminGroup{
			Name: "最高權限使用者",
			Permission: `{"admin":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"adminGroup":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"adminLevel":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"adminAccess":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"ipWhitelisting":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"ipSubnetWhitelisting":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"areaBlacklisting":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"menuGroup":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true},
						"menuSetting":{"enable":true,"view":true,"create":true,"edit":true,"copy":true,"delete":true}
					}`,
			Remark: "",
			Enable: 1,
		}
		db.Create(&adminGroup)
	}

	if !db.HasTable(&adminaccesses.AdminAccess{}) {
		db.AutoMigrate(&adminaccesses.AdminAccess{})

		adminAccesses := []adminaccesses.AdminAccess{
			{
				Code:   "enable",
				Name:   "啟用",
				Enable: 1,
			},
			{
				Code:   "view",
				Name:   "檢視",
				Enable: 1,
			},
			{
				Code:   "create",
				Name:   "新增",
				Enable: 1,
			},
			{
				Code:   "edit",
				Name:   "修改",
				Enable: 1,
			},
			{
				Code:   "copy",
				Name:   "複製",
				Enable: 1,
			},
			{
				Code:   "delete",
				Name:   "刪除",
				Enable: 1,
			},
		}

		for _, adminAccess := range adminAccesses {
			db.Create(&adminAccess)
		}
	}

	if !db.HasTable(&administrators.Administrator{}) {
		db.AutoMigrate(&administrators.Administrator{})

		administrator := administrators.Administrator{
			Account:  "admin",
			Password: "qaz123",
			Name:     "系統管理員",
			GroupID:  0,
			LevelID:  0,
			Remark:   "",
			Enable:   1,
		}
		db.Create(&administrator)
	}

	if !db.HasTable(&adminlevels.AdminLevel{}) {
		db.AutoMigrate(&adminlevels.AdminLevel{})

		adminLevels := []adminlevels.AdminLevel{
			{
				Level:  100,
				Name:   "系統管理員",
				Enable: 1,
			},
			{
				Level:  95,
				Name:   "系統開發人員",
				Enable: 1,
			},
			{
				Level:  90,
				Name:   "系統測試人員",
				Enable: 1,
			},
			{
				Level:  80,
				Name:   "最高管理員",
				Enable: 1,
			},
			{
				Level:  70,
				Name:   "資深管理員",
				Enable: 1,
			},
			{
				Level:  60,
				Name:   "一般管理員",
				Enable: 1,
			},
			{
				Level:  10,
				Name:   "測試帳號",
				Enable: 1,
			},
		}

		for _, adminLevel := range adminLevels {
			db.Create(&adminLevel)
		}
	}

	if !db.HasTable(&menugroups.MenuGroup{}) {
		db.AutoMigrate(&menugroups.MenuGroup{})

		menuGroups := []menugroups.MenuGroup{
			{
				Name:    "主選單",
				Sort:    1,
				Enable:  1,
				AdminID: 1,
			},
			{
				Name:    "系統設定",
				Sort:    2,
				Enable:  1,
				AdminID: 1,
			},
			{
				Name:    "系統選單",
				Sort:    3,
				Enable:  1,
				AdminID: 1,
			},
		}

		for _, menuGroup := range menuGroups {
			db.Create(&menuGroup)
		}
	}

	if !db.HasTable(&menusettings.MenuSetting{}) {
		db.AutoMigrate(&menusettings.MenuSetting{})

		menuSettings := []menusettings.MenuSetting{
			{
				Code:    "admin",
				Name:    "帳號管理",
				GroupID: 2,
				Icon:    "user",
				Icolor:  "#00A65A",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    1,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "adminGroup",
				Name:    "群組管理",
				GroupID: 2,
				Icon:    "peoples",
				Icolor:  "#00A65A",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    2,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "adminLevel",
				Name:    "層級管理",
				GroupID: 2,
				Icon:    "signal",
				Icolor:  "#DD4B39",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    3,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "adminAccess",
				Name:    "操作管理",
				GroupID: 2,
				Icon:    "tap",
				Icolor:  "#DD4B39",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    4,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "ipWhitelisting",
				Name:    "IP白名單管理",
				GroupID: 2,
				Icon:    "ip",
				Icolor:  "#00C0EF",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    5,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "ipSubnetWhitelisting",
				Name:    "IP網段白名單管理",
				GroupID: 2,
				Icon:    "ip-subnet",
				Icolor:  "#00C0EF",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    6,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "areaBlacklisting",
				Name:    "地區黑名單管理",
				GroupID: 2,
				Icon:    "international",
				Icolor:  "#00C0EF",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    7,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "menuGroup",
				Name:    "選單群組管理",
				GroupID: 3,
				Icon:    "nested",
				Icolor:  "#F39C12",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    8,
				Enable:  1,
				AdminID: 1,
			},
			{
				Code:    "menuSetting",
				Name:    "選單管理",
				GroupID: 3,
				Icon:    "list",
				Icolor:  "#F39C12",
				Access:  `{"enable","view","create","edit","copy","delete"}`,
				Sort:    9,
				Enable:  1,
				AdminID: 1,
			},
		}

		for _, menuSetting := range menuSettings {
			db.Create(&menuSetting)
		}
	}

	// db.Model(&administrators.Administrator{}).AddForeignKey("admin_id", "administrators(id)", "RESTRICT", "RESTRICT")
	// db.Model(&administrators.Administrator{}).AddForeignKey("group_id", "admin_groups(id)", "RESTRICT", "RESTRICT")
	// db.Model(&administrators.Administrator{}).AddForeignKey("level_id", "admin_levels(id)", "RESTRICT", "RESTRICT")

	// db.Model(&adminaccesses.AdminAccess{}).AddForeignKey("admin_id", "administrators(id)", "RESTRICT", "RESTRICT")

	// db.Model(&admingroups.AdminGroup{}).AddForeignKey("admin_id", "administrators(id)", "RESTRICT", "RESTRICT")

	// db.Model(&adminlevels.AdminLevel{}).AddForeignKey("admin_id", "administrators(id)", "RESTRICT", "RESTRICT")

	// db.Model(&menugroups.MenuGroup{}).AddForeignKey("admin_id", "administrators(id)", "RESTRICT", "RESTRICT")

	// db.Model(&menusettings.MenuSetting{}).AddForeignKey("admin_id", "administrators(id)", "RESTRICT", "RESTRICT")
}
