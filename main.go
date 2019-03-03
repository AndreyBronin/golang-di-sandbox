package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AndreyBronin/golang-di-sandbox/factory"
	"github.com/AndreyBronin/golang-di-sandbox/farm"
	"github.com/AndreyBronin/golang-di-sandbox/supermarket"
	"github.com/AndreyBronin/golang-di-sandbox/warehouse"
	"github.com/insolar/component-manager"
)

func main() {
	cm := component.NewManager(nil)
	cm.Register(farm.NewFarm(), factory.NewDoorFactory())
	cm.Register(&supermarket.Supermarket{}, &warehouse.Warehouse{})

	cm.Register(NewCustomer("Bob"), NewCustomer("Alice"))
	cm.Inject()

	ctx := context.Background()

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	var waitChannel = make(chan bool)

	go func() {
		sig := <-gracefulStop
		log.Println("caught sig: ", sig)

		log.Println("Stopping application...")
		err := cm.Stop(ctx)
		if err != nil {
			log.Fatalln("Failed to graceful stop components: ", err.Error())
		}
		close(waitChannel)
	}()

	err := cm.Start(ctx)
	if err != nil {
		log.Fatalln("Failed to start components: ", err.Error())
	}

	<-waitChannel
}
