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

    /*
    user := User{}
    user.Id = bson.NewObjectId()
    user.Username = "testGO"
    user.Password = "23123"
    user.Email = "testEmail"


    err = c.Insert(&user)
   // err = c.Insert(bson.M{"_id": bson.NewObjectId(), "username":"test", "password": "test11", "email":"123123@123.com"})
    if(err != nil){
        panic(err)
    }

        this.Data["userEmail"] = "testl"
    this.Data["username"] = "username"
    this.Data["password"] = "password"

    */
}
