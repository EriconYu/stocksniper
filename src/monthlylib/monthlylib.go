package monthlylib

import (
	"externstock"
	"fmt"
	"stocklib"
)

//MonthlySHGoroutine 沪市日K
func MonthlySHGoroutine() {
	go func() {
		for i := 1; i < 300000; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockMonthly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
}

//MonthlySZGoroutine 沪市日K
func MonthlySZGoroutine() {
	go func() {
		for i := 300000; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockMonthly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}

//MonthlyAllGoroutine 两市日K
func MonthlyAllGoroutine() {
	go func() {
		for i := 1; i < 300000; i++ {
			//
			stocklib.GetGif(externstock.SinaStockMonthly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
	go func() {
		for i := 300000; i < 999999; i++ {
			//
			stocklib.GetGif(externstock.SinaStockMonthly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}
