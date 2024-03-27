package utils

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/patrickmn/go-cache"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

var c *cache.Cache = nil

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

// SendCaptchaCode 发送邮箱验证码
func SendCaptchaCode(email string, ip string) error {
	foo, found := GetCode(email)
	if found {
		logs.Info("Email", foo)
		return fmt.Errorf("邮件已发送，5分钟内有效")
	}
	// 邮箱验证码
	emailCode := GenerateCode()
	c.Set(email, emailCode, cache.DefaultExpiration)
	emailContent := RenderCaptcha(email, emailCode, ip)
	if err := SendEmail(email, "登录/注册邮箱验证码", emailContent); err != nil {
		return err
	}
	return nil
}

// SendEmail 发送邮件
func SendEmail(email string, emailHeader string, emailContent string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "haoxuan1009@foxmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", emailHeader)
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
