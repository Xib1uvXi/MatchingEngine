package orderbook

import "github.com/shopspring/decimal"

var testTimestamp = 123452342343
var testQuanity, _ = decimal.NewFromString("0.1")
var testPrice, _ = decimal.NewFromString("0.1")
var testOrderId = 1
var testTradeId = 1

var testTimestamp1 = 123452342345
var testQuanity1, _ = decimal.NewFromString("0.2")
var testPrice1, _ = decimal.NewFromString("0.1")
var testOrderId1 = 2
var testTradeId1 = 2

var testTimestamp2 = 123452342340
var testQuanity2, _ = decimal.NewFromString("0.2")
var testPrice2, _ = decimal.NewFromString("0.3")
var testOrderId2 = 3
var testTradeId2 = 3

var testTimestamp3 = 1234523
var testQuanity3, _ = decimal.NewFromString("200.0")
var testPrice3, _ = decimal.NewFromString("1.3")
var testOrderId3 = 3
var testTradeId3 = 3
