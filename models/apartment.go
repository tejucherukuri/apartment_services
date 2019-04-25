package models

import (
	"time"
)

type Apartment struct {
	Id            int       `json:"apartment_id"`
	ApartmentName string    `json:"apartment_name"`
	Street        string    `json:"street"`
	Address       string    `json:"address"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code"`
	CreatedDate   time.Time `json:"-"`
	LastModified  time.Time `json:"-"`
	ImageUrl      string    `json:"apartment_img_url"`
}

type ApartmentDetails struct {
	Id           int        `json:"apartment_details_id"`
	Apartment    *Apartment `orm:"rel(fk)"`
	BlockNumber  string     `json:"block_name"`
	VFlatNumber  string     `json:"flat_number"`
	CreatedDate  time.Time  `json:"-"`
	LastModified time.Time  `json:"-"`
}

type ApartmentLabour struct {
	Id           int        `json:"id"`
	Apartment    *Apartment `orm:"rel(fk)"`
	LabourName   string     `json:"name"`
	PhoneNumber  string     `json:"phonenumber"`
	Address      string     `json:"address"`
	LabourType   string     `json:"type"`
	Gender       string     `json:"gender"`
	ImageUrl     string     `json:"img_url"`
	AuthorId     int        `json:"author_id"`
	CreatedDate  time.Time  `json:"-"`
	LastModified time.Time  `json:"-"`
}

type ApartmentMeeting struct {
	Id           int        `json:"id"`
	Apartment    *Apartment `orm:"rel(fk)"`
	Author       *User      `orm:"rel(fk)"`
	MeetingName  string     `json:"meeting_name"`
	MeetingDate  time.Time  `json:"meeting_date"`
	MeetingPlace string     `json:"meeting_place"`
	Instructions string     `json:"instructions"`
	CreatedDate  time.Time  `json:"created_date"`
	LastModified time.Time  `json:"last_modified"`
}

type ApartmentBills struct {
	Id              int        `json:"id"`
	Apartment       *Apartment `orm:"rel(fk)"`
	BillType        string     `json:"bill_type";orm:"null"`
	BillName        string     `json:"bill_name";orm:"null"`
	BillDescription string     `json:"bill_description";orm:"null"`
	BillAmount      int        `json:"bill_amount";orm:"null"`
	BillDate        time.Time  `json:"bill_date"`
	Author          *User      `orm:"rel(fk)"`
	CreatedDate     time.Time  `json:"created_date"`
	LastModified    time.Time  `json:"-"`
}


type ApartmentIssues struct {
	
}
