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
func (b *BookController) Today() {
	b.Data["json"] = models.BookInfo()
	b.ServeJSON()
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
func (b *BookController) GetRecommend() {
	index, indexError := b.GetInt("index")
	size, sizeError := b.GetInt("size")
	if indexError != nil || sizeError != nil {
		b.Ctx.Output.SetStatus(400)
		b.Data["json"] = map[string]string{
			"data": "分页参数错误"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]models.PageData{
		"data": models.Recommend(index, size)}
	b.ServeJSON()
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
func (b *BookController) GetNewBook() {
	index, indexError := b.GetInt("index")
	size, sizeError := b.GetInt("size")
	if indexError != nil || sizeError != nil {
		b.Ctx.Output.SetStatus(400)
		b.Data["json"] = map[string]string{
			"data": "分页参数错误"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]models.PageData{
		"data": models.NewBook(index, size)}
	b.ServeJSON()
}

// @Title 获取最受欢迎的书籍
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Param	index		query 	int	true		"The index is required"
// @Param	size		query 	int	true		"The size is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /popular_book [get]
func (b *BookController) GetPopularBook() {
	index, indexError := b.GetInt("index")
	size, sizeError := b.GetInt("size")
	if indexError != nil || sizeError != nil {
		b.Ctx.Output.SetStatus(400)
		b.Data["json"] = map[string]string{
			"data": "分页参数错误"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]models.PageData{
		"data": models.PopularBook(index, size)}
	b.ServeJSON()
}

// @Title 获取热门书籍
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Param	index		query 	int	true		"The index is required"
// @Param	size		query 	int	true		"The size is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /hot_book [get]
func (b *BookController) GetHotBook() {
	index, indexError := b.GetInt("index")
	size, sizeError := b.GetInt("size")
	if indexError != nil || sizeError != nil {
		b.Ctx.Output.SetStatus(400)
		b.Data["json"] = map[string]string{
			"data": "分页参数错误"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]models.PageData{
		"data": models.HotBook(index, size)}
	b.ServeJSON()
}

// @Title 获取书籍详情
// @Description create books
// @Param	token		header 	string	true		"The token is required"
// @Param	book_id		path 	int	true		"The book_id is required"
// @Success 200 {object} models.BookInfo
// @Failure 403 body is empty
// @Failure 400 token验证失败
// @router /detail/:book_id [get]
func (b *BookController) GetBookDetail() {
	bookId, bookIdError := b.GetInt(":book_id")
	if bookIdError != nil {
		b.Ctx.Output.SetStatus(400)
		b.Data["json"] = map[string]string{
			"data": "请求参数错误"}
		b.ServeJSON()
		return
	}
	b.Data["json"] = map[string]models.SiteAppBook{
		"data": models.BookDetail(bookId)}
	b.ServeJSON()
}