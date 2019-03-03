package factory

import "github.com/AndreyBronin/golang-di-sandbox/core"

func NewDoorFactory() core.Factory {
	return &doorFactory{}
}

type doorFactory struct {
}

func (t *doorFactory) ProduceProduct(GoodsType, count uint) core.Product {
	return &door{}
}

type door struct{}

func (t *door) Name() string {
	return "door"
}

func (t *door) Weight() uint32 {
	return 100
}
