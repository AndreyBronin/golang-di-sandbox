package main

import (
	"context"
	"github.com/AndreyBronin/golang-di-sandbox/core"
	"log"
)

type Customer struct {
	name        string
	Supermarket core.Buyer `inject:""`
}

func NewCustomer(name string) *Customer {
	return &Customer{}
}

func (c *Customer) Start(ctx context.Context) error {
	return nil
}

func (c *Customer) Stop(ctx context.Context) error {
	return nil
}

func (c *Customer) BuyGoods(shoppingList []core.GoodsType) []core.Object {
	result := make([]core.Object, 0)
	for _, good := range shoppingList {
		thing, err := c.Supermarket.Buy(good)
		if err != nil {
			log.Println("Failed to buy ", good.String())
			continue
		}
		log.Printf("Customer %s have bought a %s", c.name, good.String())
		result = append(result, thing)
	}
	return result
}

func (c *Customer) CreateRandomShoppingList() []core.GoodsType {
	panic("implement me")
}
