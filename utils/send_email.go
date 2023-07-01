package utils

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/patrickmn/go-cache"
	"gopkg.in/gomail.v2"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var c *cache.Cache = nil

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func SendEmail() {
	foo, found := c.Get("283731869@qq.com")
	if found {
		fmt.Println(foo)
		return
	}
	content, _ := os.ReadFile("resource/docs/email_template.html")
	c.Set("283731869@qq.com", GenerateCode(), cache.DefaultExpiration)
	m := gomail.NewMessage()
	m.SetHeader("From", "283731869@qq.com")
	m.SetHeader("To", "283731869@qq.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", string(content))

	port, err := strconv.ParseInt(GetAppConfigValue("emailport"), 10, 64)
	if err != nil {
		logs.Error(err.Error())
		return
	}

	dialer := gomail.NewDialer(
		GetAppConfigValue("emailhost"),
		int(port),
		GetAppConfigValue("emailaccount"),
		GetAppConfigValue("emailpassword"),
	)

	if err := dialer.DialAndSend(m); err != nil {
		logs.Error(err.Error())
	}
}

// GenerateCode 生成随机验证码
func GenerateCode() string {
	min := 100000
	max := 999999
	randData := rand.NewSource(time.Now().UnixNano())
	code := rand.New(randData).Intn(max-min+1) + min
	return strconv.Itoa(code)
}

// GetCode 获取缓存的验证码
func GetCode(key string) (interface{}, bool) {
	return c.Get(key)
}
