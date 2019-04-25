package routers

import (
	"apartment_services/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/afe/user", &controllers.Login{}, "post:Authorize")
	beego.Router("/afe/usersignup", &controllers.User{}, "post:UserSignUp")
	beego.Router("/afe/apartmentsearch", &controllers.Apartment{}, "get:SearchApartments")
	beego.Router("/afe/fetchapartmentdetails",&controllers.Apartment{},"get:FetchApartmentDetails")
	beego.Router("/afe/addLocalService", &controllers.Apartment{},"post:AddLocalService")
	beego.Router("/afe/GetLocalService",&controllers.Apartment{},"get:GetLocalService")
	beego.Router("/afe/setupMeeting",&controllers.Apartment{},"post:SetupMeeting")
	beego.Router("/afe/addfinancialbills",&controllers.Apartment{},"post:AddFinancialBillAdmin")
	//	beego.
}
