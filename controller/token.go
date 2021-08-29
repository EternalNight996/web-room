// controller/token.go
package controller

import (
	"github.com/EternalNight996/web-room/model"
	"github.com/EternalNight996/web-room/model/repo"
	"github.com/kataras/iris/v12"
)

// GetTokenInfo 验证token是否有效，如果有效则返回token携带的信息
func GetTokenInfo(ctx iris.Context) {
	logined := ctx.Values().Get("logined").(model.Logined)

	rep := repo.GetTokenInfo{
		ID:       logined.ID,
		Username: logined.Username,
	}
	ctx.JSON(rep)
}
