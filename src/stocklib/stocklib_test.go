package stocklib

import "testing"

//TestGetStockInfo ....
func TestGetStockInfo(t *testing.T) {
	var stockInfo StockPriceInfo
	stockInfo.SalesCity = "sz"
	stockInfo.StockID = "000025"
	info, err := stockInfo.GetStockInfo()
	if err != nil {
		t.Error("TestGetStockInfo err:", err)
	} else {
		t.Log("stockInfo is", info)
	}
	//fmt.Println("err is ", err, "info is ", info)
}

//TestParseStockInfo ...
func TestParseStockInfo(t *testing.T) {
	var stockPriceInfo StockPriceInfo
	stockPriceInfo.SalesCity = "sz"
	stockPriceInfo.StockID = "000025"
	if info, err := stockPriceInfo.GetStockInfo(); err == nil {
		stockPriceInfo.ParseStockInfo(info)
	}
}
