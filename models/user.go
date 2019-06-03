package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	UserList map[string]*User
)

func init() {}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

type Registered struct {
	Name     string `json:"name"`     // 姓名
	Email    string `json:"email"`    //邮箱账号
	Phone    string `json:"phone"`    // 手机号码
	Password string `json:"password"` // 密码
}

// 登录信息
type LoginInfo struct {
	Email    string `json:"email";valid:"Email; MaxSize(100)"`
	Password string `json:"password";valid:"Required"`
}

func AddUser(u Registered) (int64, error) {
	o := orm.NewOrm()
	// 邮箱是否存在
	exist := o.QueryTable("site_app_user").Filter("email", u.Email).Exist()
	if exist {
		return 0, errors.New("邮箱已存在")
	}
	var user SiteAppUser
	user.Email = u.Email
	user.Name = u.Name
	user.Password = u.Password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info(err)
		return id, err
	}
	return id, nil
}

func GetUser(uid float64) (ok bool, userInfo SiteAppUser) {
	o := orm.NewOrm()
	user := SiteAppUser{Id: int(uid)}
	err := o.Read(&user, "Id")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return false, user
	}
	if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return false, user
	}
	return true, user
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("用户不存在")
}

func Login(email, password string) (ok bool, userInfo SiteAppUser) {
	o := orm.NewOrm()
	user := SiteAppUser{Email: email}
	err := o.Read(&user, "Email")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return false, user
	}
	if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return false, user
	}
	return true, user
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
