package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/models"
	"golang_apiv2/utils"
	"strconv"
	"time"
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

// SendEmailCode ...
// @Title 发送邮箱验证码
// @Description create UserRelated
// @Param	body		body 	models.SendCode	true		"body for UserRelated content"
// @Success 201 {object} {code: 200}
// @Failure 403 body is empty
// @router /send-code [post]
func (c *UserRelatedController) SendEmailCode() {
	// 获取body的json数据
	requestBody := models.SendCode{}
	if jsonErr := c.BindJSON(&requestBody); jsonErr != nil {
		logs.Error(jsonErr.Error())
	}
	valid := validation.Validation{}
	valid.Email(requestBody.Email, "email")
	if valid.HasErrors() {
		c.Ctx.Output.SetStatus(400)
		if err := c.Ctx.Output.JSON(
			models.ResponseData{
				Code:    400400,
				Data:    valid.Errors[0],
				Message: "邮箱不符合校验规则",
			},
			false,
			false,
		); err != nil {
			logs.Error(err.Error())
		}
		return
	}
	// 发送邮件
	emailError := utils.SendCaptchaCode(requestBody.Email, c.Ctx.Input.IP())
	if emailError != nil {
		c.Ctx.Output.SetStatus(200)
		if err := c.Ctx.Output.JSON(
			models.ResponseData{
				Code:    200500,
				Data:    "邮件发送失败",
				Message: emailError.Error(),
			},
			false,
			false,
		); err != nil {
			logs.Error(err.Error())
		}
		return
	}
	if err := c.Ctx.Output.JSON(
		models.ResponseData{
			Code:    200200,
			Data:    nil,
			Message: "邮件发送成功，5分钟内有效",
		},
		false,
		false,
	); err != nil {
		logs.Error(err.Error())
	}
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

// Register ...
// @Title 用户注册
// @Description 登录后获取token
// @Param	body		body 	models.RegisterUser	true		"body for UserFiled content"
// @Success 201 {object} models.UserRelated
// @Failure 403 body is empty
// @router /register [post]
func (c *UserRelatedController) Register() {
	// 获取body的json数据
	userData := models.RegisterUser{}
	if jsonErr := c.BindJSON(&userData); jsonErr != nil {
		return
	}
	// 添加判断是否已经存在
	//userId, addUserError := models.AddUser(&userData)
	//if addUserError != nil {
	//	c.Ctx.Output.SetStatus(400)
	//	err := c.Ctx.Output.JSON(
	//		models.ResponseData{
	//			Code:    400400,
	//			Data:    nil,
	//			Message: addUserError.Error(),
	//		},
	//		false,
	//		false,
	//	)
	//	if err != nil {
	//		return
	//	}
	//	return
	//}
	logs.Info(userData.Code)
	//token, tokenError := utils.CreateToken(userId)
	//if tokenError != nil {
	//	c.Ctx.Output.SetStatus(400)
	//	err := c.Ctx.Output.JSON(
	//		models.ResponseData{
	//			Code:    400400,
	//			Data:    nil,
	//			Message: tokenError.Error(),
	//		},
	//		false,
	//		false,
	//	)
	//	if err != nil {
	//		logs.Error(err.Error())
	//	}
	//	return
	//}
	//// 获取body的json数据
	//data := models.LoginModel{}
	//jsonErr := c.BindJSON(&data)
	//if jsonErr != nil {
	//	return
	//}
	//err := c.Ctx.Output.JSON(
	//	models.ResponseData{
	//		Code:    200200,
	//		Data:    map[string]string{"token": token, "userName": data.Account, "password": data.Password},
	//		Message: "获取数据成功",
	//	},
	//	false,
	//	false,
	//)
	//if err != nil {
	//	logs.Error(err)
	//	return
	//}
}

// UserLogin ...
// @Title 用户登录
// @Description 登录后获取token
// @Param	body		body 	models.LoginModel	true		"body for Haouxuan content"
// @Success 201 {object} models.UserRelated
// @Failure 403 body is empty
// @router /login [post]
func (c *UserRelatedController) UserLogin() {
	// 获取body的json数据
	requestBody := models.LoginModel{}
	jsonErr := c.BindJSON(&requestBody)
	if jsonErr != nil {
		return
	}
	// 校验数据是否正确
	valid := validation.Validation{}
	//自定义错误信息
	//valid.SetError("Password", "c")
	pass, _ := valid.Valid(&requestBody)
	// 有错误
	if !pass {
		c.Ctx.Output.SetStatus(400)
		err := c.Ctx.Output.JSON(
			models.ResponseData{
				Code:    400400,
				Data:    nil,
				Message: valid.Errors[0].Message,
			},
			false,
			false,
		)
		if err != nil {
			return
		}
		return
	}
	// 添加判断是否已经存在
	userData, searchUserError := models.GetUsersByAccount(requestBody.Account, requestBody.Password)
	if searchUserError != nil {
		c.Ctx.Output.SetStatus(400)
		err := c.Ctx.Output.JSON(
			models.ResponseData{
				Code:    400400,
				Data:    nil,
				Message: "账号或密码错误！",
			},
			false,
			false,
		)
		if err != nil {
			return
		}
		return
	}
	token, tokenError := utils.CreateToken(userData.Id)
	if tokenError != nil {
		c.Ctx.Output.SetStatus(400)
		err := c.Ctx.Output.JSON(
			models.ResponseData{
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
	// cookie保存一个token
	c.Ctx.SetCookie("token", token, 60*30)
	if err := c.Ctx.Output.JSON(
		models.ResponseData{
			Code: 200200,
			Data: map[string]string{
				"token":      token,
				"createTime": strconv.FormatInt(time.Now().UnixMilli(), 10),
				//"userName": userData.Name,
				//"id":       strconv.FormatInt(userData.Id, 10),
			},
			Message: "获取数据成功",
		},
		false,
		false,
	); err != nil {
		logs.Error(err)
		return
	}
}
