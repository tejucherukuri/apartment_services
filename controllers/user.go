package controllers

import (
	"apartment_services/models"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	AServiceBaseController
}

func (this *User) GetUserInfo() {
	db := models.NewOrm()
	user := PreProcess(this, db, "", []string{})
	beego.Error("what is the error:", user)
	this.Data["json"] = &user
	this.ServeJSON()
}

func (this *User) ForgotPassword() {
	
}

func (this *User) UserSignUp() {
	db := models.NewOrm()
	//	PreProcess(this, db, "", []string{"first_name", "last_name", "email", "phone_number", "password"})
	Email := this.GetString("email")
	PhoneNumber := this.GetString("phone_number")
	//	LastName := this.GetString("last_name")
	//	FirstName := this.GetString("first_name")
	//	password := this.GetString("password")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(this.GetString("password")), 10)
	iuser := models.User{}
	emailCheck := db.QueryTable("User").Filter("Email", Email).One(&iuser)
	phoneNumber := db.QueryTable("User").Filter("PrimaryPhone", PhoneNumber).One(&iuser)
	if emailCheck != nil && phoneNumber != nil {
		userinsert := models.User{
			Email:        Email,
			PrimaryPhone: PhoneNumber,
			Password:     string(hashedPassword),
			DisplayName:  this.GetString("first_name") + " " + this.GetString("last_name"),
			IsActive:     false,
		}
		db.InsertB(&userinsert)
	} else if emailCheck == nil {
		this.AServiceAbort(502, "Email already exists , please login to continue")
	} else {
		this.AServiceAbort(501, "PhoneNumber already exists, please login to continue")
	}

	// db.QueryTable("user").Filter("Email",e)
	beego.Error("signup function here :")
	this.Data["json"] = &map[string]string{
		"msg": "User signedin successfully",
	}
	this.ServeJSON()
}
