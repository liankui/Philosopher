package article

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/model"
	"log"
	"net/http"
	"strconv"
)

func Insert(context *gin.Context) {
	article := model.Article{}
	var id = -1
	if err := context.ShouldBindJSON(&article); err == nil {
		id = article.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id不是int类型，id转化失败", e.Error())
	}
	article := model.Article{
		Id: i,
	}
	artic := article.FindById()
	context.JSON(http.StatusOK, gin.H{
		"article": artic,
	})
}
