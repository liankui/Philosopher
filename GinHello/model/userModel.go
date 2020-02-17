package model

import (
	"github.com/yueekee/Philosopher/GinHello/initDB"
	"log"
)

type UserModel struct {
	Id				int			`form:"id"`
	Email			string		`form:"email" binding:"email"`
	Password		string		`form:"password"`
	PasswordAgain	string		`form:"password-again" binding:"eqfield=Password"`
}

// 增加用户
func (user *UserModel) Save() int64 {
	result, e := initDB.Db.Exec("insert into ginhello.user (email, password) values (?,?);",
		user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert err:", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id err:", err.Error())
	}
	return id
}

// 验证密码，登录
func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where email = ?", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password)
	if e != nil {
		log.Println(e)
	}
	return u
}
