// project_tj project main.go
package main

import (
	"apartment_services/models"
	_ "apartment_services/routers"
	api_session "apartment_services/session"
	"fmt"
	"time"

	"github.com/astaxie/beego/session"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var globalSessions *session.Manager

func init() {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// beego.Debug(orm.RegisterDataBase("default", "mysql", "tejaswi:7fcef841-920d-4863-b090-4de2b2c3c1a1@(byndr-read.co7gbazco3kw.us-east-1.rds.amazonaws.com:3306)/byndr?charset=utf8"))
	beego.Debug(orm.RegisterDataBase("default", "mysql",
		beego.AppConfig.String("mysql_user")+":"+
			beego.AppConfig.String("mysql_password")+"@"+
			beego.AppConfig.String("mysql_db_host")+"/"+
			beego.AppConfig.String("mysql_db")+"?charset=utf8&parseTime=True"))
	orm.RegisterModel(new(models.User),
		new(models.ApiSession), new(models.UserDevice), new(models.Apartment), new(models.ApartmentDetails),
		new(models.ApartmentLabour), new(models.ApartmentMeeting), new(models.ApartmentBills))
	//, new(models.Apartment), new(models.ApartmentUser), new(models.ApartmentLabour), new(models.ApartmentLabourReviews),
}

func main() {
	orm.Debug = true
	beego.Error("i am in the main function")
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.BConfig.WebConfig.Session.SessionOn = false

	//beego.SetLevel(beego.LevelDebug)
	api_session.Initialize()
	//	queue.InitializeQueues()
	//	defer queue.QueueConnection.Close()
	//	defer queue.Channel.Close()
	beego.Run(beego.AppConfig.String("app_httpport"))
	beego.Error("i am in the main function2")
	fmt.Println("Hello World!")
}
