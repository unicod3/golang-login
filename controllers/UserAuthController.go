package controllers

import (
	"github.com/astaxie/beego"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
)

type User struct{
    Id         bson.ObjectId `bson:"_id,omitempty"`
    Username   string    `bson:"username"`
    Password   string    `bson:password`
    Email      string    `bson:"email"`
}



type UserAuthController struct {
	beego.Controller
}

func (c *UserAuthController) Login() {
    s := c.GetSession("sesid")
    if s != nil {
        c.Redirect("/home",302)
        return
    }
	c.TplName = "userAuth/login.tpl"
}

func (c *UserAuthController) Logout() {
    s := c.GetSession("sesid")
    if s != nil {
       c.DelSession("sesid")
    }
    c.Redirect("/",302)
        return
}

func (c *UserAuthController) Register() {
	s := c.GetSession("sesid")
    if(s != nil){
        c.Redirect("/home",302)
        return
    }
    c.TplName = "userAuth/register.tpl"
}


func (this *UserAuthController) LoginHandler() {

    flash := beego.NewFlash()
    email := this.GetString("email")
    password := this.GetString("password")

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
    err = c.Find(bson.M{"email" : email, "password":password}).One(&result)

    if(err != nil){
        flash.Error("Username or Password is not correct")
        flash.Store(&this.Controller)
        this.Redirect("/login",302)
        return
    }else{
        this.Data["Email"] = result.Email
        this.Data["Username"] = result.Username
        //user logged in set session
        this.SetSession("sesid", result.Email)
        this.Redirect("/home",302)
    }
}


func (this *UserAuthController) RegisterHandler() {

    flash := beego.NewFlash()
    email := this.GetString("email")
    password := this.GetString("password")

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

    user := User{}
    user.Id = bson.NewObjectId()
    user.Password = password
    user.Email = email
    err = c.Insert(&user)

    if(err != nil){
        flash.Error("Error Occured: Could not insert to database")
        flash.Store(&this.Controller)
        this.Redirect("/register",302)
        return
    }else{
        //user logged in set session
        this.SetSession("sesid", email)
        this.Redirect("/profile",302)
    }
}
