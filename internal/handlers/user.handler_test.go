package handlers

import (
	"biFebriansyah/gogin/config"
	"biFebriansyah/gogin/internal/repositories"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoUserMock = repositories.RepoUserMock{}
var reqBody = `{
	"user_id": "123",
	"username": "testing",
	"password": "abcd1234",
	"role": "user"
}`

func TestPostData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := &config.Result{Data: "User created"}
	repoUserMock.On("CreateUser", mock.Anything).Return(expectedResult, nil)

	r.POST("/create", handler.PostData)
	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"data": "User created", "status":"OK"}`, w.Body.String())

}
