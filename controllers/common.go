package controllers

import (
	"bytes"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/nfnt/resize"
	"golang_apiv2/utils"
	"image"
	"log"
	"os"
	"path"
	"strconv"
)

// CommonController operations for Common
type CommonController struct {
	beego.Controller
}

// ImageResize ...
// @Title ImageResize
// @Description 图片自定义尺寸
// @Param   width   query   int  false       "宽度"
// @Param   height   query   int  false       "高度"
// @Param   img     path    string  true        "6Co6W3DogOqW6TxwDp8Vb.png"
// @Success 200 获取图片成功
// @router /picture/:img [get]
func (c *CommonController) ImageResize() {
	var buffer bytes.Buffer
	imgPath := c.GetString(":img")
	var compressImage image.Image

	currentWidth, currentWidthErr := strconv.ParseInt(c.Ctx.Input.Query("width"), 10, 16)
	currentHeight, currentHeightErr := strconv.ParseInt(c.Ctx.Input.Query("height"), 10, 16)
	file, err := os.Open("resource/images/" + imgPath)
	fileSuffix := path.Ext(imgPath)
	fmt.Println(fileSuffix, imgPath, "fileSuffix")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := utils.DecodePicture(fileSuffix, file)
	if err != nil {
		logs.Error(err.Error())
		return
	}
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
	// 判断图片宽高比
	// 压缩图片

	compressImage = resize.Resize(uint(currentWidth), 0, img, resize.Bicubic)
	// 裁剪图片
	curImage, _ := utils.CutPicture(compressImage, 0, 0, int(currentWidth), int(currentHeight))
	pictureErr := utils.EncodePicture(fileSuffix, &buffer, curImage)
	if pictureErr != nil {
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("无法关闭文件:", err)
		}
	}()
	data := buffer.Bytes()
	if bodyErr := c.Ctx.Output.Body(data); bodyErr != nil {
		return
	}
}
