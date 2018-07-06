package main

import (
	"github.com/ServiceComb/go-chassis"
	"github.com/ServiceComb/go-chassis/core/lager"
	_ "github.com/ServiceComb/go-chassis/bootstrap"
	//_ "github.com/ServiceComb/go-chassis/config-center"
	"github.com/asifdxtreme/chassis-apollo-example/client/schema"
)

func main(){
	chassis.RegisterSchema("rest", &schema.RestFulHello{})
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed.", err)
		return
	}
	chassis.Run()
}