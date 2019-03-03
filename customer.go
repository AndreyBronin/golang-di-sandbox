package main

import (
	"context"
	"github.com/AndreyBronin/golang-di-sandbox/core"
	"log"
	"math/rand"
	"time"
)

const (
	maxTimeoutSec      = 10
	maxShoppingListLen = 20
)

type Customer struct {
	name   string
	cancel context.CancelFunc

	Supermarket core.Buyer `inject:""`
}

func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (c *Customer) Start(ctx context.Context) error {
	ctx, c.cancel = context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(randomTimeout()):
				c.BuyGoods(createRandomShoppingList())
			}
		}
	}(ctx)
	return nil
}

func (c *Customer) Stop(ctx context.Context) error {
	c.cancel()
	return nil
}

func (c *Customer) BuyGoods(shoppingList []core.GoodsType) []core.Object {
	result := make([]core.Object, 0)
	for _, good := range shoppingList {
		thing, err := c.Supermarket.Buy(good)
		if err != nil {
			log.Printf("%s: Failed to buy %s", c.name, good.String())
			continue
		}
		log.Printf("Customer %s have bought a %s", c.name, good.String())
		result = append(result, thing)
	}
	return result
}

func createRandomShoppingList() []core.GoodsType {
	var count int
	for {
		count = rand.Intn(maxShoppingListLen)
		if count > 0 {
			break
		}
	}
	result := make([]core.GoodsType, count)
	for i, _ := range result {
		result[i] = core.GoodsType(rand.Intn(int(core.GoodsTypeDoor) + 1))
	}

	return result
}

func randomTimeout() time.Duration {
	return time.Second * time.Duration(rand.Intn(maxTimeoutSec))
}
