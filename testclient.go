package main

import (
	"fmt"
	"log"

	proto "github.com/contetto/user-service/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

/*
 This file can be deleted later
*/

// Setup and the client
func runClient(service micro.Service) {

	// Create new client
	client := proto.NewUsersClient("user-service", service.Client())

	// Call the client to create new user
	log.Println("Call Users.Post to create a new account")
	postrsp, err := client.Post(context.TODO(), &proto.User{
		FirstName: "Long",
		LastName:  "Hoang",
		Email:     "long@mindworker.de",
		Password:  "contetto",
		Platform:  "Chrome?",
		Role:      "user",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println(postrsp)

	log.Println("New ID: ", postrsp.ID)

	// Call the client
	log.Println("Call Users.Get to retrieve that new account")
	getrsp, err := client.Get(context.TODO(), &proto.GetReq{ID: postrsp.ID})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(getrsp)
}
