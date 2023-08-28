package utils

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"mime/multipart"
	"net/http"
	"regexp"
	"time"
)

var (
	signingKey = []byte("7e6c8b94a77412")
	tracker    *IPTracker
)

// ResponseData 返回的数据统一格式
type ResponseData struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func init() {
	tracker = NewIPTracker()
}

// Interceptor 拦截器
func Interceptor(ctx *context.Context) {
	// 目前暂时不需要token
	fmt.Println(ctx.Input.URL() == "/v1/user/", ctx.Input.URL(), "ctx=============================")
	logs.Info("路由拦截")
	token := ctx.Request.Header["Token"]
	logs.Info(token, "ctx.Input.URL()")
	//// 登录、注册账号（过滤掉）
	//filterUrl, _ := regexp.MatchString("/login", ctx.Request.RequestURI)
	//logs.Info(ctx.Request.RequestURI, filterUrl, "filterUrl")
	//if ctx.Input.Method() == "POST" && filterUrl {
	//	return
	//}
	// /picture专门处理图片的，裁剪一类（过滤掉）
	//public, _ := regexp.MatchString("/public/", ctx.Request.RequestURI)
	//if public {
	//	return
	//}
	// beego.Info(token, ctx.Request.RequestURI, "----------")
	//match, _ := regexp.MatchString("/user/login", ctx.Request.RequestURI)
	//if !match {
	//
	//}
	if len(token) == 0 {
		ctx.Output.SetStatus(400)
		err := ctx.Output.JSON(
			map[string]interface{}{"code": 400404, "message": "Token不存在"},
			false,
			true,
		)
		if err != nil {
			return
		}
	} else {
		// Verifying and validating a JWT
		token, err := jwt.Parse(token[0], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return signingKey, nil
		})

		claims, isOk := token.Claims.(jwt.MapClaims)
		fmt.Println(claims, isOk, "claims")

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//// 修改时间戳，生成新的token,30分钟过期
			//claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
			//newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			//newTokenString, _ := newToken.SignedString(signingKey)
			//fmt.Println(newTokenString, "认证成功,新的token")
			////tokenString := CreateToken(string(claims["name"]), claims["id"])
			//// 将新的token放在cookie返回
			//ctx.SetCookie("token", newTokenString)
			return
		}
		// token无效或过期
		statusCode := 400400
		// token过期
		if err.Error() == "Token is expired" {
			statusCode = 400410
		}
		ctx.Output.SetStatus(400)
		jsonErr := ctx.Output.JSON(
			map[string]interface{}{"code": statusCode, "message": err.Error()},
			false,
			true,
		)
		if jsonErr != nil {
		}
	}
}

// CreateToken 生成token
func CreateToken(userId int64) (string, error) {
	//	登录操作，发放一个token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"iss":    "jwt-go",
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
	})
	tokenString, err := newToken.SignedString(signingKey)
	fmt.Println(tokenString, err, "newToken")
	//	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiaGFveHVhbiIsIm5iZiI6MTQ0NDQ3ODQwMH0.MG4hdXS5BX_HPgET36Dc6YgOuPjloAiG_M-DFHYT71I
	return tokenString, err
}

func TokenToUserId(jwtToken string) (int, error) {
	if len(jwtToken) < 1 {
		return 0, errors.New("token is invalid")
	}
	token, tokenError := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return signingKey, nil
	})
	// token有效就解析
	if tokenError != nil && token.Valid {
		claims, isOk := token.Claims.(jwt.MapClaims)
		fmt.Println(claims["userId"], isOk, "claims")
		return 0, nil
	}
	return 0, errors.New("token is invalid")
}

// GetAppConfigValue 获取appConfig的对应值
func GetAppConfigValue(key string) string {
	s, err := beego.AppConfig.String(key)
	if err != nil {
		return ""
	}
	return s
}

// RequestInterceptor 全局拦截器
func RequestInterceptor(ctx *context.Context) {
	if !tracker.TrackIP(ctx.Input.IP()) {
		ctx.Output.SetStatus(500)
		jsonErr := ctx.Output.JSON(
			map[string]interface{}{"code": 500500, "message": "Visits are too frequent"},
			false,
			true,
		)
		if jsonErr != nil {
			logs.Error(jsonErr.Error())
		}
	}
}

// RandomNumber 取随机数
func RandomNumber(count int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := rand.Intn(count)
	return randomNum
}

// RequestOutInput 返回响应的数据
func RequestOutInput(ctx *context.Context, status int, code int64, data any, message string) {
	ctx.Output.SetStatus(status)
	err := ctx.Output.JSON(
		ResponseData{
			Code:    code,
			Data:    data,
			Message: message,
		},
		false,
		false,
	)
	logs.Error(err)
}

// GenerateRandomString 获取随机字符串
func GenerateRandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[RandomNumber(len(charset))]
	}
	return string(result)
}

// GetFileExtension 获取文件后缀名
func GetFileExtension(fileHeader *multipart.FileHeader) string {
	file, err := fileHeader.Open()
	if err != nil {
		return ""
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	buf := make([]byte, 512) // Read the first 512 bytes
	_, err = file.Read(buf)
	if err != nil {
		return ""
	}

	contentType := http.DetectContentType(buf)
	switch contentType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "application/pdf":
		return ".pdf"
	// Add more cases for other file types
	default:
		pattern := `\.([a-zA-Z0-9]+)$`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(fileHeader.Filename)
		if len(matches) == 2 {
			extension := matches[1]
			return "." + extension
		}
		return ""
	}
}
