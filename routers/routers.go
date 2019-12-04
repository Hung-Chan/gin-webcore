package routers

import (
	"gin-webcore/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter .
func InitRouter() *gin.Engine {

	router := gin.Default()

	// 帳號管理
	admins := router.Group("admins")
	{
		admins.GET("", controllers.AdminsList)
		admins.GET("/groups", controllers.AdminsGroups)
		admins.GET("/levels", controllers.AdminsLevels)

		// 操作api (新增、檢視、修改、複製、刪除)
		admins.POST("", controllers.AdminsCreate)
		admins.GET("/view/:id", controllers.AdminsView)
		admins.PATCH("/:id", controllers.AdminsUpdate)
		admins.PUT("", controllers.AdminsCopy)
		admins.DELETE("/:id", controllers.AdminsDelete)
	}

	// 群組管理
	adminGroups := router.Group("admin-groups")
	{
		adminGroups.GET("", controllers.AdminGroupsList)
		adminGroups.GET("/permission", controllers.AdminGroupsPermission)

		// 操作api (新增、檢視、修改、複製、刪除)
		adminGroups.POST("", controllers.AdminGroupsCreate)
		adminGroups.GET("/view/:id", controllers.AdminGroupsView)
		adminGroups.PATCH("/:id", controllers.AdminGroupsUpdate)
		adminGroups.PUT("", controllers.AdminGroupsCopy)
		adminGroups.DELETE("/:id", controllers.AdminGroupsDelete)
	}

	// 層級管理
	adminLevels := router.Group("admin-levels")
	{
		adminLevels.GET("", controllers.AdminLevelsList)

		// 操作api (新增、檢視、修改、複製、刪除)
		adminLevels.POST("", controllers.AdminLevelsCreate)
		adminLevels.GET("/view/:id", controllers.AdminLevelsView)
		adminLevels.PATCH("/:id", controllers.AdminLevelsUpdate)
		adminLevels.PUT("", controllers.AdminLevelsCopy)
		adminLevels.DELETE("/:id", controllers.AdminLevelsDelete)
	}

	// 操作管理
	adminAccesses := router.Group("admin-accesses")
	{
		adminAccesses.GET("", controllers.AdminAccessesList)

		// 操作api (新增、檢視、修改、複製、刪除)
		adminAccesses.POST("", controllers.AdminAccessesCreate)
		adminAccesses.GET("/view/:id", controllers.AdminAccessesView)
		adminAccesses.PATCH("/:id", controllers.AdminAccessesUpdate)
		adminAccesses.PUT("", controllers.AdminAccessesCopy)
		adminAccesses.DELETE("/:id", controllers.AdminAccessesDelete)
	}

	// 選單群組管理
	menuGroups := router.Group("menu-groups")
	{
		menuGroups.GET("", controllers.MenuGroupsList)

		// 操作api (新增、檢視、修改、複製、刪除)
		menuGroups.POST("", controllers.MenuGroupCreate)
		menuGroups.GET("/view/:id", controllers.MenuGroupView)
		menuGroups.PATCH("/:id", controllers.MenuGroupUpdate)
		menuGroups.PUT("", controllers.MenuGroupsCopy)
		menuGroups.DELETE("/:id", controllers.MenuGroupDelete)
	}

	// 選單管理
	menuSettings := router.Group("menu-settings")
	{
		menuSettings.GET("", controllers.MenuSettingsList)
		menuSettings.GET("/groups", controllers.MenuSettingsGroups)
		menuSettings.GET("/accesses", controllers.MenuSettingsAccesses)

		// 操作api (新增、檢視、修改、複製、刪除)
		menuSettings.POST("", controllers.MenuSettingsCreate)
		menuSettings.GET("/view/:id", controllers.MenuSettingsView)
		menuSettings.PATCH("/:id", controllers.MenuSettingsUpdate)
		menuSettings.PUT("", controllers.MenuSettingsCopy)
		menuSettings.DELETE("/:id", controllers.MenuSettingsDelete)
		menuSettings.PATCH("", controllers.MenuSettingsSort)
	}

	auth := router.Group("auth")
	{
		auth.POST("/login", controllers.Login)
		auth.GET("/info", controllers.Info)
		// 	auth.GET("/sidebarMenu", controllers.SidebarMenu)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"CODE": 200,
		})
	})

	return router
}
