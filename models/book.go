package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 今日推荐书籍
type TodayInfo struct {
	Recommends   []SiteAppBook `json:"recommendBooks"`
	NewBooks     []SiteAppBook `json:"newBooks"`
	PopularBooks []SiteAppBook `json:"popularBooks"`
}

// 分页数据
type PageData struct {
	Data []SiteAppBook `json:"data"`
	Index int `json:"index"`
	Size int `json:"size"`
	TotalElements int64 `json:"totalElements"`
} 

func BookInfo() TodayInfo {
	var recommends []SiteAppBook
	var newBooks []SiteAppBook
	var popularBooks []SiteAppBook

	o := orm.NewOrm()
	book := new(SiteAppBook)
	qs := o.QueryTable(book)
	// 查询数据库popularBooks按浏览量排序，newBooks按发布时间排序，recommends按时间排序
	_, popularBooksErr := qs.Filter("view__gt", 0).OrderBy("-view").Limit(10).All(&popularBooks)
	_, newBooksErr := qs.OrderBy("-created_at").Limit(10).All(&newBooks)
	_, recommendErr := qs.Filter("recommend", true).OrderBy("-created_at").Limit(10).All(&recommends)
	// 报错提示
	if recommendErr != nil {
		beego.Info("查询推荐列表出错了")
	}
	if newBooksErr != nil {
		beego.Info("查询最新列表出错了")
	}
	if popularBooksErr != nil {
		beego.Info("查询热门列表出错了")
	}
	return TodayInfo{
		Recommends:   recommends,
		NewBooks:     newBooks,
		PopularBooks: popularBooks}
}

// 推荐列表（全部）
func Recommend(index int, size int) PageData {
	var data []SiteAppBook
	o := orm.NewOrm()
	book := new(SiteAppBook)
	qs := o.QueryTable(book)
	// 分页
	_, error := qs.Filter("recommend",true).Limit(size, (index-1)*size).All(&data)
	// 查询总的条数
	totalElements, _ := qs.Filter("recommend", true).Count()
	if error != nil {
		beego.Info("查询全部推荐列表出错")
	}
	return PageData{
		Data:data,
		Index:index,
		TotalElements:totalElements,
		Size:size}
}

// 新书推荐（全部）
func NewBook(index int, size int) PageData {
	var data []SiteAppBook
	o := orm.NewOrm()
	book := new(SiteAppBook)
	qs := o.QueryTable(book)
	// 分页
	_, error := qs.OrderBy("-created_at").Limit(size, (index-1)*size).All(&data)
	// 查询总的条数
	totalElements, _ := qs.Count()
	if error != nil {
		beego.Info("查询全部推荐列表出错")
	}
	return PageData{
		Data:data,
		Index:index,
		TotalElements:totalElements,
		Size:size}
}

// 最受欢迎的书籍（全部）
func PopularBook(index int, size int) PageData {
	var data []SiteAppBook
	o := orm.NewOrm()
	book := new(SiteAppBook)
	qs := o.QueryTable(book)
	// 分页
	_, error := qs.OrderBy("-view").Limit(size, (index-1)*size).All(&data)
	// 查询总的条数
	totalElements, _ := qs.Count()
	if error != nil {
		beego.Info("查询全部推荐列表出错")
	}
	return PageData{
		Data:data,
		Index:index,
		TotalElements:totalElements,
		Size:size}
}

// 书籍详情
func BookDetail(bookId int) SiteAppBook {
	o := orm.NewOrm()
	data := SiteAppBook{Id: int(bookId)}
	err := o.Read(&data, "Id")
	if err != nil {
		beego.Info("查询书籍详情出错")
	}
	return data
}
