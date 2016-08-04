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
        this.Redirect("/",302)
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
        this.Data["email"] = result.Email
        this.Data["firstName"] = result.FirstName
        this.Data["lastName"] = result.LastName
        this.Data["company"] = result.Company
        this.Data["categories"] = result.Categories
        this.Data["address1"] = result.Address1
        this.Data["address2"] = result.Address2
        this.Data["phone"] = result.Phone
    }
}
