package minlib

import (
	"extern"
	"fmt"
	"stocklib"
	"time"
)

//MinSHGoroutine 沪市所有股票分时图片获取
func MinSHGoroutine() {
	//分时
	go func() {
		for i := 1; i < 999999; i++ {
			stocklib.GetGif(extern.SinaStockMin, "sh", fmt.Sprintf("%.6d", i))
		}
		shGetGif <- time.Now().String()[:19]
	}()
}

//MinSZGoroutine 深市所有股票分时图片获取
func MinSZGoroutine() {
	//分时
	go func() {
		for i := 1; i < 999999; i++ {
			stocklib.GetGif(extern.SinaStockMin, "sz", fmt.Sprintf("%.6d", i))
		}
		szGetGif <- time.Now().String()[:19]
	}()
}

//MinAllGoroutine 所有股票分时图片获取
func MinAllGoroutine() {
	//分时
	go func() {
		for i := 1; i < 999999; i++ {
			stocklib.GetGif(extern.SinaStockMin, "sh", fmt.Sprintf("%.6d", i))
		}
		shGetGif <- time.Now().String()[:19]
	}()
	go func() {
		for i := 1; i < 999999; i++ {
			stocklib.GetGif(extern.SinaStockMin, "sz", fmt.Sprintf("%.6d", i))
		}
		szGetGif <- time.Now().String()[:19]
	}()
}
