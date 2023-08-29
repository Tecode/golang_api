package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// SendEmailCode ...
// @Title 发送邮箱验证码
// @Description create UserRelated
// @Param	body	body 	models.SendCode	true		"body for UserRelated content"
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
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, valid.Errors[0].Message)
		return
	}

	// 查询邮箱是否已经被注册
	if account, err := models.GetUsersByAccount(requestBody.Email); err == nil && account != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, "The email address is already registered")
		return
	}

	// 发送邮件
	emailError := utils.SendCaptchaCode(requestBody.Email, c.Ctx.Input.IP())
	if emailError != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, emailError.Error())
		return
	}
	utils.RequestOutInput(c.Ctx, 200, 200200, nil, "邮件发送成功，5分钟内有效")

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

// Put ...
// @Title Put
// @Description update the UserRelated
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body	 	models.UserRelated	true		"body for UserRelated content"
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
		utils.RequestOutInput(c.Ctx, 500, 500500, nil, jsonErr.Error())
		logs.Error(jsonErr)
		return
	}
	valid := validation.Validation{}
	pass, _ := valid.Valid(&userData)
	if !pass {
		for _, err := range valid.Errors {
			utils.RequestOutInput(c.Ctx, 400, 400400, nil, err.Error())
			break
		}
		return
	}
	cacheCode, ok := utils.GetCode(userData.Email)
	if !ok || (ok && cacheCode.(string) != userData.Code) {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, "The email verification code is incorrect")
		return
	}
	// 添加判断是否已经存在
	userId, addUserError := models.AddUser(&userData)
	if addUserError != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, addUserError.Error())
		return
	}
	utils.RequestOutInput(c.Ctx, 400, 400400,
		map[string]any{
			"userId":     userId,
			"createTime": strconv.FormatInt(time.Now().UnixMilli(), 10),
			//"userName": userData.Name,
			//"id":       strconv.FormatInt(userData.Id, 10),
		}, "Register successes")
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
// @Param	body		body	models.LoginModel	true		"body for Haouxuan content"
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
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, valid.Errors[0].Message)
		return
	}
	// 添加判断是否已经存在
	userData, searchUserError := models.LoginAccount(requestBody.Account, requestBody.Password)
	if searchUserError != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, "Account or password error")
		return
	}
	token, tokenError := utils.CreateToken(userData.Id)
	if tokenError != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, tokenError.Error())
		return
	}
	// cookie保存一个token
	c.Ctx.SetCookie("token", token, 60*30)
	utils.RequestOutInput(c.Ctx, 200, 200200,
		map[string]string{
			"token":      token,
			"createTime": strconv.FormatInt(time.Now().UnixMilli(), 10),
			//"userName": userData.Name,
			//"id":       strconv.FormatInt(userData.Id, 10),
		}, "success")
}

// RecordAccount 记录邮箱账号
// @Title 发送邮箱验证码
// @Description create UserRelated
// @Param	body    body	models.SendCode	true		"body for UserRelated content"
// @Success 201 {object} {code: 200}
// @Failure 403 body is empty
// @router /record-account [post]
func (c *UserRelatedController) RecordAccount() {
	// 发送socket消息
	//GetSocket().hub.SendMessage()
	// 获取body的json数据
	requestBody := models.SendCode{}
	if jsonErr := c.BindJSON(&requestBody); jsonErr != nil {
		logs.Error(jsonErr.Error())
	}
	valid := validation.Validation{}
	valid.Email(requestBody.Email, "email")
	if valid.HasErrors() {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, valid.Errors[0].Message)
		return
	}

	// 查询邮箱是否已经被注册，没有被注册的邮箱不能发送重置密码邮件
	account, err := models.GetUsersByAccount(requestBody.Email)
	if account == nil || err != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, "The email address is not registered")
		return
	}
	// 查询是否账号已经被记录
	if record, _ := models.GetRecordAccountByEmail(requestBody.Email); record != nil {
		dateTime := record.CreatedAt
		duration := time.Now().Sub(dateTime)
		fmt.Println(duration.Hours(), dateTime, time.Now())
		if record.Valid && duration.Minutes() < 15 {
			utils.RequestOutInput(c.Ctx, 400, 400400, nil, "The email is valid for 15 minutes after it was sent")
			return
		}
		// 15分钟后让邮件失效
		if duration.Minutes() > 15 {
			err := models.UpdateRecordAccountById(record)
			if err != nil {
				logs.Error(err.Error())
				utils.RequestOutInput(c.Ctx, 400, 400400, nil, err.Error())
				return
			}
		}
	}
	// 生成MD5验证码
	hash := md5.New()
	token := hex.EncodeToString(hash.Sum([]byte(utils.GenerateRandomString(8))))
	content, err := utils.RenderTemplate[string]("resource/docs/reset_password.html", &token)
	if err != nil {
		logs.Error(err)
	}
	// 记录修改的账户以及验证token
	m := models.RecordAccount{
		UserId: account.Id,
		Valid:  true,
		RecordFiled: models.RecordFiled{
			Email:       requestBody.Email,
			RecordToken: token,
		},
	}
	if _, err := models.AddRecordAccount(&m); err != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, "Insert db error")
		return
	}
	// 发送邮件
	emailError := utils.SendEmail(requestBody.Email, "重置密码", content)
	if emailError != nil {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, emailError.Error())
		return
	}
	utils.RequestOutInput(c.Ctx, 200, 200200, nil, "The email is valid for 15 minutes after it was sent")
}

// ResetPassword 重置密码
// @Title 重置密码，需要邮箱和token验证
// @Description create UserRelated
// @Param	body	body 	models.ResetPasswordType	true		"body for ResetPasswordType content"
// @Success 201 {object} {code: 200}
// @Failure 403 body is empty
// @router /reset-password [post]
func (c *UserRelatedController) ResetPassword() {
	// 获取body的json数据
	requestBody := models.ResetPasswordType{}
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
	//	密码和确认密码是否相同
	if requestBody.Password != requestBody.ConfirmPassword {
		utils.RequestOutInput(c.Ctx, 400, 400400, nil, "密码和确认密码不相同")
		return
	}
	// 查找是否有重置密码的记录
	if v, _ := models.GetRecordAccountByToken(&requestBody); v != nil {
		// 设置token无效
		err := models.UpdateRecordAccountById(v)
		if err != nil {
			logs.Error(err.Error())
		}
		// 重置密码
		updateErr := models.UpdateUsersById(&models.Users{
			Id:        v.UserId,
			UserFiled: models.UserFiled{Password: requestBody.Password},
		}, "Password")
		if updateErr != nil {
			return
		}
		utils.RequestOutInput(c.Ctx, 200, 200200, nil, "密码重置成功")
		return
	}
	utils.RequestOutInput(c.Ctx, 400, 400400, nil, "重置密码失败，请重新发送邮件")
}
