package routers

import (
	"gin-webcore/middleware"
	"gin-webcore/controllers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// InitRouter .
func InitRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/auth/login", controllers.Login)

	router.Use(middleware.Jwt())
	{
		auth := router.Group("auth")
		{
			auth.GET("/info", controllers.Info)
			auth.GET("/sidebarMenu", controllers.SidebarMenu)
			auth.POST("/logout", controllers.Logout)
		}

		// 層級管理
		adminLevels := router.Group("admin-levels")
		{
			var adminLevelControllers controllers.AdminLevelController

			adminLevels.GET("", adminLevelControllers.AdminLevelsList)

			// 操作api (新增、檢視、修改、複製、刪除)
			adminLevels.POST("", adminLevelControllers.AdminLevelCreate)
			adminLevels.GET("/view/:id", adminLevelControllers.AdminLevelView)
			adminLevels.PATCH("/:id", adminLevelControllers.AdminLevelUpdate)
			adminLevels.PUT("", adminLevelControllers.AdminLevelCopy)
			adminLevels.DELETE("/:id", adminLevelControllers.AdminLevelDelete)
		}

		// 操作管理
		adminAccesses := router.Group("admin-accesses")
		{
			var adminAccessControllers controllers.AdminAccessController

			adminAccesses.GET("", adminAccessControllers.AdminAccessesList)

			// 操作api (新增、檢視、修改、複製、刪除)
			adminAccesses.POST("", adminAccessControllers.AdminAccessCreate)
			adminAccesses.GET("/view/:id", adminAccessControllers.AdminAccessView)
			adminAccesses.PATCH("/:id", adminAccessControllers.AdminAccessUpdate)
			adminAccesses.PUT("", adminAccessControllers.AdminAccessCopy)
			adminAccesses.DELETE("/:id", adminAccessControllers.AdminAccessDelete)
		}

		// IP白名單管理
		ipWhitelistings := router.Group("ip-whitelistings")
		{
			var ipWhitelistingControllers controllers.IPWhitelistingController

			ipWhitelistings.GET("", ipWhitelistingControllers.IPWhitelistingsList)

			// 操作api (新增、檢視、修改、複製、刪除)
			ipWhitelistings.POST("", ipWhitelistingControllers.IPWhitelistingCreate)
			ipWhitelistings.GET("/view/:id", ipWhitelistingControllers.IPWhitelistingView)
			ipWhitelistings.PATCH("/:id", ipWhitelistingControllers.IPWhitelistingUpdate)
			ipWhitelistings.PUT("", ipWhitelistingControllers.IPWhitelistingCopy)
			ipWhitelistings.DELETE("/:id", ipWhitelistingControllers.IPWhitelistingDelete)
		}

		// IP網段白名單管理
		ipSubnetWhitelistings := router.Group("ip-subnet-whitelistings")
		{
			var ipSubnetWhitelistingControllers controllers.IPSubnetWhitelistingController

			ipSubnetWhitelistings.GET("", ipSubnetWhitelistingControllers.IPSubnetWhitelistingsList)

			// 操作api (新增、檢視、修改、複製、刪除)
			ipSubnetWhitelistings.POST("", ipSubnetWhitelistingControllers.IPSubnetWhitelistingCreate)
			ipSubnetWhitelistings.GET("/view/:id", ipSubnetWhitelistingControllers.IPSubnetWhitelistingView)
			ipSubnetWhitelistings.PATCH("/:id", ipSubnetWhitelistingControllers.IPSubnetWhitelistingUpdate)
			ipSubnetWhitelistings.PUT("", ipSubnetWhitelistingControllers.IPSubnetWhitelistingCopy)
			ipSubnetWhitelistings.DELETE("/:id", ipSubnetWhitelistingControllers.IPSubnetWhitelistingDelete)
		}
	}

	// 帳號管理
	admins := router.Group("admins")
	{
		admins.GET("", controllers.AdministratorsList)
		admins.GET("/groups", controllers.AdministratorGroups)
		admins.GET("/levels", controllers.AdministratorLevels)
		admins.GET("/group-permission/:id", controllers.AdministratorGroupPermission)

		// 操作api (新增、檢視、修改、複製、刪除)
		admins.POST("", controllers.AdministratorCreate)
		admins.GET("/view/:id", controllers.AdministratorView)
		admins.PATCH("/:id", controllers.AdministratorUpdate)
		admins.PUT("", controllers.AdministratorCopy)
		admins.DELETE("/:id", controllers.AdministratorDelete)
	}

	// 群組管理
	adminGroups := router.Group("admin-groups")
	{
		adminGroups.GET("", controllers.AdminGroupsList)
		adminGroups.GET("/permission", controllers.AdminGroupsPermission)

		// 操作api (新增、檢視、修改、複製、刪除)
		adminGroups.POST("", controllers.AdminGroupCreate)
		adminGroups.GET("/view/:id", controllers.AdminGroupView)
		adminGroups.PATCH("/:id", controllers.AdminGroupUpdate)
		adminGroups.PUT("", controllers.AdminGroupCopy)
		adminGroups.DELETE("/:id", controllers.AdminGroupDelete)
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
		menuSettings.GET("/groups", controllers.MenuGroupsOption)
		menuSettings.GET("/accesses", controllers.MenuAccessesOption)

		// 操作api (新增、檢視、修改、複製、刪除)
		menuSettings.POST("", controllers.MenuSettingCreate)
		menuSettings.GET("/view/:id", controllers.MenuSettingView)
		menuSettings.PATCH("/:id", controllers.MenuSettingUpdate)
		menuSettings.PUT("", controllers.MenuSettingCopy)
		menuSettings.DELETE("/:id", controllers.MenuSettingDelete)
		menuSettings.PATCH("", controllers.MenuSettingsSort)
	}

	// 地區黑名單管理
	areaBlacklistings := router.Group("area-blacklistings")
	{
		areaBlacklistings.GET("", controllers.AreaBlacklistingsList)

		// 操作api (新增、檢視、修改、複製、刪除)
		areaBlacklistings.POST("", controllers.AreaBlacklistingCreate)
		areaBlacklistings.GET("/view/:id", controllers.AreaBlacklistingView)
		areaBlacklistings.PATCH("/:id", controllers.AreaBlacklistingUpdate)
		areaBlacklistings.PUT("", controllers.AreaBlacklistingCopy)
		areaBlacklistings.DELETE("/:id", controllers.AreaBlacklistingDelete)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"CODE": 200,
		})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
