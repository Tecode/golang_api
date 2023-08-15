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
		return nil
	})
	// 加入全局的计划任务列表
	task.AddTask("emailTask", emailTask)
	// 开始执行全局的任务
	task.StartTask()
}
