package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"golang_api/models"
)

// 书籍信息相关
type BookController struct {
	beego.Controller
}

// @Title 首页展示信息
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /today [get]
func (u *BookController) Today() {
	u.Data["json"] = models.BookInfo()
	u.ServeJSON()
}

// @Title 获取推荐书籍信息
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /recommend [get]
func (u *BookController) GetRecommend() {
	fmt.Println("书籍信息")
	u.Data["json"] = map[string]string{"uid": "书籍信息"}
	u.ServeJSON()
}
