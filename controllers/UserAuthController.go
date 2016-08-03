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

func (c *UserAuthController) Get() {
	c.TplName = "userAuth/login.tpl"
}


func (this *UserAuthController) Post() {

    username := this.GetString("Username")
//    password := this.GetString("password")

    this.TplName = "userAuth/registration.tpl"
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
    err = c.Find(bson.M{"username" : username}).One(&result)
    if(err != nil){
        panic(err)
    }
    this.Data["userEmail"] = result.Email
    this.Data["username"] = result.Username
    this.Data["password"] = result.Password


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
