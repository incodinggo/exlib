package imgex_test

import (
	"fmt"
	"github.com/incodinggo/exlib/imgex"
	"image/jpeg"
	"os"
	"path"
	"testing"
)

func TestCompress(t *testing.T) {
	s := `/Users/AC/Desktop/1.jpg`
	img, _ := imgex.Compress(s, 128)
	if f, err := os.Create(path.Join(path.Dir(s), "1"+"_compress"+path.Ext(s))); err != nil {
		fmt.Println(err)
		return
	} else {
		defer f.Close()
		jpeg.Encode(f, img, nil)
	}
}
