package stocklib

import (
	"extern"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/httplib"
	"github.com/axgle/mahonia"
)

//StockPriceInfo 股票价格信息
type StockPriceInfo struct {
	SalesCity             string  // 交易市场 “sz”  "sh"
	StockID               string  // 股票代码
	StockName             string  // 股票名字
	OpeningPriceToday     float64 // 今日开盘价
	ClosingPriceYesterday float64 // 昨日收盘价
	PriceNow              float64 // 当前价格
	HighestPriceToday     float64 // 今日最高价
	LowestPriceToday      float64 // 今日最低价
	BuyPrice              float64 // 竞买价 买一
	SellPrice             float64 // 竞卖价 卖一
	//成交股票数，由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百；
	TradedNumber uint64
	//成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万；
	TradedAmount float64
	BuyHands1    float64 // 买一申请手
	BuyPrice1    float64 // 买一申请价格
	BuyHands2    float64 // 买二申请手
	BuyPrice2    float64 // 买二申请价格
	BuyHands3    float64 // 买三申请手
	BuyPrice3    float64 // 买三申请价格
	BuyHands4    float64 // 买四申请手
	BuyPrice4    float64 // 买四申请价格
	BuyHands5    float64 // 买五申请手
	BuyPrice5    float64 // 买五申请价格
	SoldHands1   float64 // 卖一申请手
	SoldPrice1   float64 // 卖一申请价格
	SoldHands2   float64 // 卖二申请手
	SoldPrice2   float64 // 卖二申请价格
	SoldHands3   float64 // 卖三申请手
	SoldPrice3   float64 // 卖三申请价格
	SoldHands4   float64 // 卖四申请手
	SoldPrice4   float64 // 卖四申请价格
	SoldHands5   float64 // 卖五申请手
	SoldPrice5   float64 // 卖五申请价格
	Data         string  // 日期
	Time         string  // 时间
}

//GetStockInfo 获取股票信息
func (s *StockPriceInfo) GetStockInfo() (stockInfo string, err error) {
	stockURL := extern.SinaStockAPI + s.SalesCity + s.StockID
	req := httplib.Get(stockURL)
	if input, err := req.Bytes(); err == nil {
		dec := mahonia.NewDecoder("gbk")
		stockInfo = dec.ConvertString(string(input))
	} else {
		fmt.Println("GetStockInfo err:", err)
	}
	return stockInfo, err
}

func (s *StockPriceInfo) parseStockID(stockInfo string) {
	stockid := strings.Split(stockInfo, "=")[0]
	stockid = strings.TrimSpace(stockid)
	stockid = stockid[11:]
	s.SalesCity = stockid[:2]
	s.StockID = stockid[2:]
	//fmt.Println("city is", s.SalesCity)
	//fmt.Println("id is", s.StockID)
}

func (s *StockPriceInfo) parseStockPrice(stockInfo string) {
	lenth := len(stockInfo)
	stockPriceInfo := stockInfo[21 : lenth-3]
	fmt.Println("stockPriceInfo is", stockPriceInfo)
	PriceArray := strings.Split(stockPriceInfo, ",")
	//lenth = len(PriceArray)
	// fmt.Println("lenth is ", lenth)  //33
	s.StockName = PriceArray[0]
	s.OpeningPriceToday, _ = strconv.ParseFloat(PriceArray[1], 64)
	s.ClosingPriceYesterday, _ = strconv.ParseFloat(PriceArray[2], 64)
	s.PriceNow, _ = strconv.ParseFloat(PriceArray[3], 64)
	s.HighestPriceToday, _ = strconv.ParseFloat(PriceArray[4], 64)
	s.LowestPriceToday, _ = strconv.ParseFloat(PriceArray[5], 64)
	s.BuyPrice, _ = strconv.ParseFloat(PriceArray[6], 64)
	s.SellPrice, _ = strconv.ParseFloat(PriceArray[7], 64)
	s.TradedNumber, _ = strconv.ParseUint(PriceArray[8], 10, 64)
	s.TradedAmount, _ = strconv.ParseFloat(PriceArray[9], 64)
	s.BuyHands1, _ = strconv.ParseFloat(PriceArray[10], 64)
	s.BuyPrice1, _ = strconv.ParseFloat(PriceArray[11], 64)
	s.BuyHands2, _ = strconv.ParseFloat(PriceArray[12], 64)
	s.BuyPrice2, _ = strconv.ParseFloat(PriceArray[13], 64)
	s.BuyHands3, _ = strconv.ParseFloat(PriceArray[14], 64)
	s.BuyPrice3, _ = strconv.ParseFloat(PriceArray[15], 64)
	s.BuyHands4, _ = strconv.ParseFloat(PriceArray[16], 64)
	s.BuyPrice4, _ = strconv.ParseFloat(PriceArray[17], 64)
	s.BuyHands5, _ = strconv.ParseFloat(PriceArray[18], 64)
	s.BuyPrice5, _ = strconv.ParseFloat(PriceArray[19], 64)
	s.SoldHands1, _ = strconv.ParseFloat(PriceArray[20], 64)
	s.SoldPrice1, _ = strconv.ParseFloat(PriceArray[21], 64)
	s.SoldHands2, _ = strconv.ParseFloat(PriceArray[22], 64)
	s.SoldPrice2, _ = strconv.ParseFloat(PriceArray[23], 64)
	s.SoldHands3, _ = strconv.ParseFloat(PriceArray[24], 64)
	s.SoldPrice3, _ = strconv.ParseFloat(PriceArray[25], 64)
	s.SoldHands4, _ = strconv.ParseFloat(PriceArray[26], 64)
	s.SoldPrice4, _ = strconv.ParseFloat(PriceArray[27], 64)
	s.SoldHands5, _ = strconv.ParseFloat(PriceArray[28], 64)
	s.SoldPrice5, _ = strconv.ParseFloat(PriceArray[29], 64)
	s.Data = PriceArray[30]
	s.Time = PriceArray[31]
	//fmt.Println(s)
}

//ParseStockInfo ...
func (s *StockPriceInfo) ParseStockInfo(stockInfo string) {
	s.parseStockID(stockInfo)
	s.parseStockPrice(stockInfo)
}
