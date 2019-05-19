package helpers

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

func CompressPicture(ctx *context.Context) {
	fmt.Println("图片压缩方法")
	ctx.Output.JSON(map[string]interface{}{"data": "455"}, true, false)
}
