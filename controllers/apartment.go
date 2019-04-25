package controllers

import (
	"github.com/astaxie/beego"
	"apartment_services/models"
	"strconv"
	"apartment_services/util"
	"time"

	"github.com/astaxie/beego/orm"
)

type Apartment struct {
	AServiceBaseController
}

func (this *Apartment) SearchApartments() {
	db := models.NewOrm()
	PreProcess(this, db, "", []string{"apt_string"})
	var apartments []*models.Apartment
	qry := db.QueryTable("Apartment")
	cond := orm.NewCondition()
	search := cond.And("ApartmentName__icontains", this.GetString("apt_string"))
	qry = qry.SetCond(search)
	qry.All(&apartments)
	jsonOutput := make(map[string][]map[string]string)
	apartmentList := make([]map[string]string, len(apartments))
	for j, searchapartmentName := range apartments {
		aMap := map[string]string{
			"name":           searchapartmentName.ApartmentName,
			"street":         searchapartmentName.Street,
			"state":          searchapartmentName.State,
			"country":        searchapartmentName.Country,
			"pincode":        searchapartmentName.PostalCode,
			"displayname": searchapartmentName.ApartmentName + "," + searchapartmentName.Street + "-" + searchapartmentName.State,
			"id":             strconv.Itoa(searchapartmentName.Id),
		}
		apartmentList[j] = aMap
	}
	jsonOutput["apartments"] = apartmentList
	this.Data["json"] = &jsonOutput
	this.ServeJSON()
}

func (this *Apartment)  FetchApartmentDetails() {
	db := models.NewOrm()
	PreProcess(this,db, "",[]string{"apt_id"})
	apartmentId, _ := this.GetInt("apt_id")
    var apartmentDetails []*models.ApartmentDetails
	if apartmentId > 0{
		apartment := models.Apartment{
			Id: apartmentId,
		}
		err := db.Read(&apartment)
		if err == nil {
			db.QueryTable("ApartmentDetails").Filter("Apartment__Id",apartment.Id).All(&apartmentDetails)
			
		} else {
			this.AServiceAbort(404, "Apartment you are searching doesn't exist")
		}

	} 
	// jsonOutput := make(map[string][]map[string]string)
	this.Data["json"] = apartmentDetails
	this.ServeJSON()	
}

func (this *Apartment) GetLocalService() {
	db := models.NewOrm()
	user := PreProcess(this, db , "", []string{"apt_id"})
	db.Read(user)
	apartmentId, _ := this.GetInt("apt_id")
	apartment := models.Apartment {
		Id: apartmentId,
	}
	err := db.Read(&apartment)
	if err != nil {
		this.AServiceAbort(501,"Apartment not found")
	}
	var apartmentlabours []*models.ApartmentLabour
	db.QueryTable("ApartmentLabour").Filter("Apartment__Id",apartment.Id).All(&apartmentlabours)
	this.Data["json"] = &apartmentlabours
	this.ServeJSON()
}

func (this *Apartment) AddLocalService() {
	db := models.NewOrm()
	user := PreProcess(this,db,"",[]string{"name","phone_number","address","labour_type","gender","apt_id","user_type"})
	db.Read(user)	
	apartmentId, _ := this.GetInt("apt_id")
	labourAddition := models.ApartmentLabour{}
	apartment := models.Apartment {
		Id: apartmentId,
	}
	err := db.Read(&apartment)
	if err != nil {
		this.AServiceAbort(501, "Apartment is not found in database")
	} else {
        labourAddition = models.ApartmentLabour {
			Gender: this.GetString("gender"),
			LabourName: this.GetString("name"),
			Address: this.GetString("address"),
			LabourType: this.GetString("labour_type"),
			PhoneNumber: this.GetString("phone_number"),
			AuthorId: user.Id,
			Apartment: &apartment,
			// Apar: apartment,
		} 
		_,err = db.InsertB(&labourAddition)
		beego.Error("what is the err:",err)
	}
	this.Data["json"] = &map[string]string{
		"msg": "Added Successfully",
	}
	this.ServeJSON()
}

//This is just for admin - do check permissions here, which you havent yet done
func (this *Apartment) SetupMeeting() {
	db := models.NewOrm()
	user := PreProcess(this,db,"",[]string{"apt_id","meeting_name","meeting_date","meeting_place","user_type"})
	beego.Error("what is the user here:",user)
	apartmentId, _ := this.GetInt("apt_id")
	apartment := models.Apartment{
		Id: apartmentId,
	}
	apartmentMeeting := models.ApartmentMeeting{}
	err := db.Read(&apartment)
	if err != nil {
		this.AServiceAbort(501, "Apartment is not found in database")
	}
	var meetingTime time.Time
	meetingTime, _ = time.Parse(util.Yyyymmddform,this.GetString("meeting_date"))
	beego.Error("what is the meetingTime here:",meetingTime)

	apartmentMeeting = models.ApartmentMeeting {
		Author: user,
		Apartment: &apartment,
		MeetingName: this.GetString("meeting_name"),
		MeetingDate: meetingTime,
		MeetingPlace: this.GetString("meeting_place"),
		Instructions: this.GetString("instruction"),
	}
	id, err := db.InsertB(&apartmentMeeting)
	beego.Error("what is the id and the err:",id, err)
	this.Data["json"] = &map[string]string{
		"msg": "Meeting has been setup",
	}
	this.ServeJSON()

}


func  (this *Apartment) AddFinancialBillAdmin() {
	db := models.NewOrm()
	user := PreProcess(this,db,"",[]string{"bill_type","bill_name","bill_description","bill_amount","bill_date","user_type","apartment_id"})
	beego.Error("LET's check the user here:",user)
	apartmentId,_ := this.GetInt("apartment_id")
	billAmount, _ := this.GetInt("bill_amount")
	apartment := models.Apartment{
		Id: apartmentId,
	}
	err := db.Read(&apartment)
	if err != nil {
		this.AServiceAbort(501, "Apartment doesn't exist in the database")
	}
	var billingDate time.Time
	billingDate,_ = time.Parse(util.Yyyymmddform,this.GetString("bill_date"))
	beego.Error("print the billingDate here:",billingDate)
	addBill := models.ApartmentBills{}
	addBill = models.ApartmentBills {
		Author: user,
		Apartment: &apartment,
		BillName: this.GetString("bill_name"),
		BillType: this.GetString("bill_type"),
		BillDescription: this.GetString("bill_description"),
		BillDate: billingDate,
		BillAmount: billAmount,
	}
	id, err := db.InsertB(&addBill)
	beego.Error("what is the id and the err:",id, err)
	this.Data["json"] = &map[string]string{
		"msg": "Bill has been added successfully",
	}
	this.ServeJSON()    
}

// func (this *Apartment) ApartmentJoinRequests() {
// 	db := models.NewOrm()
// 	user := PreProcess(this,db,"",[]string{"user_type","apt_id"})
// }

func (this *Apartment) SendRequestForApartmentId() {
	db := models.NewOrm()
	user := PreProcess(this,db,"",[]string{"apt_details_id"})
	beego.Error("what is the user here:", user)
	apartmentId,_ := this.GetInt("apartment_id")
	apartment := models.Apartment{
		Id: apartmentId,
	}
	err := db.Read(&apartment)
	if err != nil  {
		// beego.
	}



}



// func (this *Apartment) ApartmentName() {
// 	db := models.NewOrm()
// 	PreProcess(this, db, "", []string{"apt_name","street","area","city","state","pincode","secretary_name","mobile_number"} )

// }
