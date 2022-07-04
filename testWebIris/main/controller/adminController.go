package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"goWebIris/testWebIris/main/model"
	"goWebIris/testWebIris/main/service"
)

//我们使⽤mvc包模式来进⾏功能开发，在进⾏了结构体定义以后，我们接着定义控制器。控制器负责来完成我们请求的逻辑流程控制，是我
//们功能开发的核⼼枢纽。在AdminController定义中，包含iris.Context上下⽂处理对象，⽤于数据功能处理的管理员模块功能实现
//AdminService，还有⽤于session管理的对象。定义PostLogin⽅法来处理⽤户登陆请求。

type AdminController struct {
	//Iris框架自动为每个请求都绑定上上下文对象，可作为接收参数
	Ctx iris.Context
	//引入service接口
	Service service.Adminservice
	//引入session对象
	Session *sessions.Session
}

//定义常量
const (
	ADMIN = "admin"
)

//定义请求登录接口的变量结构体

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**

定义接口：管理员登录功能：json
接口：/admin/login
*/

func (ac *AdminController) PostLogin(context iris.Context) mvc.Result {
	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)
	//登录参数校验
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "⽤户名或密码为空,请重新填写后尝试登录 ",
			},
		}
	}
	var admin model.Admin
	var exist bool
	//根据用户名密码，查询数据库
	admin, exist = ac.Service.GetByAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)
	fmt.Println(admin)
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "1",
				"success": "登录失败",
				"message": "⽤户名或者密码错误,请重新登录",
			},
		}
	}
	format := admin.CreateTime.Format("2006-01-02 15:04:05")
	fmt.Println(format)
	//parse, _ := time.Parse("2006-01-02 15:04:05", format)
	//admin.CreateTime = parse
	//管理员存在设置session【对象序列化】
	userByte, err := json.Marshal(admin)
	if err != nil {
		return mvc.Response{}
	}
	//将用户信息存入缓存
	ac.Session.Set(ADMIN, userByte)
	return mvc.Response{
		//Object: map[string]interface{}{
		//	"status":  "1",
		//	"success": "登录成功",
		//	"message": "管理员登录成功",
		//},
		Object: "admin",
		//Object: admin,
	}

}
