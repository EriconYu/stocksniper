package handler

import (
	"stocklib/realtimelib"

	"fmt"

	"github.com/kataras/iris/context"
)

type StockWatchInfo struct {
	StockID string `json:"stock_id"`
}

func IndexHandler(ctx context.Context) {
	ctx.View("index.html")
}

func WatchHandler(ctx context.Context) {
	var stock_watched StockWatchInfo
	err := ctx.ReadJSON(&stock_watched)
	if err != nil {
		ctx.WriteString("error:" + err.Error())
		return
	}
	fmt.Printf("stock_watched is %v", stock_watched)
	stockInfo, _ := realtimelib.RealtimeOneStock(stock_watched.StockID)
	ctx.JSON(stockInfo)

}
