// controller/message.go
package controller

import (
	"github.com/EternalNight996/web-room/model"
	"github.com/EternalNight996/web-room/model/repo"
	"github.com/EternalNight996/web-room/model/reqo"
	"github.com/kataras/iris/v12"
)

// PostMessage send message
func PostMessage(ctx iris.Context) {
	req := reqo.PostMessage{}
	// 将ReceiverID和消息写人req
	ctx.ReadJSON(&req)
	// 获取用户id和用户名
	logined := ctx.Values().Get("logined").(model.Logined)

	// 先将消息、SenderID、ReceiverID等插入到数据库，并返回插入ID。
	insertID, err := messageService.Insert(logined.ID, req.ReceiverID, req.Content)
	if err != nil {
		// 写入错误状态码如404
		ctx.StatusCode(iris.StatusInternalServerError)
		// code 1,插入错误
		ctx.JSON(model.ErrorInsertDatabase(err))
		return
	}

	// 更新消息在数据库位置
	rep := repo.PostMessage{
		ID: insertID,
	}

	ctx.JSON(rep)
}

// GetMessage get all message
func GetMessage(ctx iris.Context) {
	req := reqo.GetMessage{}
	ctx.ReadQuery(&req)
	logined := ctx.Values().Get("logined").(model.Logined)

	msgList, err := messageService.Query(
		req.BeginID,
		req.BeginTime,
		req.EndTime,
		logined.ID,
	)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	// Build repponse object
	repList := []repo.GetMessage{}

	for _, msg := range msgList {
		private := false
		if msg.Receiver.ID != 0 {
			private = true
		}
		// Get single rep
		rep := repo.GetMessage{
			ID:         msg.Message.ID,
			SenderID:   msg.Message.SenderID,
			SenderName: msg.Sender.Username,
			Content:    msg.Message.Content,
			SendTime:   msg.Message.SendTime,
			Private:    private,
		}

		// Add into repList
		repList = append(repList, rep)
	}

	ctx.JSON(repList)
}
