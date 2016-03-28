package models

import "gopkg.in/mgo.v2/bson"

// User model defines the user attributes
//swagger:model User
type User struct {
	ID                  bson.ObjectId `json:"_id" valid:"-" bson:"_id,omitempty"`
	FirstName           string        `json:"FirstName" valid:"ascii,required"`
	LastName            string        `json:"LastName" valid:"ascii"`
	Email               string        `json:"Email" valid:"email,required"`
	MobileNumber        string        `json:"MobileNumber" valid:"ascii"`
	PhoneNumber         string        `json:"PhoneNumber" valid:"ascii"`
	ReceiveNotification bool          `json:"ReceiveNotification"`
	Address             string        `json:"Address" valid:"ascii"`
	City                string        `json:"City" valid:"ascii"`
	State               string        `json:"State" valid:"ascii"`
	Country             string        `json:"Country" valid:"ascii"`
	Zip                 string        `json:"Zip" valid:"ascii"`
	Password            string        `json:"Password" valid:"ascii,required"`
	ProfilePhoto        string        `json:"ProfilePhoto" valid:"ascii"`
	PromoCodeID         string        `json:"PromoCodeId" valid:"ascii"`
	JoinMailList        bool          `json:"JoinMailList"`
	Organization        string        `json:"Organization" valid:"ascii"`
	WhitelistHosts      []string      `json:"WhitelistHosts"`
	WhitelistIPs        []string      `json:"WhitelistIPs"`
	SubscriptionStart   float64       `json:"SubscriptionStart" valid:"-"`
	SubscriptionEnd     float64       `json:"SubscriptionEnd" valid:"-"`
	SubscriptionPlan    string        `json:"SubscriptionPlan" valid:"ascii"`
	IsSubscriptionValid bool          `json:"IsSubscriptionValid"`
	Platform            string        `json:"Platform" valid:"ascii,required"`
	AppID               string        `json:"AppID"`
	AppSecret           string        `json:"AppSecret"`
	Role                string        `json:"Role" valid:"ascii,required"`
	IsEnabled           bool          `json:"IsEnabled"`
	IsVerified          bool          `json:"IsVerified"`
}
