package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/models"
	"golang_apiv2/utils"
)

// UserRelatedController operations for UserRelated
type UserRelatedController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserRelatedController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create UserRelated
// @Param	body		body 	models.UserRelated	true		"body for UserRelated content"
// @Success 201 {object} models.UserRelated
// @Failure 403 body is empty
// @router / [post]
func (c *UserRelatedController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get UserRelated by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserRelated
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserRelatedController) GetOne() {
	c.Data["json"] = map[string]string{"GetOne": c.Ctx.Input.Param(":haoxuan"), "appname": "appName"}
	err := c.ServeJSON()
	if err != nil {
		return
	}
}

// GetAll ...
// @Title GetAll
// @Description get UserRelated
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UserRelated
// @Failure 403
// @router / [get]
func (c *UserRelatedController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the UserRelated
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UserRelated	true		"body for UserRelated content"
// @Success 200 {object} models.UserRelated
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserRelatedController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the UserRelated
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserRelatedController) Delete() {

}

// UserLogin ...
// @Title 用户登录
// @Description 登录后获取token
// @Param	body		body 	models.LoginModel	true		"body for Haouxuan content"
// @Success 201 {object} models.UserRelated
// @Failure 403 body is empty
// @router /login [post]
func (c *UserRelatedController) UserLogin() {
	token, tokenError := utils.CreateToken()
	if tokenError != nil {
		c.Ctx.Output.SetStatus(400)
		err := c.Ctx.Output.JSON(
			models.ResponseData[map[string]string]{
				Code:    400400,
				Data:    nil,
				Message: tokenError.Error(),
			},
			false,
			false,
		)
		if err != nil {
			logs.Error(err.Error())
		}
		return
	}
	// 获取body的json数据
	data := models.LoginModel{}
	jsonErr := c.BindJSON(&data)
	if jsonErr != nil {
		return
	}
	err := c.Ctx.Output.JSON(
		models.ResponseData[map[string]string]{
			Code:    200200,
			Data:    map[string]string{"token": token, "userName": data.Account, "password": data.Password},
			Message: "获取数据成功",
		},
		false,
		false,
	)
	if err != nil {
		logs.Error(err)
		return
	}
}