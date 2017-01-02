package realtimelib

import (
	"fmt"
	"stocklib"

	"github.com/cihub/seelog"
)

//RealtimeOneStock 特定股票的实时信息
func RealtimeOneStock(StockID string) {
	go func() {
		var stockinfo stocklib.StockPriceInfo
		if []byte(StockID)[0] == '0' {
			stockinfo.SalesCity = "sh"
		} else {
			stockinfo.SalesCity = "sz"
		}
		stockinfo.StockID = StockID
		if _, ok := stockinfo.GetStockInfo(); ok == true {
			//这里可以各种high了
			seelog.Info(stockinfo.StockName)
		} else {
			fmt.Printf("Get %s Info ok:%v\n", stockinfo.StockID, ok)
		}
	}()
}

//RealtimeSHStocks 沪市所有股票的实时信息
func RealtimeSHStocks() {
	//实时数据
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sh"
		for i := 1; i < 300000; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, ok := stockinfo.GetStockInfo(); ok == true {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
			}
		}
	}()
}

//RealtimeSZStocks 深市所有股票的实时信息
func RealtimeSZStocks() {
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sz"
		for i := 300000; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, ok := stockinfo.GetStockInfo(); ok == true {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
			}
		}
	}()

}

//RealtimeAllStocks 两市所有股票的实时信息
func RealtimeAllStocks() {
	//实时数据
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sh"
		for i := 1; i < 300000; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, ok := stockinfo.GetStockInfo(); ok == true {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
			}
		}
	}()
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sz"
		for i := 300000; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, ok := stockinfo.GetStockInfo(); ok == true {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info ok:%v\n", stockinfo.SalesCity+stockinfo.StockID, ok)
			}
		}
	}()

}
