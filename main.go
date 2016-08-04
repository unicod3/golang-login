package main

import (
	_ "beego-docker/routers"
	"github.com/astaxie/beego"
)

func main() {
    beego.BConfig.WebConfig.TemplateLeft="<<<"
    beego.BConfig.WebConfig.TemplateRight=">>>"
    beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

