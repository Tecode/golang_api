package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func CompressPicture(ctx *context.Context) {
	var buf bytes.Buffer
	size := 170
	file, err := os.Open("resource/images/6Co6W3DogOqW6TxwDp8Vb.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	fmt.Println(width,height, "图片宽高----------")
	// 压缩图片
	m := resize.Resize(200, 0, img, resize.Bicubic)
	// 裁剪图片
	curImage, _ := ImageCopy(m, 0, 0, size, size)
	jpeg.Encode(&buf, curImage, nil)
	file.Close()
	data := buf.Bytes()
	ctx.Output.Body(data)
}

func ImageCopy(src image.Image, x, y, w, h int) (image.Image, error) {

	var subImg image.Image

	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {

		return subImg, errors.New("图片解码失败")
	}

	return subImg, nil
}
