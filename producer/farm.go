package producer

import (
	"context"
	"fmt"
	"github.com/AndreyBronin/golang-di-sandbox/core"
	"log"
	"time"
)

func NewFarm() core.Producer {
	return &farm{}
}

type farm struct {
	Warehouse core.Warehouse `inject:""`

	cancel context.CancelFunc
}

func (f *farm) Produce(t core.GoodsType, count uint) core.Product {
	// switch t {
	// case core.GoodsTypeMilk:
	// default:
	// 	panic("The Farm can't produce " + t.String())
	//
	// }
	return &milk{}
}

func (f *farm) Start(ctx context.Context) error {
	ctx, f.cancel = context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
				milk := f.Produce(core.GoodsTypeMilk, 1)
				fmt.Println(milk.Name())
				f.Warehouse.Put(milk)
				// TODO use channels intead
				// f.Warehouse.GetChannel() <- milk
			}
		}
	}(ctx)

	log.Println("Farm started.")
	return nil
}

func (f *farm) Stop(ctx context.Context) error {
	f.cancel()
	log.Println("Farm stopped.")
	return nil
}

type milk struct{}

func (milk) Name() string {
	return "milk"
}

func (milk) Weight() uint32 {
	return 1
}

func (m *milk) Expiration() time.Time {
	return time.Now().Add(time.Hour * 24 * 7)
}

type bread struct{}

func (bread) Weight() uint32 {
	return 1
}

func (bread) Expiration() time.Time {
	return time.Now().Add(time.Hour * 24)
}
