package realtimelib

import (
	"fmt"
	"stocklib"
)

//RealtimeOneStock 特定股票的实时信息
func RealtimeOneStock(StockID string) (stockinfo stocklib.StockPriceInfo, ok bool) {
	stockinfo.SalesCity = stocklib.GetSalesCity(StockID)
	stockinfo.StockID = StockID
	if _, ok := stockinfo.GetStockInfo(); ok == true {
		//这里可以各种high了
		fmt.Printf("stockinfo is %v", stockinfo)
	} else {
		fmt.Printf("Get %s Info ok:%v\n", stockinfo.StockID, ok)
	}
	return stockinfo, ok
}

//RealtimeSHStocks 沪市所有股票的实时信息
func RealtimeSHStocks() {
	//实时数据
	var stockinfo stocklib.StockPriceInfo
	stockinfo.SalesCity = "sh"
	for i := 600000; i < 699999; i++ {
		stockinfo.StockID = fmt.Sprintf("%.6d", i)
		if _, ok := stockinfo.GetStockInfo(); ok == true {
			//seelog.Info(string(info))
		} else {
			fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
		}
	}
}

//RealtimeSZStocks 深市所有股票的实时信息
func RealtimeSZStocks() {
	var stockinfo stocklib.StockPriceInfo
	stockinfo.SalesCity = "sz"
	for i := 1; i < 400000; i++ {
		stockinfo.StockID = fmt.Sprintf("%.6d", i)
		if _, ok := stockinfo.GetStockInfo(); ok == true {
			//seelog.Info(string(info))
		} else {
			fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
		}
	}
	for i := 700000; i < 800000; i++ {
		stockinfo.StockID = fmt.Sprintf("%.6d", i)
		if _, ok := stockinfo.GetStockInfo(); ok == true {
			//seelog.Info(string(info))
		} else {
			fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
		}
	}
}

//RealtimeAllStocks 两市所有股票的实时信息
func RealtimeAllStocks() {
	//实时数据
	go func() { //深市
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sz"
		for i := 1; i < 400000; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if info, ok := stockinfo.GetStockInfo(); ok == true {
				fmt.Printf("stockinfo is %v /r/n", info)
			} else {
				fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
			}
		}
	}()
	go func() { //沪市
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sh"
		for i := 600000; i < 699999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if info, ok := stockinfo.GetStockInfo(); ok == true {
				fmt.Printf("stockinfo is %v /r/n", info)
			} else {
				fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
			}
		}
	}()

}
