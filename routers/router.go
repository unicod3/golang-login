package routers

import (
	"beego-docker/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.UserAuthController{},"get:Login;post:LoginHandler")
    beego.Router("/register", &controllers.UserAuthController{},"get:Register")
    beego.Router("/home", &controllers.HomeController{},"get:Index")
  //  beego.Router("/logout", &controllers.LoginController{})
}
