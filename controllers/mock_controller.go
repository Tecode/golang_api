package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"golang_apiv2/utils"
)

// MockController operations for MockController
type MockController struct {
	beego.Controller
}

// MockGet mock请求数据的GET方法
func MockGet(ctx *context.Context) {
	jsonStr := `{"name": "haoxuan","age":20}`
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

// MockPost mock请求数据的POST方法
func MockPost(ctx *context.Context) {
	ctx.GetCookie("5556")
}

// MockDelete mock请求数据的DELETE方法
func MockDelete(ctx *context.Context) {
	ctx.GetCookie("5556")
}

// MockPut mock请求数据的PUT方法
func MockPut(ctx *context.Context) {
	ctx.GetCookie("5556")
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
// @Success 201 {object} models.MockBaseData
// @Failure 403 body is empty
// @router /add [post]
func (c *MockController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get MockController by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MockController
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MockController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get MockController
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MockController
// @Failure 403
// @router /getList [get]
func (c *MockController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the MockController
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.MockController	true		"body for MockController content"
// @Success 200 {object} models.MockController
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MockController) Put() {

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
