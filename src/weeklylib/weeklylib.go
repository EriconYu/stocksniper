package weeklylib

import (
	"extern"
	"fmt"
	"stocklib"
)

//WeeklySHGoroutine 沪市日K
func WeeklySHGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(extern.SinaStockWeekly, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
}

//WeeklySZGoroutine 深市日K
func WeeklySHGoroutine() {
	go func() {
		for i := 1; i < 999999; i++ {
			//
			go stocklib.GetGif(extern.SinaStockWeekly, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}

//WeeklyAllGoroutine 沪市日K
func WeeklySHGoroutine() {
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
