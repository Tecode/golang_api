package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/models"
	"golang_apiv2/utils"
	"math"
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
	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(&r)
	if err != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, err.Error())
		return
	}
	if !b {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, valid.Errors[0].Error())
		return
	}
	//综合环境系数
	coefficient := r.WorkEnvironmentCoefficient * r.OppositeSex * r.ColleagueEnvironment
	//	工作性价比 = （平均日薪 x 综合环境系数）/ 35*（通勤时长 - 0.5*摸鱼时长） x 学历系数
	price := (r.AverageDailySalary * coefficient) / ((35 * (r.WorkingHours + r.CommutingHours - 0.5*r.IdleDuration)) * r.EducationCoefficient) * r.WorkTime
	result := math.Round(float64(price*100)) / 100
	fmt.Println(result)
	if result == 0 {
		utils.RequestOutInput(c.Ctx, 200, 200200, result, "太惨了")
		return
	}
	percent, insertErr := models.AddQualityWork(&r, float32(result))
	if insertErr != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, insertErr.Error())
		return
	}
	utils.RequestOutInput(
		c.Ctx, 200, 200200,
		map[string]any{
			"percent": utils.RoundToTwoDecimal(float64(percent)),
			"result":  result,
		},
		"OK",
	)
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
