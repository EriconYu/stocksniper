package main

import (
	"fmt"
	"handler"
	"stocklib/realtimelib"
	"time"

	"os"

	"github.com/kataras/iris"
)

func main() {
	fmt.Println("Hello,I am stock sniper!Biu Biu Biu ~")

	app := iris.New()

	//app.StaticServe("./static/html/", "/static/")
	app.StaticServe("./static/css/", "/css/")
	app.StaticServe("./static/js/", "/js/")

	htmlTmpl := iris.HTML("./static/html/", ".html")
	htmlTmpl.Reload(true)
	app.RegisterView(htmlTmpl)

	app.Get("/", handler.IndexHandler)
	app.Get("/index", handler.IndexHandler)
	app.Get("/watch/:stockid", handler.WatchHandler)
	app.Post("/watch", handler.WatchHandler)

	err := app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
	if err != nil {
		fmt.Printf("iris app run err:", err)
		os.Exit(-1)
	}
	// UserAPPCode
	realtimelib.RealtimeAllStocks()
	//	monthlylib.MonthlyAllGoroutine()

	time.Sleep(10 * time.Second)
}
