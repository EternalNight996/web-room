// model/dbe/user.go
package dbe

import "time"

// User 用户实体，对应表user
type User struct {
	ID        int64
	Username  string
	Passwd    string
	Gender    int64 // 1 -> female, 2 -> male
	Nickname  string
	Mail      string
	CreatedAt time.Time `xorm:"created"` // 这个Field将在Insert时自动赋值为当前时间
	UpdatedAt time.Time `xorm:"updated"` // 这个Field将在Insert或Update时自动赋值为当前时间
	DeletedAt time.Time `xorm:"deleted"` // 如果带DeletedAt这个字段和标签，xorm删除时自动软删除
}
