package controllers

import (
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

// @Title 获取推荐书籍信息（全部）
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Param	index		query 	int	true		"The index is required"
// @Param	size		query 	int	true		"The size is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /recommend [get]
func (u *BookController) GetRecommend() {
	index, indexError := u.GetInt("index")
	size, sizeError := u.GetInt("size")
	if indexError != nil || sizeError != nil {
		u.Data["json"] = map[string]string{
			"data": "分页参数错误"}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]models.PageData{
		"data": models.Recommend(index, size)}
	u.ServeJSON()
}

// @Title 获取最新发布的书籍（全部）
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Param	index		query 	int	true		"The index is required"
// @Param	size		query 	int	true		"The size is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /new_book [get]
func (u *BookController) GetNewBook() {
	index, indexError := u.GetInt("index")
	size, sizeError := u.GetInt("size")
	if indexError != nil || sizeError != nil {
		u.Data["json"] = map[string]string{
			"data": "分页参数错误"}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]models.PageData{
		"data": models.NewBook(index, size)}
	u.ServeJSON()
}

// @Title 获取最受欢迎的书籍（全部）
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Param	index		query 	int	true		"The index is required"
// @Param	size		query 	int	true		"The size is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /popular_book [get]
func (u *BookController) GetPopularBook() {
	index, indexError := u.GetInt("index")
	size, sizeError := u.GetInt("size")
	if indexError != nil || sizeError != nil {
		u.Data["json"] = map[string]string{
			"data": "分页参数错误"}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]models.PageData{
		"data": models.PopularBook(index, size)}
	u.ServeJSON()
}