package farm

import (
	"context"
	"fmt"
	"github.com/AndreyBronin/golang-di-sandbox/core"
	"log"
	"time"
)

type Farm interface {
	ProduceFood() core.Food
}

func NewMilkFarm() Farm {
	return &milkFarm{}
}

type milkFarm struct {
	//Warehouse warehouse.Warehouse `inject:""`

	cancel context.CancelFunc
}

func (m *milkFarm) ProduceFood() core.Food {
	return &milk{}
}

func (m *milkFarm) Start(ctx context.Context) error {
	ctx, m.cancel = context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
				milk := m.ProduceFood()
				fmt.Println(milk.Name())
				//m.Warehouse.PutFood(milk.Name())
			}
		}
	}(ctx)

	log.Println("Milk farm started.")
	return nil
}

func (m *milkFarm) Stop(ctx context.Context) error {

	m.cancel()

	log.Println("Milk farm stopped.")
	return nil
}

type milk struct{}

func (m *milk) Name() string {
	return "milk"
}

func (m *milk) Weight() uint32 {
	return 1
}

func (m *milk) Expiration() time.Time {
	return time.Now().Add(time.Hour * 24 * 7)
}
