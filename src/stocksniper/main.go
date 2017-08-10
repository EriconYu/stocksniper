package main

import (
	"fmt"
	"handler"
	"stocklib/monthlylib"
	"stocklib/realtimelib"
	"time"

	"os"

	"github.com/kataras/iris"
)

func main() {
	fmt.Println("Hello,I am stock sniper!Biu Biu Biu ~")

	app := iris.New()

	app.Get("/index", handler.IndexHandler)
	app.Get("/watch/:stockid", handler.WatchHandler)
	app.Post("/watch", handler.WatchHandler)

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		fmt.Printf("iris app run err:", err)
		os.Exit(-1)
	}
	// UserAPPCode
	realtimelib.RealtimeAllStocks()
	monthlylib.MonthlyAllGoroutine()

	time.Sleep(10 * time.Second)
}
