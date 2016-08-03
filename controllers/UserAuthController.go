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
	c.TplName = "userAuth/login.tpl"
}


func (c *UserAuthController) Register() {
	c.TplName = "userAuth/register.tpl"
}


func (this *UserAuthController) LoginHandler() {

    username := this.GetString("Username")
    password := this.GetString("Password")

    this.TplName = "welcome/index.tpl"
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
    err = c.Find(bson.M{"username" : username, "password": password}).One(&result)

    this.Data["Email"] = ""
    this.Data["Username"] = ""
    this.Data["Error"] = ""

    if(err != nil){
        log.Fatal(err)
        this.Data["Error"] = err
    }else{
        this.Data["Email"] = result.Email
        this.Data["Username"] = result.Username
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
