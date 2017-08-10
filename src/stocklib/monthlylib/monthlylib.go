package monthlylib

import (
	"externstock"
	"fmt"
	"stocklib"
)

func MonthlyKLine(StockID string) {
	var stockinfo stocklib.StockPriceInfo
	stockinfo.SalesCity = stocklib.GetSalesCity(StockID)

	stockinfo.StockID = StockID
	stocklib.GetGif(externstock.SinaStockMonthly, stockinfo.SalesCity, stockinfo.StockID)
}

//MonthlySHGoroutine 沪市月K
func MonthlySHGoroutine() {
	go func() {
		for i := 1; i < 300000; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockMonthly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
}

//MonthlySZGoroutine 沪市月K
func MonthlySZGoroutine() {
	go func() {
		for i := 300000; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockMonthly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}

//MonthlyAllGoroutine 两市月K
func MonthlyAllGoroutine() {
	go func() {
		for i := 1; i < 100000; i++ {
			//
			MonthlyKLine(fmt.Sprintf("%.6d", i))
		}
	}()
	go func() {
		for i := 300000; i < 400000; i++ {
			//
			MonthlyKLine(fmt.Sprintf("%.6d", i))
		}
	}()
	go func() {
		for i := 600000; i < 800000; i++ {
			//
			MonthlyKLine(fmt.Sprintf("%.6d", i))
		}
	}()
}
