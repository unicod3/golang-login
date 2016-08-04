package controllers

import (
	"github.com/astaxie/beego"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
)


type HomeController struct {
	beego.Controller
}


func (this *HomeController) Index() {

    sesid := this.GetSession("sesid")

    if(sesid == nil){
        this.Redirect("/login",302)
        return
    }
    this.TplName = "home/index.tpl"

    url := "172.17.0.2:27017"
    database := "testDB"
    collection := "users"
    session, err := mgo.Dial(url)

    if(err != nil){
       log.Fatal(err)
    }
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)
    c := session.DB(database).C(collection)

    result := User{}
    err = c.Find(bson.M{"email" : sesid}).One(&result)

    if(err != nil){
        this.Redirect("/login",302)
        return
    }else{
        this.Data["Email"] = result.Email
        this.Data["Username"] = result.Username
    }
}
