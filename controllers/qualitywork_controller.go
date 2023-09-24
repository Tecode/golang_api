package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/models"
	"golang_apiv2/utils"
	"math"
	"strconv"
)

// QualityWorkController operations for QualityWorkController
type QualityWorkController struct {
	beego.Controller
}

// URLMapping ...
func (c *QualityWorkController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// QualityWork 这B班上的值不值测算版
func (c *QualityWorkController) QualityWork() {
	c.TplName = "quality-work.html"
}

// Post ...
// @Title Create
// @Description create QualityWorkController
// @Param	body		body 	models.QualityWorkData	true		"body for QualityWorkController content"
// @Success 201 {object} models.QualityWorkController
// @Failure 403 body is empty
// @router /quality-work [post]
func (c *QualityWorkController) Post() {
	// 获取body的json数据
	r := models.QualityWorkData{}
	if jsonErr := c.BindJSON(&r); jsonErr != nil {
		logs.Error(jsonErr.Error())
	}
	//综合环境系数
	coefficient := r.WorkEnvironmentCoefficient * r.OppositeSex * r.ColleagueEnvironment
	//	工作性价比 = （平均日薪 x 综合环境系数）/ 35*（通勤时长 - 0.5*摸鱼时长） x 学历系数
	price := (r.AverageDailySalary * coefficient) / ((35 * (r.WorkingHours + r.CommutingHours - 0.5*r.IdleDuration)) * r.EducationCoefficient) * r.WorkTime
	result := strconv.FormatFloat(math.Round(float64(price)), 'f', 2, 64)
	fmt.Println(result)
	utils.RequestOutInput(c.Ctx, 200, 200200, result, "OK")
}

// GetOne ...
// @Title GetOne
// @Description get QualityWorkController by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.QualityWorkController
// @Failure 403 :id is empty
// @router /:id [get]
func (c *QualityWorkController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get QualityWorkController
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.QualityWorkController
// @Failure 403
// @router / [get]
func (c *QualityWorkController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the QualityWorkController
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.QualityWorkController	true		"body for QualityWorkController content"
// @Success 200 {object} models.QualityWorkController
// @Failure 403 :id is not int
// @router /:id [put]
func (c *QualityWorkController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the QualityWorkController
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *QualityWorkController) Delete() {

}
