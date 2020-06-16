package orderbook

import (
	"github.com/HuKeping/rbtree"
	"github.com/shopspring/decimal"
)

type OrderList struct {
	length    int
	headOrder *Order
	tailOrder *Order
	lastOrder *Order
	volume    decimal.Decimal
	price     decimal.Decimal
}

func NewOrderList(price decimal.Decimal) *OrderList {
	return &OrderList{headOrder: nil, tailOrder: nil, length: 0, volume: decimal.Zero,
		lastOrder: nil, price: price}
}

func (ol *OrderList) Less(than rbtree.Item) bool {
	return ol.price.LessThan(than.(*OrderList).price)
}

func (ol *OrderList) Length() int {
	return ol.length
}

func (ol *OrderList) HeadOrder() *Order {
	return ol.headOrder
}

func (ol *OrderList) AppendOrder(order *Order) {
	if ol.Length() == 0 {
		order.nextOrder = nil
		order.prevOrder = nil
		ol.headOrder = order
		ol.tailOrder = order
	} else {
		order.prevOrder = ol.tailOrder
		order.nextOrder = nil
		ol.tailOrder.nextOrder = order
		ol.tailOrder = order
	}
	ol.length = ol.length + 1
	ol.volume = ol.volume.Add(order.quantity)
}

func (ol *OrderList) RemoveOrder(order *Order) {
	ol.volume = ol.volume.Sub(order.quantity)
	ol.length = ol.length - 1
	if ol.length == 0 {
		return
	}

	nextOrder := order.nextOrder
	prevOrder := order.prevOrder

	if nextOrder != nil && prevOrder != nil {
		nextOrder.prevOrder = prevOrder
		prevOrder.nextOrder = nextOrder
	} else if nextOrder != nil {
		nextOrder.prevOrder = nil
		ol.headOrder = nextOrder
	} else if prevOrder != nil {
		prevOrder.nextOrder = nil
		ol.tailOrder = prevOrder
	}
}

func (ol *OrderList) MoveToTail(order *Order) {
	// This Order is not the first Order in the OrderList
	if order.prevOrder != nil {
		// Link the previous Order to the next Order, then move the Order to tail
		order.prevOrder.nextOrder = order.nextOrder
	} else {
		// This Order is the first Order in the OrderList
		// Make next order the first
		ol.headOrder = order.nextOrder
	}
	order.nextOrder.prevOrder = order.prevOrder

	// Move Order to the last position. Link up the previous last position Order.
	ol.tailOrder.nextOrder = order
	ol.tailOrder = order
}
