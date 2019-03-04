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

type Product interface {
	Name() string
	Weight() uint32
	Expiration() time.Time
}

type Producer interface {
	Produce(GoodsType, uint) Product
}

type Warehouse interface {
	Get(GoodsType) (Product, error)
	Put(Product) error

	DumpExpiredFood()
}

type Buyer interface {
	Buy(GoodsType) (Product, error)
}
