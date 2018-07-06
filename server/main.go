package main

import (
	_ "github.com/ServiceComb/go-chassis/bootstrap"
	_ "github.com/ServiceComb/go-chassis/config-center"

	"github.com/ServiceComb/go-chassis"
	"github.com/asifdxtreme/chassis-apollo-example/server/schema"
	"log"
)

func main() {
	chassis.RegisterSchema("rest", &schema.RestFulHello{})
	if err := chassis.Init(); err != nil {
		log.Print("Chassis Init Failed")
		return
	}
	chassis.Run()
}
