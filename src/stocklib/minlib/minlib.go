package minlib

import (
	"externstock"
	"fmt"
	"stocklib"
)

//MinSHGoroutine 沪市所有股票分时图片获取
func MinSHGoroutine() {
	//分时
	for i := 100000; i < 200000; i++ {
		stocklib.GetGif(externstock.SinaStockMin, "sh", fmt.Sprintf("%.6d", i))
	}
	for i := 600000; i < 700000; i++ {
		stocklib.GetGif(externstock.SinaStockMin, "sh", fmt.Sprintf("%.6d", i))
	}
}

//MinSZGoroutine 深市所有股票分时图片获取
func MinSZGoroutine() {
	//分时
	for i := 1; i < 100000; i++ {
		stocklib.GetGif(externstock.SinaStockMin, "sz", fmt.Sprintf("%.6d", i))
	}
	for i := 300000; i < 400000; i++ {
		stocklib.GetGif(externstock.SinaStockMin, "sz", fmt.Sprintf("%.6d", i))
	}
	for i := 700000; i < 800000; i++ {
		stocklib.GetGif(externstock.SinaStockMin, "sz", fmt.Sprintf("%.6d", i))
	}
}

//MinAllGoroutine 所有股票分时图片获取
func MinAllGoroutine() {
	//分时
	go MinSHGoroutine()
	go MinSZGoroutine()
}
