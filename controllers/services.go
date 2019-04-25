package controllers

import (
	"apartment_services/auth"
	"apartment_services/models"
	api_session "apartment_services/session"
	"errors"
	"time"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
	//	"github.com/pborman/uuid"
)

var ErrUnauthorizedAPIKey = errors.New("API Key Expired or Does not exist.  Please login again")

type AServiceBaseController struct {
	beego.Controller
	start time.Time
}

func (this *AServiceBaseController) Prepare() {
	this.start = time.Now()
}

type Login struct {
	beego.Controller
}

type LoginResponse struct {
	Id          int       `json:"id"`
	ApiKey      string    `json:"api_key"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	UserType    string    `json:"user_type"`
	LastLogin   time.Time `json:"last_login"`
	ImageUrl    string    `json:"image_url"`
	Phone       string    `json:"phone_number"`
}

func (this *Login) Authorize() {
	beego.Error("wat is this here");
	db := models.NewOrm()
	succ, user, apartmentServiceActive, _ := auth.Login(this.Ctx, db)
	beego.Error("Log it here ,lets check the apartmentservice active here:", apartmentServiceActive, succ)
	if succ {
		ResponseSuccess(&this.Controller, PostLogin(db, &user, &this.Controller))

	} else {
		//		if succ {
		//			inst := &models.Institution{Id: user.InstitutionId}
		//			db.Read(inst)
		//			ResponseError(&this.Controller, "1", inst.InstitutionName+" has deactivated Byndr.  Please contact your college administrator for the details.")
		//			return
		//		}
		ResponseError(&this.Controller, "1", "Your username or password is incorrect.")
	}
}

func PostLogin(db *models.AServiceOrm, user *models.User, controller *beego.Controller) LoginResponse {
	// inst := &models.Institution{Id: user.InstitutionId}
	// db.Read(inst)
	// if user.ProgramTerm == nil {
	// 	user.ProgramTerm = &models.ProgramTerm{}
	// } else {
	// 	db.Read(user.ProgramTerm)
	// }
	newApiKey := string(uuid.NewUUID().String())
	sessionCookie := controller.Ctx.GetCookie("beegosessionID")
	api_session.Set(newApiKey, sessionCookie, user.Id, db)
	source := "web"
	beego.Error("what is the source here:", source)
	deviceId := controller.GetString("deviceId")
	appVersion := controller.GetString("appVersion")
	if deviceId != "" {
		auth.SaveDeviceId(deviceId, newApiKey, appVersion, user, db)
		source = "app"
	}
	user.LastLogin = time.Now()
	db.UpdateB(user)
	//	log := &models.UserActivityLog{
	//		UserId:       user.Id,
	//		Source:       source,
	//		Activity:     "login",
	//		ActivityTime: time.Now(),
	//	}
	//	beego.Debug(db.InsertB(log))
	return LoginResponse{
		Id:          user.Id,
		ApiKey:      newApiKey,
		Email:       user.Email,
		LastLogin:   user.LastLogin,
		ImageUrl:    user.ImageUrl,
		Phone:       user.PrimaryPhone,
		DisplayName: user.DisplayName,
	}

}

func (this *AServiceBaseController) Finish() {
	timeDur := time.Since(this.start)
	devInfo := fmt.Sprintf("| % -10s | % -40s | % -16s | % -10s |", this.Ctx.Request.Method, this.Ctx.Request.URL.Path, timeDur.String(), "success")
	beego.Informational(devInfo)
}

func (this *AServiceBaseController) AServiceAbort(statusCode int, msg string) {
	timeDur := time.Since(this.start)
	devInfo := fmt.Sprintf("| % -10s | % -40s | % -16s | % -10s | %s", this.Ctx.Request.Method, this.Ctx.Request.URL.Path, timeDur.String(), "abort", msg)
	beego.Error(devInfo)
	this.Ctx.ResponseWriter.WriteHeader(statusCode)
	this.CustomAbort(statusCode, msg)
}

func (this *AServiceBaseController) GetUserId(apiKey string, cookie string, deviceId string, db *models.AServiceOrm) (int, error) {
	userId := api_session.Get(apiKey, cookie, deviceId, db)
	if userId == 0 {
		return 0, ErrUnauthorizedAPIKey
	}
	return userId, nil
}

func (this *AServiceBaseController) GetProtocol() string {
	if proto, ok := this.Ctx.Request.Header[beego.AppConfig.String("protocol_header")]; ok {
		if len(proto) > 0 {
			return proto[0]
		}
	}
	return ""
}

func (this *AServiceBaseController) GetCookie(name string) string {
	return this.Ctx.GetCookie(name)
}

type PreProcessor interface {
	Abort(string)
	AServiceAbort(int, string)
	GetString(string, ...string) string
	GetInt(string, ...int) (int, error)
	GetCookie(string) string
	GetUserId(apiKey string, cookie string, deviceId string, db *models.AServiceOrm) (int, error)
	GetProtocol() string
}

func PreProcess(this PreProcessor, db *models.AServiceOrm, permissionRequired string,
	fieldsRequired []string) *models.User {
	if this.GetProtocol() == "http" {
		this.AServiceAbort(403, "{\"error\":\"Protocol not allowed\"")
		return nil
	}
	userId, apiError := this.GetUserId(this.GetString("apiKey"), this.GetCookie("beegosessionID"), this.GetString("deviceId"), db)
	if apiError != nil {
		beego.Debug("API Key Not found")
		this.AServiceAbort(401, "{\"error\":\"You session may have expired. Please login again.\"}")
		return nil
	}
	for _, field := range fieldsRequired {
		if this.GetString(field) == "" {
			beego.Debug("Required Field ", field, "Not Found")
			this.AServiceAbort(400, "{\"error\":\"Missing field required.\"}")
			return nil
		}
	}
	user, _ := models.GetUserById(userId, db)
	if &user == nil || (permissionRequired != "") {
		beego.Debug("User Type Not Valid for Operation")
		this.AServiceAbort(403, "{\"error\":\"You are not authorized to perform this action.\"}")
	}
	return &user

}
