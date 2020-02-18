package model

import (
	"database/sql"
	"github.com/yueekee/Philosopher/GinHello/initDB"
	"log"
)

type UserModel struct {
	Id				int			`form:"id"`
	Email			string		`form:"email" binding:"email"`
	Password		string		`form:"password"`
	Avatar			sql.NullString
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
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Println(e)
	}
	return u
}

func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where id = ?", id)
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Println(e)
	}
	return u, e
}

func (user *UserModel) Update(id int) error {
	stmt, e := initDB.Db.Prepare("update user set password=?,avatar=? where id=?")
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	_, err := stmt.Exec(user.Password, user.Avatar.String, user.Id)
	if err != nil {
		log.Panicln("exec发生了错误", err.Error())
	}
	return err
}

