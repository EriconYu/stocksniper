package main

import (
	"extern"
	"fmt"
	"stocklib"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Hello,I am stock sniper!Bui Bui Bui ~")

	var rlimit syscall.Rlimit
	syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)

	fmt.Println("rlimit is ", rlimit)
	rlimit.Cur = 65536
	rlimit.Max = 65536
	syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	var shGetGif = make(chan string)
	var szGetGif = make(chan string)
	var shGetInfo = make(chan string)
	var szGetInfo = make(chan string)
	go func() {
		for i := 1; i < 999999; i++ {
			oksh := stocklib.GetGif(extern.SinaStockMin, "sh", fmt.Sprintf("%.6d", i))
			fmt.Println("i is", i, ";oksh is", oksh)
		}
		shGetGif <- time.Now().String()[:19]
	}()
	go func() {
		for i := 1; i < 999999; i++ {
			oksz := stocklib.GetGif(extern.SinaStockMin, "sz", fmt.Sprintf("%.6d", i))
			fmt.Println("i is", i, ";oksz is", oksz)
		}
		szGetGif <- time.Now().String()[:19]
	}()
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sh"
		for i := 1; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if info, e := stockinfo.GetStockInfo(); e == nil {
				fmt.Println(string(info))
			} else {
				fmt.Printf("Get %s Info err:%s\n", stockinfo.SalesCity+stockinfo.StockID, e.Error())
			}
		}
		shGetInfo <- time.Now().String()[:19]
	}()
	go func() {
		var stockinfo stocklib.StockPriceInfo
		stockinfo.SalesCity = "sz"
		for i := 1; i < 999999; i++ {
			stockinfo.StockID = fmt.Sprintf("%.6d", i)
			if info, e := stockinfo.GetStockInfo(); e == nil {
				fmt.Println(string(info))
			} else {
				fmt.Printf("Get %s Info err:%s\n", stockinfo.SalesCity+stockinfo.StockID, e.Error())
			}
		}
		shGetInfo <- time.Now().String()[:19]
	}()
	fmt.Println("shGetGif is", <-shGetGif)
	fmt.Println("szGetGif is", <-szGetGif)
	fmt.Println("shGetInfo is", <-shGetInfo)
	fmt.Println("szGetInfo is", <-szGetInfo)
}
