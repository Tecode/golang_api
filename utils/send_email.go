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

func SendEmail(email string) error {
	foo, found := c.Get(email)
	if found {
		fmt.Println(foo)
		return fmt.Errorf("邮箱已发送，5分钟内有效")
	}
	content, _ := os.ReadFile("resource/docs/captcha_template.html")
	// 邮箱验证码
	emailCode := GenerateCode()
	c.Set("283731869@qq.com", emailCode, cache.DefaultExpiration)
	emailContent := fmt.Sprintf(string(content), "您好", email, emailCode)
	m := gomail.NewMessage()
	m.SetHeader("From", "283731869@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "登录/注册邮箱验证码")
	m.SetBody("text/html", emailContent)

	port, err := strconv.ParseInt(GetAppConfigValue("emailport"), 10, 64)
	if err != nil {
		logs.Error(err.Error())
		return err
	}

	dialer := gomail.NewDialer(
		GetAppConfigValue("emailhost"),
		int(port),
		GetAppConfigValue("emailaccount"),
		GetAppConfigValue("emailpassword"),
	)

	if err := dialer.DialAndSend(m); err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
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
