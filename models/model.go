package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// SiteAppUser 用户基本信息
type SiteAppUser struct {
	Id        int       `orm:"auto_now_add"json:"id"`                       // id
	Name      string    `orm:"size(40)"json:"name"`                         // 姓名
	Nickname  string    `orm:"size(40)"json:"nickname"`                     // 昵称
	Email     string    `json:"email"`                                      //邮箱账号
	Phone     string    `json:"phone"`                                      // 手机号码
	Password  string    `json:"password"`                                   // 密码
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"json:"createdAt"` // 注册时间
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"json:"updatedAt"`     // 更新时间
}

// SiteAppBook 书籍基本信息
type SiteAppBook struct {
	Id          int       `orm:"auto_now_add"json:"id"`                       // id
	Title       string    `orm:"size(60)"json:"title"`                        // 书籍标题
	Description string    `orm:"size(500)"json:"description"`                 // 书籍的描述
	Author      string    `json:"author"`                                     // 作者
	BookImage   string    `json:"bookImage"`                                  // 书籍封面
	Content     string    `orm:"type(text)"json:"content"`                    // 书籍内容
	Score       uint      `json:"score"`                                      // 评分
	Delete      bool      `orm:"default(false)"json:"delete"`                 // 是否下架
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"json:"createdAt"` // 创建时间
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"json:"updatedAt"`     // 更新时间
	View        int32     `json:"view"`                                       //浏览次数
	Recommend   bool      `orm:"default(false)"json:"recommend"`              // 推荐
}

func init() {
	// 设置数据库的基本信息,web_site_db数据库名称
	// SiteAppUser->site_app_user
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/web_site_db?charset=utf8")
	orm.RegisterModel(
		new(SiteAppUser),
		new(SiteAppBook))
	//orm.RunSyncdb("default", false, true)
}
