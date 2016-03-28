package main

import (
	"fmt"
	"os"

	"github.com/contetto/go-micro"
	"github.com/micro/cli"
	proto "github.com/micro/go-micro/examples/service/proto"
	"golang.org/x/net/context"
)

/*

Example usage of top level service initialisation

*/

// Setup and the client
func runClient(service micro.Service) {
	// Create new greeter client
	greeter := proto.NewUsersClient("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// Setup some flags. Specify --client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(cli.BoolFlag{
			Name:  "client",
			Usage: "Launch the client",
		}),
	)

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
