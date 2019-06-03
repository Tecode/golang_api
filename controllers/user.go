package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"golang_api/models"
)

// 用户相关
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.Registered	true		"body for user content"
// @Success 200 {int} models.SiteAppUser.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.Registered
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.Ctx.Output.SetStatus(400)
		u.Data["json"] = map[string]string{"uid": "参数错误"}
		u.ServeJSON()
		return
	}
	uid, insertErr := models.AddUser(user)
	if insertErr != nil {
		u.Ctx.Output.SetStatus(400)
		u.Data["json"] = map[string]interface{}{"error": insertErr}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Param	token		header 	string	true		"The token is required"
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.SiteAppUser	true		"body for user content"
// @Success 200 {object} models.SiteAppUser
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title 登录
// @Description 登录获取token username:admin password:123456
// @Param	body		body 	models.LoginInfo	true		"body for user content"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	//valid := validation.Validation{}
	var loginInfo models.LoginInfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &loginInfo)
	email := loginInfo.Email
	password := loginInfo.Password
	// 判断用户是否存在
	fmt.Println(email, password, "账号密码")
	if ok, userInfo := models.Login(email, password); ok {
		// 生成token
		tokenString := CreateToken(email, userInfo.Id)
		u.Data["json"] = map[string]interface{}{
			"code":  200200,
			"token": tokenString}
	} else {
		u.Ctx.Output.SetStatus(404)
		u.Data["json"] = map[string]interface{}{"code": 404404, "message": "用户不存在"}
	}
	u.ServeJSON()
}

// @Title 登出
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

type UserInfo struct {
	Name string  `json:"name"`
	Id   float64 `json:"userId"`
}

// @Title token验证
// @Description 验证token
// @Param	token		header 	string	true		"The token is required"
// @Success 200 {struct} {"name":"admin","id":"125"}
// @router /user_info [get]
func (u *UserController) UserInfo() {
	token := u.Ctx.Request.Header["Token"]
	_, id := Token(token[0])
	if len(token) == 0 {
		u.Ctx.Output.SetStatus(404)
		u.Data["json"] = map[string]interface{}{"code": 404404, "message": "token验证失败"}
		u.ServeJSON()
	} else {
		if ok, userInfo := models.GetUser(id); ok {
			fmt.Println(userInfo, id, "0==")
			u.Data["json"] = map[string]interface{}{
				"data": userInfo}
			u.Ctx.Output.SetStatus(200)
			u.ServeJSON()
		}
	}
}
