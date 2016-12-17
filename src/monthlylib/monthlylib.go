package monthlylib

import (
	"extern"
	"fmt"
	"stocklib"
)

//MonthlySHGoroutine 沪市日K
func MonthlySHGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(extern.SinaStockMonthly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
}

//MonthlySZGoroutine 沪市日K
func MonthlySZGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(extern.SinaStockMonthly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}

//MonthlyAllGoroutine 两市日K
func MonthlyAllGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(extern.SinaStockMonthly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(extern.SinaStockMonthly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}
