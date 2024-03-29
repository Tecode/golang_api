package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type QualityWorkData struct {
	AverageDailySalary         float32 `json:"average_daily_salary" valid:"Required" orm:"description(平均日薪酬)"`
	WorkingHours               float32 `json:"working_hours" valid:"Required" orm:"description(工作时长)"`
	CommutingHours             float32 `json:"commuting_hours" valid:"Required" orm:"description(通勤时长)"`
	IdleDuration               float32 `json:"idle_duration" valid:"Required" orm:"description(摸鱼时长)"`
	EducationCoefficient       float32 `json:"education_coefficient" valid:"Required" orm:"description(学历系数)"`
	WorkEnvironmentCoefficient float32 `json:"work_environment_coefficient" valid:"Required" orm:"description(工作环境系数)"`
	OppositeSex                float32 `json:"opposite_sex" valid:"Required" orm:"description(异性环境系数)"`
	ColleagueEnvironment       float32 `json:"colleague_environment" valid:"Required" orm:"description(同事环境系数)"`
	ProfessionalQualifications float32 `json:"professional_qualifications" valid:"Required" orm:"description(职业资格系数)"`
	WorkTime                   float32 `json:"work_time" valid:"Required" orm:"description(是否8:30前上班)"`
}

type QualityWork struct {
	Id int64 `orm:"auto" json:"id"`
	QualityWorkData
	WorkScore float32   `json:"work_score" orm:"description(工作性价比)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"` // 注册时间
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`     // 更新时间
}

// AddQualityWork insert a new QualityWork into database and returns
// last inserted I'd on success.
func AddQualityWork(m *QualityWorkData, workScore float32) (id float32, err error) {
	o := orm.NewOrm()
	table := o.QueryTable(new(QualityWork))
	total, _ := table.Count()
	count, queryErr := table.Filter("WorkScore__lt", workScore).Count()
	if queryErr != nil {
		return 0, errors.New("查询失败")
	}
	var data QualityWork
	data.AverageDailySalary = m.AverageDailySalary
	data.WorkingHours = m.WorkingHours
	data.CommutingHours = m.CommutingHours
	data.IdleDuration = m.IdleDuration
	data.EducationCoefficient = m.EducationCoefficient
	data.WorkEnvironmentCoefficient = m.WorkEnvironmentCoefficient
	data.OppositeSex = m.OppositeSex
	data.ColleagueEnvironment = m.ColleagueEnvironment
	data.ProfessionalQualifications = m.ProfessionalQualifications
	data.WorkTime = m.WorkTime
	data.WorkScore = workScore
	_, err = o.Insert(&data)
	result := float32(count) / float32(total)
	return result, err
}

// GetQualityWorkById retrieves QualityWork by id. Returns error if
// Id doesn't exist
func GetQualityWorkById(id int64) (v *QualityWork, err error) {
	o := orm.NewOrm()
	v = &QualityWork{Id: id}
	if err = o.QueryTable(new(QualityWork)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllQualityWork retrieves all QualityWork matches certain condition. Returns empty list if
// no records exist
func GetAllQualityWork(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(QualityWork))
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

	var l []QualityWork
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

// UpdateQualityWork updates QualityWork by Id and returns error if
// the record to be updated doesn't exist
func UpdateQualityWorkById(m *QualityWork) (err error) {
	o := orm.NewOrm()
	v := QualityWork{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteQualityWork deletes QualityWork by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQualityWork(id int64) (err error) {
	o := orm.NewOrm()
	v := QualityWork{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&QualityWork{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
