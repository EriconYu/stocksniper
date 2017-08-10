package weeklylib

import (
	"externstock"
	"fmt"
	"stocklib"
)

func WeeklyKLine(StockID string) {
	var stockinfo stocklib.StockPriceInfo
	if []byte(StockID)[0] == '0' || []byte(StockID)[0] == '3' || []byte(StockID)[0] == '7' {
		stockinfo.SalesCity = "sz"
	} else {
		stockinfo.SalesCity = "sh"
	}
	stockinfo.StockID = StockID
	stocklib.GetGif(externstock.SinaStockWeekly, stockinfo.SalesCity, stockinfo.StockID)
}

//WeeklySHGoroutine 沪市周K
func WeeklySHGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockWeekly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
}

//WeeklySZGoroutine 深市周K
func WeeklySZGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockWeekly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}

//WeeklyAllGoroutine 沪市周K
func WeeklyALLGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockMonthly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockMonthly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}
