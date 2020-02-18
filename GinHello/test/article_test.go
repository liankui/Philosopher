package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yueekee/Philosopher/GinHello/initRouter"
	"github.com/yueekee/Philosopher/GinHello/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func init() {
	router = initRouter.SetupRouter()
}

func TestInsertArticle(t *testing.T) {
	article := model.Article{
		Type: 		"go",
		Content:	"hello gin",
	}
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/article",
		bytes.NewBufferString(string(marshal)))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, "{id:-1}", w.Body.String())
}
