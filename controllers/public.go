package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/nfnt/resize"
	"golang_api/helpers"
	"image/jpeg"
	"log"
	"os"
	"path"
	"strconv"
)

// 公共的一些api
type PublicController struct {
	beego.Controller
}

// @Title 公共的一些api
// @Description 图片裁剪
// @Param   width   query   int  false       "宽度"
// @Param   height   query   int  false       "高度"
// @Param   img     path    string  true        "6Co6W3DogOqW6TxwDp8Vb.jpg"
// @Success 200 获取图片成功
// @router /picture/:img [get]
func (p *PublicController) Get() {
	var buffer bytes.Buffer
	imgPath := p.Ctx.Input.Param(":img")


	currentWidth, currentWidthErr := strconv.ParseInt(p.Ctx.Input.Query("width"), 10, 16)
	currentHeight, currentHeightErr := strconv.ParseInt(p.Ctx.Input.Query("height"), 10, 16)
	file, err := os.Open("resource/images/" + imgPath)
	fileSuffix := path.Ext(imgPath)
	fmt.Println(fileSuffix, "fileSuffix")
	if err != nil {
		log.Fatal(err)
	}
	// 根据图片类型转码
	switch fileSuffix {
	case ".jpg", ".jpeg":
		fmt.Println(".jpg, jpeg")
		break
	case ".png":
		fmt.Println(".png")
		break
	case ".gif":
		fmt.Println(".gif")
		break
	default:
		break

	}
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	b := img.Bounds()
	// 图片原始宽高
	width := b.Max.X
	height := b.Max.Y

	// 传入的宽高是否有错
	if currentHeightErr != nil || currentWidthErr != nil {
		fmt.Println("出错了")
		currentWidth = int64(b.Max.X)
		currentHeight = int64(b.Max.Y)
	}
	fmt.Println(width, height, int(currentWidth), int(currentHeight), "图片宽高----------")
	// 压缩图片
	m := resize.Resize(200, 0, img, resize.Bicubic)
	// 裁剪图片
	curImage, _ := helpers.CutPicture(m, 0, 0, int(currentWidth), int(currentHeight))
	jpeg.Encode(&buffer, curImage, nil)
	file.Close()
	data := buffer.Bytes()
	p.Ctx.Output.Body(data)
}
