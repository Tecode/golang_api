package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"golang_apiv2/utils"
	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

type UserRelated struct {
	Id int `json:"id"`
}

type Haouxuan struct {
	Id int `json:"id"`
}

type LoginModel struct {
	Account  string `valid:"Email" json:"account"`
	Password string `valid:"Required" json:"password"`
}

func init() {
	// 设置数据库的基本信息,golang_apiv2数据库名称
	// SiteAppUser->site_app_user
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",
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
	orm.RegisterModel(new(RecordAccount))
	if utils.GetAppConfigValue("createtable") == "true" {
		err := orm.RunSyncdb("default", false, true)
		if err != nil {
			logs.Error(err.Error())
			return
		}
		logs.Info("🚀🚀创建表成功")
	}
}
