package utils

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

func RunTimedTask() {
	// 定时发送哟件每天 06:06 执行
	emailTask := task.NewTask("task001", "0 06 06 * * *", func(ctx context.Context) error {
		logs.Info("定时任务 0 17 20")
		emailContent := RenderEveryDayEmail()
		if err01 := SendEmail("283731869@qq.com", "元气满满的一天开始啦，要开心噢（づ￣ 3￣)づ", emailContent); err01 != nil {
			return err01
		}
		if err02 := SendEmail("964856415@qq.com", "元气满满的一天开始啦，要开心噢（づ￣ 3￣)づ", emailContent); err02 != nil {
			return err02
		}
		return nil
	})
	// 加入全局的计划任务列表
	task.AddTask("emailTask", emailTask)
	// 开始执行全局的任务
	task.StartTask()
}
