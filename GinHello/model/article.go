package model

import (
	"github.com/yueekee/Philosopher/GinHello/initDB"
	"log"
)

type Article struct {
	Id		int		`json:"id"`
	Type 	string	`json:"type"`
	Content	string	`json:"content"`
}

func (article Article) Insert() int {
	result, e := initDB.Db.Exec("insert into `article` (type, content) values (?,?)",
		article.Type, article.Content)
	if e != nil {
		log.Panicln("文章添加失败", e.Error())
	}
	i, _ := result.LastInsertId()
	return int(i)
}
