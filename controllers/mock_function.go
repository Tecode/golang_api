package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web/context"
	"golang_apiv2/utils"
)

func MockGet(ctx *context.Context) {
	jsonStr := `{"name": "haoxuan","age":20}`
	var result map[string]any
	jsonErr := json.Unmarshal([]byte(jsonStr), &result)
	if jsonErr != nil {
		utils.RequestOutInput(ctx, 500, 500500, nil, jsonErr.Error())
		return
	}
	err := ctx.Output.JSON(result, false, true)
	if err != nil {
		utils.RequestOutInput(ctx, 500, 500500, nil, err.Error())
		return
	}
}

func MockPost(ctx *context.Context) {
	ctx.GetCookie("5556")
}

func MockDelete(ctx *context.Context) {
	ctx.GetCookie("5556")
}

func MockPut(ctx *context.Context) {
	ctx.GetCookie("5556")
}
