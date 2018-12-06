package controllers

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

// 验证token，token出错就返回错误信息
var TokenValid = func(ctx *context.Context) {
	fmt.Println(ctx, "------------------")
}
