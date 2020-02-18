package article

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/model"
	"net/http"
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
