package models

import (
	"log"

	"github.com/contetto/micro-mongo"
	"github.com/micro/go-micro"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

type UserModel struct {
	session  *mongodb.MongoSession
	UserColl *mgo.Collection
}

func NewUserModel(service micro.Service) *UserModel {
	session, err := mongodb.New(service)
	if err != nil {
		log.Fatal("This service can not run without a mongoDB\n", err)
	}

	return &UserModel{
		session:  session,
		UserColl: session.GetCollection("users"),
	}
}

func (u *UserModel) Get(id string) (User, error) {
	var user User

	err := u.UserColl.FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserModel) GetFromAppID(appID string) (User, error) {
	var user User

	err := u.UserColl.Find(bson.M{"appid": appID}).One(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserModel) Insert(user *User) error {
	err := u.UserColl.Insert(user)
	if err != nil {
		return err
	}

	return nil
}
