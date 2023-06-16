package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"regexp"
	"time"
)

type TokenValidController struct {
	beego.Controller
}

// MyCustomClaims jwt
type MyCustomClaims struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	jwt.StandardClaims
}

// Filter 验证token，token出错就返回错误信息
var Filter = func(ctx *context.Context) {
	// 目前暂时不需要token
	return
	fmt.Println(ctx.Input.URL() == "/v1/user/", ctx.Input.URL(), "ctx=============================")
	beego.Info("路由拦截")
	token := ctx.Request.Header["Token"]
	// 注册账号（过滤掉）
	if ctx.Input.Method() == "POST" && ctx.Input.URL() == "/v1/user/" {
		return
	}
	// /picture专门处理图片的，裁剪一类（过滤掉）
	public, _ := regexp.MatchString("/public/", ctx.Request.RequestURI)
	if public {
		return
	}
	//beego.Info(token, ctx.Request.RequestURI, "----------")
	match, _ := regexp.MatchString("/user/login", ctx.Request.RequestURI)
	if !match {
		if len(token) == 0 {
			ctx.Output.SetStatus(400)
			ctx.Output.JSON(
				map[string]interface{}{"code": 400404, "message": "Token不存在"},
				true,
				false)
		} else {
			mySigningKey := []byte("7e6c8b94a77412")
			// Verifying and validating a JWT
			token, err := jwt.Parse(token[0], func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return mySigningKey, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// 修改时间戳，生成新的token,30分钟过期
				claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
				newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				newTokenString, _ := newToken.SignedString(mySigningKey)
				fmt.Println(newTokenString, "认证成功,新的token")
				//tokenString := CreateToken(string(claims["name"]), claims["id"])
				// 将新的token放在cookie返回
				ctx.SetCookie("token", newTokenString)
			} else {
				ctx.Output.SetStatus(400)
				ctx.Output.JSON(err.Error(), true, false)
			}
		}
	}
}

// 生成token，过期时间为30分钟
func CreateToken(email string, id int) string {
	// create json web token
	mySigningKey := []byte("7e6c8b94a77412")
	// Create the Claims
	claims := MyCustomClaims{
		email,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Id:        "100030",
			//IssuedAt:  now.Unix(),
			Issuer: "bandzest-auth",
			//NotBefore: now.Add(30 * time.Minute).Unix(),
			Subject: "someone",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(mySigningKey)
	return tokenString
}

// 获取token里面的信息
func Token(tokenString string) (email string, id float64) {
	mySigningKey := []byte("7e6c8b94a77412")
	token, err := jwt.Parse(tokenString, func(readToken *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := readToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", readToken.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["name"].(string), claims["id"].(float64)
	} else {
		fmt.Println(err)
	}
	return "", 0
}
