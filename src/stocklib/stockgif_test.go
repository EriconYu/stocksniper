package stocklib

import (
	"externstock"
	"testing"
)

//TestGetGifDirec ...
func TestGetGifDirec(t *testing.T) {
	//getGifDirec(externstock.SinaStockMin)
}

//TestGetGif ....
func TestGetGif(t *testing.T) {
	ok := GetGif(externstock.SinaStockMin, "sh", "000001")
	if ok == false {
		t.Error("GetGif Error")
	}

	// f, e := os.OpenFile("ss", os.O_RDWR, 0777)
	// if e != nil {
	// 	t.Error(e)
	// }
	// defer f.Close()
	// f.Write([]byte("hello"))

}
