package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yueekee/Philosopher/GinHello/initRouter"
	"net/http"
	"net/http/httptest"
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

