package models

import "time"

// 推荐书籍
type BookInfo struct {
	Id          int       `json:"bookId"`      // id
	Title       string    `json:"title"`       // 书籍标题
	Description string    `json:"description"` // 书籍的描述
	Author      string    `json:"author"`      // 作者
	BookImage   string    `json:"bookImage"`   // 书籍封面
	Content     string    `json:"content"`     // 书籍内容
	Score       uint      `json:"score"`       // 评分
	Delete      bool      `json:"delete"`      // 是否下架
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
	View        int32     `json:"view"`        //浏览次数
}
