package producer

import (
	"github.com/AndreyBronin/golang-di-sandbox/core"
	"time"
)

func NewDoorFactory() core.Producer {
	return &doorFactory{}
}

type doorFactory struct {
}

func (d *doorFactory) Produce(t core.GoodsType, count uint) core.Product {
	return &door{}
}

type door struct{}

func (d *door) Name() string {
	return "door"
}

func (d *door) Weight() uint32 {
	return 100
}

func (d *door) Expiration() time.Time {
	return time.Time{}
}
