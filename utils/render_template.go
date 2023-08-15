package utils

import (
	"fmt"
	"html/template"
	"strings"
)

type LetterEmail struct {
	Name string
	Url  string
}

func Render() {
	temp, err := template.ParseFiles("resource/docs/letter_email.html")
	if err != nil {
		return
	}
	// 使用 strings.Builder 来捕获渲染结果
	var resultBuilder strings.Builder
	if err := temp.Execute(&resultBuilder, "Gopher"); err != nil {
		return
	}
	// 获取渲染后的数据
	renderedData := resultBuilder.String()
	fmt.Println("Rendered data:", renderedData)
}
