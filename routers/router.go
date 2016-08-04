package routers

import (
	"beego-docker/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //HomeController
    beego.Router("/home", &controllers.HomeController{},"get:Index")

    //UserAuthController
    beego.Router("/login", &controllers.UserAuthController{},"get:Login;post:LoginHandler")
    beego.Router("/logout", &controllers.UserAuthController{},"get:Logout")
    beego.Router("/register", &controllers.UserAuthController{},"get:Register;post:RegisterHandler")
    beego.Router("/profile", &controllers.UserAuthController{},"get:Profile;post:ProfileHandler")
}
