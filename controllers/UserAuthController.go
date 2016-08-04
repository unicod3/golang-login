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
    FirstName      string    `bson:"firstname"`
    LastName      string    `bson:"lastname"`
    Company      string    `bson:"company"`
    Categories      string    `bson:"categories"`
    Address1      string    `bson:"address1"`
    Address2      string    `bson:"address2"`
    Phone      string    `bson:"phone"`
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


/**
   sesid must be exist in the Profile section
*/
func (this *UserAuthController) Profile() {
	sesid := this.GetSession("sesid")
    if(sesid == nil){
        this.Redirect("/",302)
        return
    }
    this.TplName = "userAuth/profile.tpl"

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
        this.Data["firstName"] = result.FirstName
        this.Data["lastName"] = result.LastName
        this.Data["company"] = result.Company
        this.Data["categories"] = result.Categories
        this.Data["address1"] = result.Address1
        this.Data["address2"] = result.Address2
        this.Data["phone"] = result.Phone
    }
}


/** POST METHODS **/
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

func (this *UserAuthController) ProfileHandler() {

    sesid := this.GetSession("sesid")

    if(sesid == nil){
        this.Redirect("/login",302)
        return
    }
    flash := beego.NewFlash()
    firstname := this.GetString("firstName")
    lastname := this.GetString("lastName")
    company := this.GetString("company")
    categories := this.GetString("categories")
    address1 := this.GetString("address1")
    address2 := this.GetString("address2")
    phone := this.GetString("phone")

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

    userData :=bson.M{"$set": bson.M{
        "firstname":firstname,
        "lastname":lastname,
        "company":company,
        "categories":categories,
        "address1":address1,
        "address2":address2,
        "phone":phone,
    }}

    err = c.Update(bson.M{"email": sesid},userData)

    if(err != nil){
        flash.Error("Error Occured: Could not update the record")
        flash.Store(&this.Controller)
        this.Redirect("/",302)
        return
    }else{
        //user logged in set session
        this.Redirect("/home",302)
        return
    }
}
