package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type MockBaseData struct {
	Method     string `orm:"size(10)" json:"method" valid:"Required"`
	UserId     int64  `json:"userId" json:"userId" valid:"Required"`
	Path       string `orm:"size(128)" json:"path" valid:"Required"`
	StatusCode int64  `orm:"size(10)" json:"statusCode" valid:"Required"`
	Data       string `orm:"type(text)" json:"data" valid:"Required"`
}

type Mock struct {
	Id int64 `orm:"auto" json:"id"` // id
	MockBaseData
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"` // 注册时间
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`     // 更新时间
}

// AddMock insert a new Mock into database and returns
// last inserted I'd on success.
func AddMock(m *MockBaseData) (id int64, err error) {
	o := orm.NewOrm()
	exist := o.QueryTable(new(Mock)).Filter("Method", m.Method).Filter("path", m.Path).Exist()
	if exist {
		return 0, errors.New("api 已存在")
	}
	var data Mock
	data.MockBaseData.Data = m.Data
	data.MockBaseData.UserId = m.UserId
	data.MockBaseData.Path = m.Path
	data.MockBaseData.Method = strings.ToUpper(m.Method)
	data.MockBaseData.StatusCode = m.StatusCode
	id, err = o.Insert(&data)
	return
}

// GetMockByUrlAndMethod retrieves Mock by I'd. Returns error if
// I'd don't exist
func GetMockByUrlAndMethod(url string, method string) (v *Mock, err error) {
	o := orm.NewOrm()
	v = &Mock{}
	if err = o.QueryTable(new(Mock)).Filter("path", url).Filter("method", method).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetMockByPath retrieves Mock by path. Returns error if
func GetMockByPath(path string) (v *Mock, err error) {
	o := orm.NewOrm()
	v = &Mock{}
	if err = o.QueryTable(new(Mock)).Filter("Path", path).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetMockById retrieves Mock by I'd. Returns error if
// I'd doesn't exist
func GetMockById(id int64) (v *Mock, err error) {
	o := orm.NewOrm()
	v = &Mock{Id: id}
	if err = o.QueryTable(new(Mock)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMock retrieves all Mock matches certain condition. Returns empty list if
// no records exist
func GetAllMock(offset int64, limit int64) (ml map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Mock))

	var list []Mock
	qs = qs.OrderBy("id").RelatedSel()
	var listData []any
	if _, err = qs.Limit(offset, limit*offset).All(&list); err == nil {
		for _, v := range list {
			listData = append(listData, v)
		}
		ml = map[string]interface{}{}
		ml["list"] = listData
		count, err := qs.Count()
		if err != nil {
			return nil, err
		}
		ml["total"] = count
		return ml, nil
	}
	return nil, err
}

// UpdateMockById updates Mock by I'd and returns error if
// the record to be updated doesn't exist
func UpdateMockById(id int64, m *MockBaseData) (err error) {
	o := orm.NewOrm()
	v := Mock{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var data Mock
		data.Id = id
		data.MockBaseData.Data = m.Data
		data.MockBaseData.UserId = m.UserId
		data.MockBaseData.Path = m.Path
		// 都转换为大写进行保存
		data.MockBaseData.Method = strings.ToUpper(m.Method)
		data.MockBaseData.StatusCode = m.StatusCode
		if _, err = o.Update(&data, "Data", "UserId", "Path", "Method", "StatusCode"); err == nil {
		}
	}
	return
}

// DeleteMock deletes Mock by I'd and returns error if
// the record to be deleted doesn't exist
func DeleteMock(id int64) (err error) {
	o := orm.NewOrm()
	v := Mock{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Mock{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
