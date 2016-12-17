package stocklib

import (
	"externstock"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/cihub/seelog"
)

//getGifDirec ...
func getGifDirec(typeGif string) (dir string) {
	if time.Now().Weekday() == time.Saturday {
		dir = "img/" + time.Now().AddDate(0, 0, -1).String()[:10] + "/"
	} else if time.Now().Weekday() == time.Sunday {
		dir = "img/" + time.Now().AddDate(0, 0, -2).String()[:10] + "/"
	} else if time.Now().Weekday() == time.Monday {
		if time.Now().Hour() < 9 {
			dir = "img/" + time.Now().AddDate(0, 0, -3).String()[:10] + "/"
		} else {
			dir = "img/" + time.Now().String()[:10] + "/"
		}
	} else {
		if time.Now().Hour() < 9 {
			dir = "img/" + time.Now().AddDate(0, 0, -1).String()[:10] + "/"

		} else {
			dir = "img/" + time.Now().String()[:10] + "/"
		}
	}
	return dir
}

//getGifName ...
func getGifName(typeGif, salesCity, stockCode string) (gifName string) {
	if time.Now().Weekday() == time.Saturday {
		gifName = time.Now().AddDate(0, 0, -1).String()[:10] + "_" + typeGif + "_" + stockCode + ".gif"
	} else if time.Now().Weekday() == time.Sunday {
		gifName = time.Now().AddDate(0, 0, -2).String()[:10] + "_" + typeGif + "_" + stockCode + ".gif"
	} else if time.Now().Weekday() == time.Monday {
		if time.Now().Hour() < 9 {
			gifName = time.Now().AddDate(0, 0, -3).String()[:10] + "_" + typeGif + "_" + stockCode + ".gif"
		} else {
			gifName = time.Now().String()[:10] + "_" + typeGif + "_" + stockCode + ".gif"
		}
	} else {
		if time.Now().Hour() < 9 {
			gifName = time.Now().AddDate(0, 0, -1).String()[:10] + "_" + typeGif + "_" + stockCode + ".gif"

		} else {
			gifName = time.Now().String()[:10] + "_" + typeGif + "_" + stockCode + ".gif"
		}
	}
	return gifName
}

//GetGif ...
func GetGif(typeGif, salesCity, stockCode string) (ok bool) {
	url := fmt.Sprintf(externstock.SinaStockGifAPI, typeGif)
	url = url + salesCity + stockCode + ".gif"
	req := httplib.Get(url)
	if filestream, err := req.Bytes(); err == nil {
		if len(filestream) == 0 {
			seelog.Error(salesCity, stockCode, "GetGif Not Find!")
			return false
		}
		dir := getGifDirec(typeGif)
		if e := os.MkdirAll(dir, 0777); e != nil {
			seelog.Error("GetGif MkdirAll e is:", e)
			return false
		}

		f, e := os.OpenFile(dir+getGifName(typeGif, salesCity, stockCode), os.O_RDWR, 0666)
		if e == nil {
			f.Write(filestream)
			ok = true
		} else {
			//seelog.Info("OpenFile e:", e, "Now Create It")
			f, e := os.Create(dir + getGifName(typeGif, salesCity, stockCode))
			if e == nil {
				f.Write(filestream)
				ok = true
			} else {
				seelog.Error("CreateFile e:", e)
				ok = false
			}
		}
		f.Close()
	} else {
		ok = false
	}
	return ok
}
