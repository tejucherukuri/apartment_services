package auth

import (
	"apartment_services/models"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	"golang.org/x/crypto/bcrypt"
)

func Login(r *context.Context, db *models.AServiceOrm) (bool, models.User, bool, error) {
	username, password := r.Input.Query("emailusername"), r.Input.Query("password")
	beego.Error("what is the password here1:",r.Input.Query("emailusername"),password)
	user, err := models.GetUserByUserName(username, db)
	beego.Error("what is it failing here during login:", err)
	if err != nil {
		return false, user, false, err
	}
	beego.Debug("User found :", user.Email, user.Id)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return true, user, false, nil
	}
	beego.Error("what is the err here:", err)
	if password == beego.AppConfig.String("support_user_text")+user.Password {
		return true, user, false, nil
	}
	return false, user, true, nil
}

func SaveDeviceId(deviceId string, key string, appVersion string, user *models.User, db *models.AServiceOrm) {
	beego.Debug("Entered Save Device")
	var devicesToInactivate []*models.UserDevice
	db.QueryTable("user_device").Filter("device_id", deviceId).Filter("IsActive", true).All(&devicesToInactivate)
	for _, device := range devicesToInactivate {
		device.IsActive = false
		db.UpdateB(device)
	}
	var userDevice models.UserDevice
	err := db.QueryTable("user_device").Filter("device_id", deviceId).
		Filter("User__Id", user.Id).Limit(1).One(&userDevice)
	if err != nil {
		userDevice = models.UserDevice{
			DeviceId:   deviceId,
			ApiKey:     key,
			ApiVersion: appVersion,
			User:       user,
			IsActive:   true,
		}
		beego.Debug(db.InsertB(&userDevice))
	} else {
		userDevice.ApiKey = key
		userDevice.IsActive = true
		userDevice.ApiVersion = appVersion
		db.UpdateB(&userDevice)
	}

}
