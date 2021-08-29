// service/user.go
package service

import (
	"github.com/EternalNight996/web-room/model/dbe"
	"github.com/go-xorm/xorm"
)

// UserService user service
type UserService struct {
	db *xorm.Engine
}

// Query query user by username and id
func (userService *UserService) Query(username string, id uint) ([]dbe.User, error) {
	var userList []dbe.User

	// Limit username
	tmpDB := userService.db.Where("username like ?", "%"+username+"%")

	// Limit id
	if id != 0 {
		tmpDB.Where("id = ?", id)
	}

	// Execute query
	if err := tmpDB.Find(&userList); err != nil {
		return nil, err
	}

	return userList, nil
}

// QueryByUsername return one user
func (userService *UserService) QueryByUsername(username string) (dbe.User, error) {
	var user = dbe.User{
		Username: username,
	}
	has, err := userService.db.Get(&user)
	if err != nil {
		return dbe.User{}, err
	}
	if !has {
		return dbe.User{}, nil
	}
	return user, nil
}

// QueryByID return one user
func (userService *UserService) QueryByID(id int64) (dbe.User, error) {
	var user = dbe.User{
		ID: id,
	}

	if _, err := userService.db.Get(&user); err != nil {
		return dbe.User{}, err
	}

	return user, nil
}

// Insert insert a new user and return id
func (userService *UserService) Insert(user dbe.User) (int64, error) {
	if _, err := userService.db.Insert(&user); err != nil {
		return 0, err
	}
	return user.ID, nil
}

// Update update user and return current user infomation
func (userService *UserService) Update(user dbe.User) error {
	if _, err := userService.db.Update(&user); err != nil {
		return err
	}
	return nil
}
