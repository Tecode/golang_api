package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

// 书籍信息相关
type BookController struct {
	beego.Controller
}

// @Title 获取推荐书籍信息
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router / [get]
func (u *BookController) GetInfo() {
	fmt.Println("书籍信息")
	u.Data["json"] = map[string]string{
		"recommends":   "书籍信息",
		"newBooks":     "书籍信息",
		"popularBooks": "最受欢迎的书籍"}
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