package models

import (
	"errors"
	"fmt"
	"reflect"
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
	Id int64 `orm:"auto_now_add" json:"id"` // id
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
	data.MockBaseData.Method = m.Method
	data.MockBaseData.StatusCode = m.StatusCode
	id, err = o.Insert(&data)
	return
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
func GetAllMock(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Mock))
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

	var l []Mock
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
		data.MockBaseData.Method = m.Method
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
