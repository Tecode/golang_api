package helpers

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// CutPicture 图片裁剪
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

// DecodePicture 根据图片类型解码
func DecodePicture(fileSuffix string, file *os.File) (image.Image, error) {
	// 根据图片类型转码
	switch fileSuffix {
	case ".jpg", ".jpeg":
		return jpeg.Decode(file)
	case ".png":
		return png.Decode(file)
	case ".gif":
		return gif.Decode(file)
	default:
		return nil, errors.New("格式转换错误")
	}
}

// 根据图片类型编码
func EncodePicture(fileSuffix string, buffer io.Writer, curImage image.Image) error {
	// 根据图片类型转码
	switch fileSuffix {
	case ".jpg", ".jpeg":
		jpeg.Encode(buffer, curImage, nil)
		return nil
	case ".png":
		png.Encode(buffer, curImage)
		return nil
	case ".gif":
		gif.Encode(buffer, curImage, nil)
		return nil
	default:
		return errors.New("格式转换错误")
	}
}
