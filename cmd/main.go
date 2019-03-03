package main

import (
	"context"
	"fmt"
	"github.com/AndreyBronin/golang-di-sandbox/farm"
	"github.com/insolar/component-manager"
	"log"
	"time"
)

func main() {
	fmt.Println("Golang DI sandbox")

	cm := component.NewManager(nil)
	cm.Register(farm.NewMilkFarm())
	cm.Inject()

	ctx := context.Background()
	err := cm.Start(ctx)
	if err != nil {
		log.Fatalln("failed to start components: ", err.Error())
	}

	//cm.Stop(ctx)

	<-time.After(time.Second * 5)
}
