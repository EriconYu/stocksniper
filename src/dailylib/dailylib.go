package dailylib

import (
	"externstock"
	"fmt"
	"stocklib"
)

//DailySHGoroutine 沪市日K
func DailySHGoroutine() {
	go func() {
		for i := 1; i < 300000; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockDaily, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
}

//DailySZGoroutine 深市日K
func DailySZGoroutine() {
	go func() {
		for i := 300000; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockDaily, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}

//DailyAllGoroutine  两市日K
func DailyAllGoroutine() {
	go func() {
		for i := 1; i < 300000; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockDaily, "sh", fmt.Sprintf("%.6d", i))
		}
	}()
	go func() {
		for i := 300000; i < 999999; i++ {
			//
			go stocklib.GetGif(externstock.SinaStockDaily, "sz", fmt.Sprintf("%.6d", i))
		}
	}()
}
