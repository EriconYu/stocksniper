package stocklib

import (
	"extern"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/httplib"
)

//getGifDirec ...
func getGifDirec(typeGif string) (dir string) {
	dir = "img/" + time.Now().String()[:10] + "/"
	if e := os.Mkdir(dir, 0666); e != nil {
		fmt.Println("e is:", e)
	}
	// dir = dir + typeGif + "/"
	// if e := os.Mkdir(dir, 0666); e != nil {
	// 	fmt.Println("e is:", e)
	// }

	fmt.Println("dir is", dir)
	return dir

}

//getGifName ...
func getGifName(typeGif, salesCity, stockCode string) (gifName string) {
	gifName = time.Now().String()[:10] + "_" + typeGif + "_" + stockCode + ".gif"
	return gifName
}

//GetGif ... typeGif
func GetGif(typeGif, salesCity, stockCode string) (ok bool) {
	url := fmt.Sprintf(extern.SinaStockGifAPI, typeGif)
	url = url + salesCity + stockCode + ".gif"
	req := httplib.Get(url)
	if filestream, err := req.Bytes(); err == nil {
		if len(filestream) == 0 {
			fmt.Println(stockCode, "Not Find!")
			return false
		}
		f, e := os.OpenFile(getGifDirec(typeGif)+getGifName(typeGif, salesCity, stockCode), os.O_RDWR, 0666)
		if e == nil {
			f.Write(filestream)
		} else {
			fmt.Println("OpenFile e:", e, "Now Create It")
			f, e := os.Create(getGifDirec(typeGif) + getGifName(typeGif, salesCity, stockCode))
			if e == nil {
				f.Write(filestream)
				ok = true
			} else {
				fmt.Println("CreateFile e:", e)
				ok = false
			}
		}
		f.Close()
	} else {
		ok = false
	}
	return ok
}
