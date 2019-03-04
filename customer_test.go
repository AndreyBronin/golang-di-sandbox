package main

import (
	"context"
	"github.com/AndreyBronin/golang-di-sandbox/core"
	"github.com/insolar/component-manager"
	"github.com/stretchr/testify/assert"
	"testing"
)

type supermarketMock struct {
	//	count int
	event chan bool
}

func (s *supermarketMock) Buy(thing core.GoodsType) (core.Product, error) {
	//s.count++
	s.event <- true
	return nil, nil
}

func TestCustomer_BuyGoods(t *testing.T) {
	bob := NewCustomer("Bob")
	mock := &supermarketMock{event: make(chan bool)}
	cm := component.NewManager(nil)
	cm.Inject(bob, mock)

	ctx := context.Background()
	err := cm.Start(ctx)
	assert.NoError(t, err)

	<-mock.event

	err = cm.Stop(ctx)
	assert.NoError(t, err)

	//assert.NotEqual(t, 0, mock.count)
}
