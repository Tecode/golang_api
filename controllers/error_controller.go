package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// ErrorControllerController operations for ErrorController
type ErrorControllerController struct {
	beego.Controller
}

// URLMapping ...
func (c *ErrorControllerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *ErrorControllerController) Error404() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "404.html"
}

// GetOne ...
// @Title GetOne
// @Description get ErrorController by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ErrorController
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ErrorControllerController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get ErrorController
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ErrorController
// @Failure 403
// @router / [get]
func (c *ErrorControllerController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the ErrorController
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ErrorController	true		"body for ErrorController content"
// @Success 200 {object} models.ErrorController
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ErrorControllerController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the ErrorController
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ErrorControllerController) Delete() {

}
