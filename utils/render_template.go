package utils

import (
	"encoding/json"
	"fmt"
	"github.com/6tail/lunar-go/calendar"
	"github.com/beego/beego/v2/core/logs"
	"html/template"
	"os"
	"strings"
	"time"
)

type LetterEmail struct {
	Name      string
	Url       string
	TodayDate string
	To        string
	Week      string
	Lunar     string
	Content   string
}

type CaptchaEmail struct {
	Title     string
	Email     string
	EmailCode string
	IP        string
}

type LetterContent struct {
	Content string `json:"content"`
	From    string `json:"from"`
}

type LetterData struct {
	UrlList     []string        `json:"urlList"`
	ContentList []LetterContent `json:"contentList"`
}

// RenderTemplate 渲染模板
func RenderTemplate[T any](fileName string, data *T) (string, error) {
	// "resource/docs/letter_email.html"
	temp, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	// 使用 strings.Builder 来捕获渲染结果
	var resultBuilder strings.Builder
	if err := temp.Execute(&resultBuilder, data); err != nil {
		return "", err
	}
	// 获取渲染后的数据
	renderedData := resultBuilder.String()
	fmt.Println("Rendered data:", renderedData)
	return renderedData, nil
}

// RenderEveryDayEmail 渲染每日邮件
func RenderEveryDayEmail() string {
	// 读取json文件的内容
	jsonFile, jsonError := os.Open("resource/docs/letter.json")
	if jsonError != nil {
		logs.Error(jsonError)
	}
	var letterData LetterData
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&letterData); err != nil {
		logs.Error(err)
		return ""
	}
	// 关闭文件
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			logs.Error(err)
		}
	}(jsonFile)
	//fmt.Println(letterData.ContentList, letterData.UrlList)
	//fmt.Println(RandomNumber(len(letterData.UrlList)))
	//随机获取图片和文字内容
	text := letterData.ContentList[RandomNumber(len(letterData.ContentList))].Content
	url := letterData.UrlList[RandomNumber(len(letterData.UrlList))]
	// 日期相关
	t := time.Now()
	todayDate := t.Format("2006-01-02")
	lunar := calendar.NewLunarFromDate(time.Now())
	lunarText := fmt.Sprintf("%s月%s", lunar.GetMonthInChinese(), lunar.GetDayInChinese())
	data := LetterEmail{Name: "Dear,洁洁(~^O^~)", Url: url, Lunar: lunarText, To: "洁洁", TodayDate: todayDate, Content: text}
	content, err := RenderTemplate[LetterEmail]("resource/docs/letter_email.html", &data)
	if err != nil {
		logs.Error(err)
		return ""
	}
	//fmt.Println(content, lunarText, todayDate)
	return content
}

// RenderCaptcha 渲染验证码邮件
func RenderCaptcha(email string, emailCode string, ip string) string {
	data := CaptchaEmail{Title: "你好", Email: email, EmailCode: emailCode, IP: ip}
	content, err := RenderTemplate[CaptchaEmail]("resource/docs/captcha_template.html", &data)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return content
}
