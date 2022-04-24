package imgex

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
	"strings"
)

func Compress(src string, height uint) (newImg image.Image, err error) {
	reg, _ := regexp.Compile(`^.*\.((png)|(jpg)|(jpeg))$`)
	if !reg.MatchString(src) {
		return nil, fmt.Errorf("support png|jpg|jpeg only")
	}
	var f *os.File
	if f, err = os.Open(src); err != nil {
		return nil, err
	}
	defer f.Close()
	var img image.Image
	switch {
	case strings.HasSuffix(f.Name(), "png"):
		if img, err = png.Decode(f); err != nil {
			return nil, err
		}
	case strings.HasSuffix(f.Name(), "jpg") || strings.HasSuffix(f.Name(), "jpeg"):
		if img, err = jpeg.Decode(f); err != nil {
			return nil, err
		}
	}
	newImg = resize.Resize(0, height, img, resize.Lanczos3)
	return newImg, nil
}
