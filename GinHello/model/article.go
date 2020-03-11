package model

import (
	"github.com/yueekee/Philosopher/GinHello/initDB"
)

type Article struct {
	Id		int		`json:"id"`
	Type 	string	`json:"type"`
	Content	string	`json:"content"`
}

func (article *Article) Insert() int {
	//result, e := initDB.Db.Exec("insert into `article` (type, content) values (?,?);",
	//	article.Type, article.Content)
	//if e != nil {
	//	log.Panicln("文章添加失败", e.Error())
	//}
	//i, _ := result.LastInsertId()
	//return int(i)
	create := initDB.Db.CreateTable(&article)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (article Article) FindById() Article {
	//row := initDB.Db.QueryRow("select * from `article` where id = ?;", article.Id)
	//if e := row.Scan(&article.Id, &article.Type, &article.Content); e != nil {
	//	log.Panicln("FindById发生错误", e.Error())
	//}
	//return article
	initDB.Db.First(&article, article.Id)
	return article
}

func (article *Article) FindAll() []Article {
	//rows, e := initDB.Db.Query("select * from `article`;")
	//if e != nil {
	//	log.Panicln("FindAll发生错误", e.Error())
	//}
	//
	//var articles []Article
	//for rows.Next() {
	//	var a Article
	//	if e := rows.Scan(&a.Id, &a.Type, &a.Content); e == nil {
	//		articles = append(articles, a)
	//	}
	//}
	//
	//return articles
	var articles []Article
	initDB.Db.Find(&articles)
	return articles
}

func (article Article) DeleteOne() Article {
	//_, e := initDB.Db.Exec("delete from `article` where id = ?;", article.Id)
	//if e != nil {
	//	log.Panicln("删除数据发生错误", e.Error())
	//}
	initDB.Db.Delete(&article)
	return article
}