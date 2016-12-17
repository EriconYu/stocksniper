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
		if _, e := stockinfo.GetStockInfo(); e == nil {
			//这里可以各种high了
			seelog.Info(stockinfo.StockName)
		} else {
			fmt.Printf("Get %s Info err:%s\n", stockinfo.StockID, e.Error())
		}
	}()
}

//RealtimeSHStocks 沪市所有股票的实时信息
func RealtimeSHStocks() {
	//实时数据
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sh"
		for i := 1; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, e := stockinfo.GetStockInfo(); e == nil {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info err:%s\n", stockinfo.SalesCity+stockinfo.StockID, e.Error())
			}
		}
	}()
}

//RealtimeSZStocks 深市所有股票的实时信息
func RealtimeSZStocks() {
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sz"
		for i := 1; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, e := stockinfo.GetStockInfo(); e == nil {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info err:%s\n", stockinfo.SalesCity+stockinfo.StockID, e.Error())
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
		for i := 1; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, e := stockinfo.GetStockInfo(); e == nil {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info err:%s\n", stockinfo.SalesCity+stockinfo.StockID, e.Error())
			}
		}
	}()
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sz"
		for i := 1; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if _, e := stockinfo.GetStockInfo(); e == nil {
				//seelog.Info(string(info))
			} else {
				fmt.Printf("Get %s Info err:%s\n", stockinfo.SalesCity+stockinfo.StockID, e.Error())
			}
		}
	}()

}
