package routers

import (
	"beego-docker/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.UserAuthController{},"*:Get")
  //  beego.Router("/logout", &controllers.LoginController{})
  //  beego.Router("/register", &controllers.LoginController{})
}
