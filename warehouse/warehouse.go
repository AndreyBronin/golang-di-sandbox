package warehouse

import (
	"github.com/AndreyBronin/golang-di-sandbox/core"
)

type Warehouse struct {
	//Farm    core.Farmer   `inject:""`
	//Factory core.Producer `inject:""`
}

func (w *Warehouse) PutFood(food core.Product) error {
	//panic("implement me")
	return nil
}

func (w *Warehouse) PutProduct() error {
	panic("implement me")
}

func (w *Warehouse) TakeFood() error {
	panic("implement me")
}

func (w *Warehouse) TakeProduct() error {
	panic("implement me")
}

func (w *Warehouse) DumpExpiredFood() {
	panic("implement me")
}

// order food and products if empty
//
