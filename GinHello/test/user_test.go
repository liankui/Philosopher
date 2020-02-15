package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yueekee/Philosopher/GinHello/initRouter"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "eric"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+"已经保存", w.Body.String())
}

func TestUserSaveByQuery(t *testing.T) {
	username := "eric"
	age := 27
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+"年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}
