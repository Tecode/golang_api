package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"reflect"
	"strings"
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
	Id        int64     `orm:"auto_now_add" json:"id"`                       // id
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"` // 注册时间
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`     // 更新时间
	UserFiled
}

// SendCode 发送验证码
type SendCode struct {
	Email string `json:"email"`
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

// LoginAccount retrieves Users by Id. Returns error if
func LoginAccount(email string, password string) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{}
	if err = o.QueryTable(new(Users)).Filter("email", email).Filter("password", password).One(v); err == nil {
		return v, nil
	}
	logs.Error(err.Error())
	return nil, err
}

// GetAllUsers retrieves all Users matches certain condition. Returns empty list if
// no records exist
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Users
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUsersById UpdateUsers updates Users by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
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
