package main

import (
	"fmt"
	"os"

	"github.com/contetto/user-service/models"
	proto "github.com/contetto/user-service/proto"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

const ServiceName = "user-service"

var service micro.Service
var userModel *models.UserModel

func main() {
	// Create a new service. Optionally include some options here.
	service = micro.NewService(
		micro.Name(ServiceName),
		micro.Metadata(map[string]string{
			"MONGO_URL": "localhost:27017",
			"MONGO_DB":  "test",
		}),

		// Setup some flags. Specify --client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(cli.BoolFlag{
			Name:  "client",
			Usage: "Launch the client",
		}),
	)

	// Initialize db model
	userModel = models.NewUserModel(service)

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) {
			if c.Bool("client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	proto.RegisterUsersHandler(service.Server(), new(Users))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
