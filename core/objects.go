package core

import (
	"time"
)

type GoodsType int

//go:generate stringer -type=GoodsType
const (
	// GoodsTypeMilk calls method and returns request
	GoodsTypeMilk GoodsType = iota
	GoodsTypeDoor
)

type Object interface {
	Name() string
	Weight() uint32
}

type Food interface {
	Object
	Expiration() time.Time
}

type Product interface {
	Object
}

type Farmer interface {
	// ProduceFood
	ProduceFood(GoodsType, uint) Food
}

type Factory interface {
	ProduceProduct(GoodsType, count uint) Product
}

type Warehouse interface {
	PutFood(food Food) error
	PutProduct() error

	TakeFood() error
	TakeProduct() error

	DumpExpiredFood()
}

type Buyer interface {
	Buy(GoodsType) (Object, error)
}
