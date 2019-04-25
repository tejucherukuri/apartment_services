package session

import (
	"apartment_services/models"
	"time"

	"github.com/astaxie/beego"
)

type data struct {
	UserId     int
	Expiration time.Time
}

var apiSessions map[string]data

func Get(key string, cookie string, deviceId string, db *models.AServiceOrm) int {
	beego.Debug("API Key Requested ", key)
	var d data
	found := false
	if key != "" {
		d, found = apiSessions[key]
	} else if cookie != "" {
		d, found = apiSessions[cookie]
	}

	if !found {
		beego.Debug("API Key Not Found in memory - checking db")
		if db == nil {
			db = models.NewOrm()
		}
		var err error
		var apiSession models.ApiSession
		if key != "" {
			apiSession = models.ApiSession{ApiKey: key}
			err = db.Read(&apiSession)
		} else if cookie != "" {
			apiSession = models.ApiSession{SessionCookie: cookie}
			err = db.Read(&apiSession, "SessionCookie")
		} else {
			return 0
		}
		if err != nil {
			beego.Debug(err)
		} else {
			d = data{
				UserId:     apiSession.User.Id,
				Expiration: apiSession.Expiration,
			}
			apiSessions[key] = d
			if apiSession.SessionCookie != "" {
				apiSessions[apiSession.SessionCookie] = d
			}
		}
	}
	if !found && deviceId != "" {
		var userDevice models.UserDevice
		err := db.QueryTable("user_device").Filter("api_key", key).
			Filter("device_id", deviceId).Filter("is_active", true).One(&userDevice)
		if err == nil {
			user := &models.User{Id: userDevice.User.Id}
			db.Read(user)
			if userDevice.LastModified.Add(24 * time.Hour).Before(time.Now()) {
				db.UpdateB(&userDevice)
			}
			d = data{
				UserId:     userDevice.User.Id,
				Expiration: time.Now().Add(time.Hour * 8),
			}
			apiSessions[key] = d
		}
	}
	if time.Now().After(d.Expiration) {
		beego.Debug("API Key Expired ", time.Now(), d.Expiration)
		Remove(key, cookie, db)
		if deviceId == "" {
			return 0
		}
	}
	return d.UserId
}

func Set(key string, cookie string, userId int, db *models.AServiceOrm) {
	sessionTimeout, _ := beego.AppConfig.Int("api_session_timeout_hours")
	expiration := time.Now().Add(time.Hour * time.Duration(sessionTimeout))
	d := data{
		UserId:     userId,
		Expiration: expiration,
	}
	if key != "" {
		apiSessions[key] = d
	}
	if cookie != "" {
		apiSessions[cookie] = d
	}
	apiSession := models.ApiSession{
		ApiKey:        key,
		SessionCookie: cookie,
		User:          &models.User{Id: userId},
		Expiration:    expiration,
	}
	if db == nil {
		db = models.NewOrm()
	}
	db.InsertB(&apiSession)
	beego.Debug("API Key Added ", key, userId)

}

func Remove(key string, cookie string, db *models.AServiceOrm) {
	delete(apiSessions, key)
	delete(apiSessions, cookie)
	if db == nil {
		db = models.NewOrm()
	}
}

func Initialize() {
	apiSessions = make(map[string]data)
}
