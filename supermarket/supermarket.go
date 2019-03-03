package supermarket

import (
	"errors"
	"github.com/AndreyBronin/golang-di-sandbox/core"
)

type Supermarket struct {
	Warehouse core.Warehouse `inject:""`
}

func (s *Supermarket) Buy(thing core.GoodsType) (core.Object, error) {
	return nil, errors.New("not implemented")

	//s.Warehouse.
}
