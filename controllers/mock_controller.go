package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"golang_apiv2/models"
	"golang_apiv2/utils"
	"strconv"
)

// MockController operations for MockController
type MockController struct {
	beego.Controller
}

// MockAnyMethod mock请求数据的GET方法
func MockAnyMethod(ctx *context.Context) {

	jsonStr := "{\"name\":\"haoxuan\",\"code\":200200,\"status\":200,\"path\":\"/po/auth/name/data\",\"to\":\"成都\"}"
	var result map[string]any
	jsonErr := json.Unmarshal([]byte(jsonStr), &result)
	if jsonErr != nil {
		utils.RequestOutInput(ctx, 500, 500500, nil, jsonErr.Error())
		return
	}
	err := ctx.Output.JSON(result, false, true)
	if err != nil {
		utils.RequestOutInput(ctx, 500, 500500, nil, err.Error())
		return
	}
}

// URLMapping ...
func (c *MockController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create MockController
// @Param	body		body 	models.MockBaseData	true		"body for MockController content"
// @Success 201 {object} {}
// @Failure 403 body is empty
// @router /add [post]
func (c *MockController) Post() {
	// 获取body的json数据
	requestBody := models.MockBaseData{}
	if jsonErr := c.BindJSON(&requestBody); jsonErr != nil {
		logs.Error(jsonErr.Error())
	}
	valid := validation.Validation{}
	pass, _ := valid.Valid(requestBody)
	if !pass {
		for _, err := range valid.Errors {
			utils.RequestOutInput(c.Ctx, 400, 400400, nil, err.Message)
			break
		}
		return
	}
	mock, mockErr := models.AddMock(&requestBody)
	if mockErr != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, mock, mockErr.Error())
		return
	}
	utils.RequestOutInput(c.Ctx, 200, 200200, mock, "添加成功")
}

// GetOne ...
// @Title 获取mock 接口详情
// @Description 获取mock 接口详情
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MockController
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MockController) GetOne() {
	intValue, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, err.Error(), "参数错误")
		return
	}
	data, findError := models.GetMockById(intValue)
	if err != nil || data == nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, findError.Error())
		return
	}
	utils.RequestOutInput(c.Ctx, 200, 200200, *data, "查找成功")
}

// GetAll ...
// @Title 获取全部的mock api 接口
// @Description get MockController
// @Param	page	query	string	false	"第几页 例如：1"
// @Param	size	query	string	false	"每页条数 例如：10"
// @Success 200 {object} models.MockController
// @Failure 403
// @router /getList [get]
func (c *MockController) GetAll() {
	listData, err := models.GetAllMock(0, 10)
	if err != nil {
		return
	}
	utils.RequestOutInput(c.Ctx, 200, 200200, listData, "查找成功")
}

// Put ...
// @Title 修改api接口信息
// @Description 修改mock api接口信息
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.MockBaseData	true		"body for MockController content"
// @Success 200 {object} models.MockController
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MockController) Put() {
	// 获取body的json数据
	requestBody := models.MockBaseData{}
	if jsonErr := c.BindJSON(&requestBody); jsonErr != nil {
		logs.Error(jsonErr.Error())
	}
	valid := validation.Validation{}
	pass, _ := valid.Valid(requestBody)
	if !pass {
		for _, err := range valid.Errors {
			utils.RequestOutInput(c.Ctx, 400, 400400, nil, err.Message)
			break
		}
		return
	}
	intValue, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	if err != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, err.Error(), "参数错误")
		return
	}
	updateErr := models.UpdateMockById(intValue, &requestBody)
	if updateErr != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, updateErr.Error())
		return
	}
	utils.RequestOutInput(c.Ctx, 200, 200200, "mock", "添加成功")
}

// Delete ...
// @Title Delete
// @Description delete the MockController
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MockController) Delete() {

}