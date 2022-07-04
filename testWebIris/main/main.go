package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/boltdb"
	"goWebIris/testWebIris/main/config"
	"goWebIris/testWebIris/main/controller"
	"goWebIris/testWebIris/main/dataSource"
	"goWebIris/testWebIris/main/service"
	"time"
)

func main() {
	app := NewApp()
	//应用app设置
	configation(app)
	//路由设置
	mvcHandle(app)

	config := config.InitConfig()
	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}

func NewApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")
	app.HandleDir("/img", "./static/img")
	//注册视图⽂件
	//app.RegisterView(iris.HTML("./static", ".html"))
	//app.Get("/", func(context context.Context) {
	//	context.View("index.html")
	//})
	return app
}

func configation(app *iris.Application) {
	//	配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	//	错误配置(未找到资源)
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    "not found",
			"data":   iris.Map{},
		})
	})
	//接口报错
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    "interal error",
			"data":   iris.Map{},
		})
	})
}

//MVC 架构模式处理
func mvcHandle(app *iris.Application) {
	//启⽤session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})
	//session绑定db使用
	db, err := boltdb.New("session.db", 0600)
	if err != nil {
		panic(err)
	}
	//程序终端时将数据库关闭
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})
	//session绑定db数据库
	sessManager.UseDatabase(db)
	//创建Mysql数据路操作引擎
	mysqlEngine := dataSource.NewMysqlEngine()
	//管理员模块功能
	adminService := service.NewAdminService(mysqlEngine)
	admin := mvc.New(app.Party("/admin")) //设置路由组
	admin.Register(
		adminService,
		sessManager.Start,
	)
	//通过mvc的Handle⽅法进⾏控制器的指定
	admin.Handle(new(controller.AdminController))

}
