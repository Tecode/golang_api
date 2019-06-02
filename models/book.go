package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 推荐书籍
type TodayInfo struct {
	Recommends   []SiteAppBook `json:"recommendBooks"`
	NewBooks     []SiteAppBook `json:"newBooks"`
	PopularBooks []SiteAppBook `json:"popularBooks"`
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
