package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// HaouxuanController operations for Haouxuan
type HaouxuanController struct {
	beego.Controller
}

// Post ...
// @Title Create
// @Description create Haouxuan
// @Param	body		body 	models.Haouxuan	true		"body for Haouxuan content"
// @Success 201 {object} models.Haouxuan
// @Failure 403 body is empty
// @router / [post]
func (c *HaouxuanController) Post() {

}

// Get ...
// @Title GetOne
// @Description get Haouxuan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Haouxuan
// @Failure 403 :id is empty
// @router /:id [get]
func (c *HaouxuanController) Get() {
	appName, _ := beego.AppConfig.String("appname")
	c.Data["json"] = map[string]string{"GetOne": c.Ctx.Input.Param(":haoxuan"), "appname": appName}
	err := c.ServeJSON()
	if err != nil {
		return
	}
}

// GetAll ...
// @Title GetAll
// @Description get Haouxuan
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Haouxuan
// @Failure 403
// @router / [get]
func (c *HaouxuanController) GetAll() {
	c.Data["json"] = map[string]string{"GetAll": "GetAll"}
	err := c.ServeJSON()
	if err != nil {
		return
	}
}

// Put ...
// @Title Put
// @Description update the Haouxuan
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Haouxuan	true		"body for Haouxuan content"
// @Success 200 {object} models.Haouxuan
// @Failure 403 :id is not int
// @router /:id [put]
func (c *HaouxuanController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Haouxuan
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *HaouxuanController) Delete() {

}
