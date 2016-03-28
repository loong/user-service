package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/contetto/user-service/models"
	proto "github.com/contetto/user-service/proto"
	"golang.org/x/net/context"
)

/*
   TODO:
     - Properly handle errors form UserModel
*/

type Users struct{}

func (g *Users) Get(ctx context.Context, req *proto.GetReq, rsp *proto.User) error {
	// LH: @todo require auth, maybe even consider to use middleware/ interceptor

	user, err := userModel.Get(req.ID)
	if err != nil {
		log.Println(err)
	}

	// LH: @todo is there a better way to do this assignment?
	rsp.ID = user.ID.Hex()
	rsp.FirstName = user.FirstName
	rsp.LastName = user.LastName
	rsp.Email = user.Email

	return nil
}

// Post inserts a new user
func (g *Users) Post(ctx context.Context, req *proto.User, rsp *proto.User) error {
	// LH: @todo require auth, maybe even consider to use middleware/ interceptor

	appID, appSecret := generateAppTokenPair(req)

	var user models.User
	user.AppSecret = appSecret
	user.AppID = appID

	// LH: @todo is there a better way to do this assignment?
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email

	user.Password = user.Password // LH: @todo IMPORTANT use bcrypt here
	user.Role = "user"

	err := userModel.Insert(&user)
	if err != nil {
		log.Println(err)
	}

	// return request with newly created user obj
	user, err = userModel.GetFromAppID(user.AppID)
	if err != nil {
		log.Println(err)
	}

	// LH: @todo is there a better way to do this assignment?
	rsp.ID = user.ID.Hex()
	rsp.FirstName = user.FirstName
	rsp.LastName = user.LastName
	rsp.Email = user.Email

	return nil
}

func generateAppTokenPair(req *proto.User) (appID, appSecret string) {

	// Generate API ID and Secret (from original user-accounts repo)
	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	numRead, err := rand.Read(buff)
	if numRead != len(buff) || err != nil {
		rand.Read(buff)
	}
	buff2 := make([]byte, 8)
	rand.Read(buff2)

	appID = fmt.Sprintf("%x%x@%s", unix32bits, buff2[0:], req.Platform)
	appSecret = fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])

	return
}
