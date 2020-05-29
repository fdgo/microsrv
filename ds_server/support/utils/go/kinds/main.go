package main

import (
	"gitee.com/ha666/golibs"
	"gitee.com/ha666/logs"
	"github.com/gin-gonic/gin"
	backApp "gitlab.hfjy.com/service/user-service/controllers/back/app"
	backDept "gitlab.hfjy.com/service/user-service/controllers/back/dept"
	backPermission "gitlab.hfjy.com/service/user-service/controllers/back/permission"
	backPermissionGroup "gitlab.hfjy.com/service/user-service/controllers/back/permissionGroup"
	backPwd "gitlab.hfjy.com/service/user-service/controllers/back/pwd"
	backRole "gitlab.hfjy.com/service/user-service/controllers/back/role"
	backUser "gitlab.hfjy.com/service/user-service/controllers/back/user"
	openDept "gitlab.hfjy.com/service/user-service/controllers/open/dept"
	openUser "gitlab.hfjy.com/service/user-service/controllers/open/user"
	otherSystem "gitlab.hfjy.com/service/user-service/controllers/other/system"
	ssoCheck "gitlab.hfjy.com/service/user-service/controllers/sso/check"
	ssoLogin "gitlab.hfjy.com/service/user-service/controllers/sso/login"
	ssoLogout "gitlab.hfjy.com/service/user-service/controllers/sso/logout"
	"gitlab.hfjy.com/service/user-service/middleware"
	"gitlab.hfjy.com/service/user-service/us_initials/config"
	"gitlab.hfjy.com/service/user-service/worker"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const VERSION = "2018.1015.1643"

func main() {

	//region 初始化系统设置
	runtime.GOMAXPROCS(runtime.NumCPU())
	//endregion

	//region 启动调度任务
	go worker.DoWork()
	//endregion

	//region 初始化gin
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	//endregion

	//region 添加路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 404, "message": "Page not found"})
	})

	r.Any("/health", otherSystem.Time)
	api := r.Group("/api")
	api.Use(middleware.LimitHandler())
	{

		//region 统一登录api
		sso := api.Group("sso")
		sso.POST("/login", ssoLogin.Login)
		sso.POST("/preLogin", ssoLogin.PreLogin)
		sso.POST("/logout", ssoLogout.Logout)
		sso.POST("/check", ssoCheck.Check)
		//endregion

		//region 开放接口api
		open := api.Group("open").Use(middleware.UOPMiddleware())

		// 用户信息
		{
			open.POST("/user/info", openUser.Info)
		}

		// 部门信息
		{
			open.POST("/dept/subList", openDept.SubList)
			open.POST("/dept/info", openDept.Info)
			open.POST("/dept/userList", openDept.UserList)
		}

		// 应用信息
		{

		}

		// 角色权限信息
		{

		}
		//endregion

		//region 后台api
		back := api.Group("back").Use(middleware.JWTAuthHandler())

		// 员工
		{
			back.GET("/user/info", backUser.Info)
			back.POST("/user/freeze", backUser.Freeze)
			back.POST("/user/unFreeze", backUser.UnFreeze)
			back.GET("/user/mobile", backUser.GetUserByMobile)
			back.GET("/user/name", backUser.GetUserByName)
			back.GET("/user/jobNumber", backUser.GetUserByJobNumber)
		}

		// 组织架构
		{
			back.GET("/dept/subList", backDept.SubList)
			back.GET("/dept/info", backDept.Info)
			back.GET("/dept/userList", backDept.UserList)
		}

		// 密码
		{
			back.POST("/pwd/gen", backPwd.Gen)
			back.POST("/pwd/modify", backPwd.Modify)
			back.POST("/pwd/reset", backPwd.Reset)
		}

		// 应用
		{
			back.GET("/app/list", backApp.List)
			back.GET("/app/info", backApp.Info)
			back.GET("/app/secret", backApp.Secret)
			back.POST("/app/create", backApp.Create)
			back.POST("/app/resetSecret", backApp.ResetSecret)
			back.POST("/app/delete", backApp.Delete)
			back.POST("/app/update", backApp.Update)
			back.GET("/app/userList", backApp.UserList)
			back.POST("/app/userRoles", backApp.UserRoles)
			back.GET("/app/permissions", backApp.Permissions)
			back.GET("/app/permissionGroups", backApp.PermissionGroups)
			back.GET("/app/roles", backApp.Roles)
			back.GET("/app/defaultRole", backApp.DefaultRole)
		}

		// 角色
		{
			back.GET("/role/info", backRole.Info)
			back.GET("/role/infos", backRole.Infos)
			back.POST("/role/create", backRole.Create)
			back.POST("/role/edit", backRole.Edit)
			back.POST("/role/delete", backRole.Delete)
		}

		// 权限
		{
			back.GET("/permission/info", backPermission.Info)
			back.GET("/permission/infos", backPermission.Infos)
			back.POST("/permission/create", backPermission.Create)
			back.POST("/permission/edit", backPermission.Edit)
			back.POST("/permission/delete", backPermission.Delete)
			back.POST("/permission/setPermissionGroup", backPermission.SetPermissionGroup)
		}

		// 权限组
		{
			back.GET("/permissionGroup/info", backPermissionGroup.Info)
			back.GET("/permissionGroup/infos", backPermissionGroup.Infos)
			back.POST("/permissionGroup/create", backPermissionGroup.Create)
			back.POST("/permissionGroup/edit", backPermissionGroup.Edit)
			back.POST("/permissionGroup/delete", backPermissionGroup.Delete)
		}

		//endregion

	}
	//endregion

	//region 输出当前系统信息

	//region 输出golang版本信息
	logs.Info("【go】version:%s", runtime.Version())
	//endregion

	//region 输出系统信息
	logs.Info("【sys】os:%s", runtime.GOOS)
	logs.Info("【sys】cpu:%d", runtime.NumCPU())
	//endregion

	//region 输出网络信息
	logs.Info("【net】ip:%s", golibs.GetCurrentIntranetIP())
	//endregion

	//region 输出应用信息
	logs.Info("【app】path:%s", getCurrentDirectory())
	logs.Info("【app】version:%s", VERSION)
	//endregion

	//endregion

	//region 启动api服务
	r.Run(config.Config.Listen.Address)
	//endregion

}

// 获取程序运行路径
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
