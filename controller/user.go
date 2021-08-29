// controller/user.go
package controller

import (
	"errors"
	"log"

	"github.com/EternalNight996/web-room/model"
	"github.com/EternalNight996/web-room/model/dbe"
	"github.com/EternalNight996/web-room/model/repo"
	"github.com/EternalNight996/web-room/model/reqo"
	"github.com/EternalNight996/web-room/tool"
	"github.com/kataras/iris/v12"
)

// PostLogin user login
func PostLogin(ctx iris.Context) {
	req := reqo.PostLogin{}
	// 写入登录请求
	ctx.ReadJSON(&req)

	// Query user by username
	user, err := userService.QueryByUsername(req.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		// code 2 查询数据库失败
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	log.Println(user, req)
	// If passwd are inconsistent
	if user.Passwd != req.Passwd {
		ctx.StatusCode(iris.StatusBadRequest)
		// code 6 数据验证失败
		ctx.JSON(model.ErrorVerification(errors.New("用户名或密码错误")))
		return
	}

	// Login Ok
	// Get token
	token, err := tool.GetJWTString(user.Username, user.ID)
	if err != nil {
		// code 500 服务器故障
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorBuildJWT(err))
	}

	rep := repo.PostLogin{
		Username: user.Username,
		ID:       user.ID,
		Token:    token,
	}
	ctx.JSON(rep)
}

// PostUser user register
func PostUser(ctx iris.Context) {
	req := reqo.PostUser{}
	ctx.ReadJSON(&req)

	// Username and passwd can't be blank
	if req.Username == "" || req.Passwd == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(errors.New("用户名和密码不能为空")))
		return
	}

	// Query user with same username
	exist, _ := userService.QueryByUsername(req.Username)

	// Can't be same username
	if exist.Username != "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(errors.New("用户名已存在")))
		return
	}

	// New user and insert into database
	newUser := dbe.User{
		Username: req.Username,
		Passwd:   req.Passwd,
		Gender:   req.Gender,
		Nickname: req.Nickname,
		Mail:     req.Mail,
	}
	userID, err := userService.Insert(newUser)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorInsertDatabase(err))
		return
	}

	rep := repo.PostUser{
		Username: newUser.Username,
		ID:       userID,
	}
	ctx.JSON(rep)
}

// GetUser return user list
func GetUser(ctx iris.Context) {
	req := reqo.GetUser{}
	ctx.ReadQuery(&req)
	repList := []repo.GetUser{}

	userList, err := userService.Query(req.Username, req.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	for _, user := range userList {
		rep := repo.GetUser{
			ID:       user.ID,
			Username: user.Username,
			Gender:   user.Gender,
			Nickname: user.Nickname,
			Mail:     user.Mail,
		}

		repList = append(repList, rep)
	}
	ctx.JSON(repList)
}

// PutUser update user information
func PutUser(ctx iris.Context) {
	req := reqo.PutUser{}
	ctx.ReadJSON(&req)
	logined := ctx.Values().Get("logined").(model.Logined)

	// // Query user by userID
	// user, err := userService.QueryByID(userID)
	// if err != nil {
	// 	ctx.JSON(new(model.RepModel).WithError(err.Error()))
	// 	return
	// }

	user := dbe.User{
		ID: logined.ID,
	}
	// Replace if set
	if req.Gender != 0 {
		user.Gender = req.Gender
	}
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Mail != "" {
		user.Mail = req.Mail
	}

	// Update user
	err := userService.Update(user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	// Get updated user
	updatedUser, err := userService.QueryByID(user.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	rep := repo.PutUser{
		ID:       updatedUser.ID,
		Username: updatedUser.Username,
		Gender:   updatedUser.Gender,
		Nickname: updatedUser.Nickname,
		Mail:     updatedUser.Mail,
	}
	ctx.JSON(rep)
}
