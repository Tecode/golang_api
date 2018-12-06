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

// 验证token，token出错就返回错误信息
var TokenValid = func(ctx *context.Context) {
	token := ctx.Request.Header["Token"]
	fmt.Println(token, ctx.Request.RequestURI)
	match, _ := regexp.MatchString("/user/login", ctx.Request.RequestURI)
	if !match {
		if len(token) == 0 {
			ctx.Output.SetStatus(404)
			ctx.Output.JSON("token不存在", true, false)
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
				fmt.Println(claims, "认证成功")
				//tokenString := CreateToken(string(claims["name"]), claims["id"])
				ctx.Output.Cookie("token", "newCookies")
			} else {
				ctx.Output.SetStatus(400)
				ctx.Output.JSON(err.Error(), true, false)
			}
		}
	}
}

// jwt
type MyCustomClaims struct {
	Name string `json:"name"`
	Id int `json:"id"`
	jwt.StandardClaims
}

// 生成token
func CreateToken(userName string, id int) string{
	// create json web token
	mySigningKey := []byte("7e6c8b94a77412")
	//now := time.Now()
	// Create the Claims
	claims := MyCustomClaims{
		userName,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
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