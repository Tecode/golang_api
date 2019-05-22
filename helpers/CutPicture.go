package helpers

import (
	"errors"
	"fmt"
	"image"
)

// 图片裁剪
func CutPicture(src image.Image, x, y, width, height int) (image.Image, error) {

	var subImg image.Image

	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+width, y+height)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+width, y+height)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+width, y+height)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {

		return subImg, errors.New("图片解码失败")
	}

	return subImg, nil
}

// 图片转码

func EncodePicture(fileSuffix string, file image.Image) {
	// 根据图片类型转码
	switch fileSuffix {
	case ".jpg", ".jpeg":
		fmt.Println(".jpg, jpeg")
	case ".png":
		fmt.Println(".png")
		break
	case ".gif":
		fmt.Println(".gif")
		break
	default:
		break

	}
}