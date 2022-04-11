package routers

import (
	"github.com/astaxie/beego"
	"mybeego/controllers"
)

// 路由
func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/login", &controllers.LoginController{})
}
