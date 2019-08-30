package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/nfnt/resize"
	"golang_api/helpers"
	"image"
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
	imgPath := p.GetString(":img")
	var compressImage image.Image

	currentWidth, currentWidthErr := strconv.ParseInt(p.Ctx.Input.Query("width"), 10, 16)
	currentHeight, currentHeightErr := strconv.ParseInt(p.Ctx.Input.Query("height"), 10, 16)
	file, err := os.Open("resource/images/" + imgPath)
	fileSuffix := path.Ext(imgPath)
	fmt.Println(fileSuffix, "fileSuffix")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := helpers.DecodePicture(fileSuffix, file)
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
	curImage, _ := helpers.CutPicture(compressImage, 0, 0, int(currentWidth), int(currentHeight))
	helpers.EncodePicture(fileSuffix, &buffer, curImage)
	file.Close()
	data := buffer.Bytes()
	p.Ctx.Output.Body(data)
}

// @Title 版本信息
// @Description 版本信息
// @Success 200 {object} Version
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /version [get]
func (v *PublicController) GetVersion() {
	v.Data["json"] = map[string]interface{}{
		"version": "1.0.0",
		"data": []string{
			"1、升级搜索页面UI",
			"2、优化下拉卡顿的问题",
			"3、iOS 版本支持 iPhone X"}}
	v.ServeJSON()
}
