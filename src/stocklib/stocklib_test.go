package stocklib

import "testing"

//TestGetStockInfo ....
func TestGetStockInfo(t *testing.T) {
	var stockInfo StockPriceInfo
	stockInfo.SalesCity = "sz"
	stockInfo.StockID = "000025"
	//info, err := stockInfo.GetStockInfo()
	_, err := stockInfo.GetStockInfo()
	if err != true {
		t.Error("TestGetStockInfo err:", err)
	}
}

//TestParseStockInfo ...
func TestParseStockInfo(t *testing.T) {
	// var stockPriceInfo StockPriceInfo
	// stockPriceInfo.SalesCity = "sz"
	// stockPriceInfo.StockID = "000025"
	// if info, err := stockPriceInfo.GetStockInfo(); err == nil {
	// 	t.Log(info)
	// } else {
	// 	t.Error("GetStockInfo err:", err)
	// }
}
