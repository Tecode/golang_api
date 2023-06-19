package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"golang_apiv2/utils"
	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// ResponseData 返回的数据统一格式
type ResponseData[T interface{}] struct {
	Code    int64  `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type UserRelated struct {
	Id int `json:"id"`
}

type Haouxuan struct {
	Id int `json:"id"`
}

type LoginModel struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func init() {
	// 设置数据库的基本信息,web_site_db数据库名称
	// SiteAppUser->site_app_user
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		utils.GetAppConfigValue("mysqlaccount"),
		utils.GetAppConfigValue("mysqlpassword"),
		utils.GetAppConfigValue("mysqlurl"),
		utils.GetAppConfigValue("msqldatabase"),
	)
	databaseErr := orm.RegisterDataBase(
		"default",
		"mysql",
		dataSource,
	)
	if databaseErr != nil {
		logs.Error(databaseErr.Error())
		return
	}
	logs.Info("🚀🚀数据库链接成功")
	// 创建表，注意！！这里要账号有创建数据库表的权限
	orm.RegisterModel(new(Users))
	if utils.GetAppConfigValue("createtable") == "true" {
		err := orm.RunSyncdb("default", false, true)
		if err != nil {
			logs.Error(err.Error())
			return
		}
		logs.Info("🚀🚀创建表成功")
	}
}
