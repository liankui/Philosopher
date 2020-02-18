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

func (article *Article) Insert() int {
	result, e := initDB.Db.Exec("insert into `article` (type, content) values (?,?);",
		article.Type, article.Content)
	if e != nil {
		log.Panicln("文章添加失败", e.Error())
	}
	i, _ := result.LastInsertId()
	return int(i)
}

func (article Article) FindById() Article {
	row := initDB.Db.QueryRow("select * from `article` where id = ?;", article.Id)
	if e := row.Scan(&article.Id, &article.Type, &article.Content); e != nil {
		log.Panicln("FindById发生错误", e.Error())
	}
	return article
}
