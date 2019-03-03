package main

import (
	"context"
	"github.com/AndreyBronin/golang-di-sandbox/factory"
	"github.com/AndreyBronin/golang-di-sandbox/farm"
	"github.com/AndreyBronin/golang-di-sandbox/supermarket"
	"github.com/AndreyBronin/golang-di-sandbox/warehouse"
	"github.com/insolar/component-manager"
	"log"
	"time"
)

func main() {
	cm := component.NewManager(nil)
	cm.Register(farm.NewFarm(), factory.NewDoorFactory())
	cm.Register(&supermarket.Supermarket{}, &warehouse.Warehouse{})

	cm.Register(NewCustomer("Bob"), NewCustomer("Alice"))
	cm.Inject()

	ctx := context.Background()
	err := cm.Start(ctx)
	if err != nil {
		log.Fatalln("failed to start components: ", err.Error())
	}

	//cm.Stop(ctx)

	<-time.After(time.Second * 5)
}
