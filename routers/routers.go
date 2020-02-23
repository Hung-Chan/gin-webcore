package routers

import (
	"gin-webcore/controllers"
	"gin-webcore/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter .
func InitRouter() *gin.Engine {

	router := gin.Default()

	// CORS .
	router.Use(middleware.CORS())

	router.POST("/auth/login", controllers.Login)

	router.Use(middleware.Jwt())
	{
		auth := router.Group("auth")
		{
			auth.GET("/info", controllers.Info)
			auth.GET("/sidebarMenu", controllers.SidebarMenu)
			auth.POST("/logout", controllers.Logout)
		}

		// 帳號管理
		admins := router.Group("admins")
		{
			var administratorControllers controllers.AdministratorController

			admins.GET("", administratorControllers.AdministratorsList)
			admins.GET("/groups", administratorControllers.AdministratorGroups)
			admins.GET("/levels", administratorControllers.AdministratorLevels)
			admins.GET("/group-permission/:id", administratorControllers.AdministratorGroupPermission)

			// 操作api (新增、檢視、修改、複製、刪除)
			admins.POST("", administratorControllers.AdministratorCreate)
			admins.GET("/view/:id", administratorControllers.AdministratorView)
			admins.PATCH("/:id", administratorControllers.AdministratorUpdate)
			admins.PUT("", administratorControllers.AdministratorCopy)
			admins.DELETE("/:id", administratorControllers.AdministratorDelete)
		}

		// 群組管理
		adminGroups := router.Group("admin-groups")
		{
			var adminGroupControllers controllers.AdminGroupController

			adminGroups.GET("", adminGroupControllers.AdminGroupsList)
			adminGroups.GET("/permission", adminGroupControllers.AdminGroupsPermission)

			// 操作api (新增、檢視、修改、複製、刪除)
			adminGroups.POST("", adminGroupControllers.AdminGroupCreate)
			adminGroups.GET("/view/:id", adminGroupControllers.AdminGroupView)
			adminGroups.PATCH("/:id", adminGroupControllers.AdminGroupUpdate)
			adminGroups.PUT("", adminGroupControllers.AdminGroupCopy)
			adminGroups.DELETE("/:id", adminGroupControllers.AdminGroupDelete)
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

		// 地區黑名單管理
		areaBlacklistings := router.Group("area-blacklistings")
		{
			var areaBlacklistingControllers controllers.AreaBlacklistingController

			areaBlacklistings.GET("", areaBlacklistingControllers.AreaBlacklistingsList)

			// 操作api (新增、檢視、修改、複製、刪除)
			areaBlacklistings.POST("", areaBlacklistingControllers.AreaBlacklistingCreate)
			areaBlacklistings.GET("/view/:id", areaBlacklistingControllers.AreaBlacklistingView)
			areaBlacklistings.PATCH("/:id", areaBlacklistingControllers.AreaBlacklistingUpdate)
			areaBlacklistings.PUT("", areaBlacklistingControllers.AreaBlacklistingCopy)
			areaBlacklistings.DELETE("/:id", areaBlacklistingControllers.AreaBlacklistingDelete)
		}

		// 選單群組管理
		menuGroups := router.Group("menu-groups")
		{
			var menuGroupControllers controllers.MenuGroupController

			menuGroups.GET("", menuGroupControllers.MenuGroupsList)

			// 操作api (新增、檢視、修改、複製、刪除)
			menuGroups.POST("", menuGroupControllers.MenuGroupCreate)
			menuGroups.GET("/view/:id", menuGroupControllers.MenuGroupView)
			menuGroups.PATCH("/:id", menuGroupControllers.MenuGroupUpdate)
			menuGroups.PUT("", menuGroupControllers.MenuGroupsCopy)
			menuGroups.DELETE("/:id", menuGroupControllers.MenuGroupDelete)
		}

		// 選單管理
		menuSettings := router.Group("menu-settings")
		{
			var menuSettingControllers controllers.MenuSettingController

			menuSettings.GET("", menuSettingControllers.MenuSettingsList)
			menuSettings.GET("/groups", menuSettingControllers.MenuGroupsOption)
			menuSettings.GET("/accesses", menuSettingControllers.MenuAccessesOption)

			// 操作api (新增、檢視、修改、複製、刪除)
			menuSettings.POST("", menuSettingControllers.MenuSettingCreate)
			menuSettings.GET("/view/:id", menuSettingControllers.MenuSettingView)
			menuSettings.PATCH("/:id", menuSettingControllers.MenuSettingUpdate)
			menuSettings.PUT("", menuSettingControllers.MenuSettingCopy)
			menuSettings.DELETE("/:id", menuSettingControllers.MenuSettingDelete)
			menuSettings.PATCH("", menuSettingControllers.MenuSettingsSort)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"CODE": 200,
		})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
