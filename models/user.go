package models

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id                    int       `json:"user_id"`
	DisplayName           string    `json:"name"`
	Email                 string    `json:"email" orm:"unique"`
	Password              string    `json:"-"`
	ImageUrl              string    `json:"img_url"`
	PrimaryPhone          string    `json:"phone"`
	LastLogin             time.Time `orm:"null" json:"-"`
	AppDownloadedDate     time.Time `orm:"null" json:"-"`
	PushNotificationOptIn bool      `json:"push_optin"`
	EmailOptIn            bool      `json:"email_optin"`
	MessageOptIn          bool      `json:"msg_optin"`
	CreatedDate           time.Time `json:"-"`
	LastModified          time.Time `json:"-"`
	IsActive              bool      `orm:"null" json:"is_active"`
}

type UserDevice struct {
	Id           int `orm:"pk"`
	DeviceId     string
	ApiKey       string
	ApiVersion   string
	User         *User `orm:"rel(fk)"`
	IsActive     bool
	CreatedDate  time.Time
	LastModified time.Time
}

//type Apartment struct {

//}

type ApiSession struct {
	ApiKey        string `orm:"pk"`
	SessionCookie string
	User          *User `orm:"rel(fk)"`
	Expiration    time.Time
	CreatedDate   time.Time
	LastModified  time.Time
}

var ErrUserNameNotFound = errors.New("UserName not found")

//Doesnt exist scenario you still need to mention here
func GetUserByUserName(username string, db *AServiceOrm) (User, error) {
	if db == nil {
		db = NewOrm()
	}
	beego.Error("what is the email:", username)
	user := User{Email: strings.ToLower(username)}
	err := db.Read(&user, "Email")
	if err != nil {
		user.Email = strings.ToLower(username)
		if db.Read(&user, "Email") != nil {
			return user, ErrUserNameNotFound
		}
	}
	return user, nil

}

func GetUserById(userId int, db *AServiceOrm) (User, error) {
	if db == nil {
		db = NewOrm()
	}
	user := User{Id: userId}
	err := db.Read(&user)
	if err == orm.ErrNoRows {
		return user, ErrUserNameNotFound
	}
	return user, nil
}
