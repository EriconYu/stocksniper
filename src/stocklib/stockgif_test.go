package stocklib

import (
	"extern"
	"os"
	"testing"
)

//TestGetGifDirec ...
func TestGetGifDirec(t *testing.T) {
	getGifDirec(extern.SinaStockMin)
}

//TestGetGif ....
func TestGetGif(t *testing.T) {
	// ok := GetGif(extern.SinaStockMin, "sh", "000001")
	// if ok == false {
	// 	t.Error("GetGif Error")
	// }

	f, e := os.OpenFile("ss", os.O_RDWR, 0777)
	if e != nil {
		t.Error(e)
	}
	defer f.Close()
	f.Write([]byte("hello"))

}
