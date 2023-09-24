package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UserFiled struct {
	Name     string `orm:"size(40)" json:"name"`     // 姓名
	Nickname string `orm:"size(40)" json:"nickname"` // 昵称
	Gender   int    `orm:"null" json:"gender"`       // 性别
	Email    string `orm:"size(40)" json:"email"`    //邮箱账号
	Phone    string `orm:"size(40)" json:"phone"`    // 手机号码
	Password string `orm:"size(40)" json:"password"` // 密码
}

// Users 用户基本信息
type Users struct {
	Id int64 `orm:"auto" json:"id"` // id
	UserFiled
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"` // 注册时间
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`     // 更新时间
}

// SendCode 发送验证码
type SendCode struct {
	Email string `json:"email" valid:"Email"`
}

// ResetPasswordType 重置密码数据格式
type ResetPasswordType struct {
	Password        string `json:"password" valid:"Required"`
	ConfirmPassword string `json:"confirmPassword" valid:"Required"`
	RecordToken     string `json:"recordToken" valid:"Required"`
}

// RegisterUser 注册账号
type RegisterUser struct {
	Email    string `orm:"size(128)" json:"email" valid:"Email"`
	Password string `orm:"size(128)" json:"password" valid:"Required"`
	Code     string `orm:"size(128)" json:"code" valid:"Required"`
	Name     string `orm:"size(128)" json:"name" valid:"Match(/^.{2,10}$/)"`
}

// AddUser insert a new Users into database and returns
// last inserted I'd on success.
func AddUser(r *RegisterUser) (id int64, err error) {
	o := orm.NewOrm()
	exist := o.QueryTable(new(Users)).Filter("email", r.Email).Exist()
	if exist {
		return 0, errors.New("邮箱已存在")
	}
	var data Users
	data.Email = r.Email
	data.Nickname = r.Name
	data.Password = r.Password
	id, err = o.Insert(&data)
	return
}

// GetUsersByAccount retrieves Users by Id. Returns error if
// I'd doesn't exist
func GetUsersByAccount(email string) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{}
	if err = o.QueryTable(new(Users)).Filter("email", email).One(v); err == nil {
		return v, nil
	}
	logs.Error(err.Error())
	return nil, err
}

// LoginAccount retrieves Users by id. Returns error if
func LoginAccount(email string, password string) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{}
	if err = o.QueryTable(new(Users)).Filter("email", email).Filter("password", password).One(v); err == nil {
		return v, nil
	}
	logs.Error(err.Error())
	return nil, err
}

// UpdateUsersById UpdateUsers updates Users by id and returns error
// values:需要更新的列，默认会全部更新
func UpdateUsersById(m *Users, values ...string) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var u Users
		u.Id = m.Id
		u.Password = m.UserFiled.Password
		u.Gender = m.UserFiled.Gender
		u.Nickname = m.UserFiled.Nickname
		u.Phone = m.UserFiled.Phone
		u.Email = m.UserFiled.Email
		var num int64
		if num, err = o.Update(&u, values...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id int64) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
