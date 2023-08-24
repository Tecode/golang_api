package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type RecordFiled struct {
	Email       string `orm:"size(128)" json:"email"`
	RecordToken string `orm:"size(128)" json:"recordToken"`
}

type RecordAccount struct {
	Id     int64 `orm:"auto_now_add" json:"id"` // id
	UserId int64 `json:"userId"`
	RecordFiled
	Valid     bool      `orm:"default(true)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"` // 注册时间
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`     // 更新时间
}

// AddRecordAccount insert a new RecordAccount into database and returns
// last inserted id on success.
func AddRecordAccount(m *RecordAccount) (id int64, err error) {
	var data RecordAccount
	o := orm.NewOrm()
	data.Email = m.RecordFiled.Email
	data.RecordToken = m.RecordFiled.RecordToken
	data.Valid = true
	data.UserId = m.UserId
	id, err = o.Insert(&data)
	return
}

// GetRecordAccountByEmail retrieves RecordAccount by id. Returns error if
// id doesn't exist
func GetRecordAccountByEmail(email string) (v *RecordAccount, err error) {
	o := orm.NewOrm()
	v = &RecordAccount{}
	if err = o.QueryTable(new(RecordAccount)).Filter("Valid", true).Filter("Email", email).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetRecordAccountByToken 跟具邮箱token查找是否点击过重置密码
func GetRecordAccountByToken(m *ResetPasswordType) (v *RecordAccount, err error) {
	o := orm.NewOrm()
	v = &RecordAccount{}
	if err = o.QueryTable(new(RecordAccount)).Filter("Email", m.RecordToken).Filter("Valid", true).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateRecordAccountById UpdateRecordAccount updates RecordAccount by Id and returns error if
// the record to be updated doesn't exist
func UpdateRecordAccountById(m *RecordAccount) (err error) {
	o := orm.NewOrm()
	// ascertain id exists in the database
	if err = o.Read(&RecordAccount{Id: m.Id}); err == nil {
		var num int64
		// 设置为无效
		m.Valid = false
		if num, err = o.Update(m); err == nil {
			logs.Info(fmt.Println("Number of records updated in database:", num))
		}
	}
	return
}

// DeleteRecordAccount deletes RecordAccount by id and returns error if
// the record to be deleted doesn't exist
func DeleteRecordAccount(id int64) (err error) {
	o := orm.NewOrm()
	v := RecordAccount{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RecordAccount{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
