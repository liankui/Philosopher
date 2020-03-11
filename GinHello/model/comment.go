package model

import (
	"github.com/jinzhu/gorm"
	"github.com/yueekee/Philosopher/GinHello/initDB"
)

type Comment struct {
	gorm.Model
	Content 	string
}

func init() {
	// 检查之前的表是否存在
	hasTable := initDB.Db.HasTable(Comment{})
	if !hasTable {
		initDB.Db.Create(Comment{})
	}
}
