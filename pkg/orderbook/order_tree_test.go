package orderbook

import (
	"github.com/shopspring/decimal"
	"strconv"
	"testing"
)

func TestNewOrderTree(t *testing.T) {
	orderTree := NewOrderTree()

	dummyOrder := make(map[string]string)
	dummyOrder["timestamp"] = strconv.Itoa(testTimestamp)
	dummyOrder["quantity"] = testQuanity.String()
	dummyOrder["price"] = testPrice.String()
	dummyOrder["order_id"] = strconv.Itoa(testOrderId)
	dummyOrder["trade_id"] = strconv.Itoa(testTradeId)

	dummyOrder1 := make(map[string]string)
	dummyOrder1["timestamp"] = strconv.Itoa(testTimestamp1)
	dummyOrder1["quantity"] = testQuanity1.String()
	dummyOrder1["price"] = testPrice1.String()
	dummyOrder1["order_id"] = strconv.Itoa(testOrderId1)
	dummyOrder1["trade_id"] = strconv.Itoa(testTradeId1)

	dummyOrder2 := make(map[string]string)
	dummyOrder2["timestamp"] = strconv.Itoa(testTimestamp2)
	dummyOrder2["quantity"] = testQuanity2.String()
	dummyOrder2["price"] = testPrice2.String()
	dummyOrder2["order_id"] = strconv.Itoa(testOrderId2)
	dummyOrder2["trade_id"] = strconv.Itoa(testTradeId2)

	dummyOrder3 := make(map[string]string)
	dummyOrder3["timestamp"] = strconv.Itoa(testTimestamp3)
	dummyOrder3["quantity"] = testQuanity3.String()
	dummyOrder3["price"] = testPrice3.String()
	dummyOrder3["order_id"] = strconv.Itoa(testOrderId3)
	dummyOrder3["trade_id"] = strconv.Itoa(testTradeId3)

	if !(orderTree.volume.Equal(decimal.Zero)) {
		t.Errorf("orderTree.volume incorrect, got: %d, want: %d.", orderTree.volume, decimal.Zero)
	}

	if !(orderTree.Length() == 0) {
		t.Errorf("orderTree.Length() incorrect, got: %d, want: %d.", orderTree.Length(), 0)
	}

	orderTree.InsertOrder(dummyOrder)
	orderTree.InsertOrder(dummyOrder1)

	if !(orderTree.PriceExist(testPrice)) {
		t.Errorf("orderTree.numOrders incorrect, got: %d, want: %d.", orderTree.numOrders, 2)
	}

	if !(orderTree.PriceExist(testPrice1)) {
		t.Errorf("orderTree.numOrders incorrect, got: %d, want: %d.", orderTree.numOrders, 2)
	}

	if !(orderTree.Length() == 2) {
		t.Errorf("orderTree.numOrders incorrect, got: %d, want: %d.", orderTree.numOrders, 2)
	}

	orderTree.RemoveOrderById(dummyOrder1["order_id"])
	orderTree.RemoveOrderById(dummyOrder["order_id"])

	if !(orderTree.Length() == 0) {
		t.Errorf("orderTree.numOrders incorrect, got: %d, want: %d.", orderTree.numOrders, 2)
	}

	orderTree.InsertOrder(dummyOrder)
	orderTree.InsertOrder(dummyOrder1)
	orderTree.InsertOrder(dummyOrder2)
	orderTree.InsertOrder(dummyOrder3)

	if !(orderTree.MaxPrice().Equal(testPrice3)) {
		t.Errorf("orderTree.MaxPrice incorrect, got: %d, want: %d.", orderTree.MaxPrice(), testPrice3)
	}

	if !(orderTree.MinPrice().Equal(testPrice)) {
		t.Errorf("orderTree.MinPrice incorrect, got: %d, want: %d.", orderTree.MinPrice(), testPrice)
	}

	orderTree.RemovePrice(testPrice)

	if orderTree.PriceExist(testPrice) {
		t.Errorf("orderTree.MinPrice incorrect, got: %d, want: %d.", orderTree.MinPrice(), testPrice)
	}

	// TODO Check PriceList as well and verify with the orders
}
