package users_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asakuno/go-api/dto/request"
	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Register_OK(t *testing.T) {
	defer CleanUpTestData()

	route := SetUpRoutes()
	userController := SetUpUserController()

	route.POST(ROUTE_NAME, userController.Register)

	// Create test request
	createUserReq := request.CreateUserRequest{
		LoginId:  "testuser123",
		Email:    "testuser123@example.com",
		Password: "password123",
	}

	reqBody, _ := json.Marshal(createUserReq)
	req, _ := http.NewRequest(http.MethodPost, ROUTE_NAME, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response utils.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.True(t, response.Status)
	assert.Equal(t, "ユーザーが正常に登録されました", response.Message)
}

func Test_Register_DuplicateLoginID(t *testing.T) {
	defer CleanUpTestData()

	route := SetUpRoutes()
	userController := SetUpUserController()

	route.POST(ROUTE_NAME, userController.Register)

	// Insert an initial user
	users, err := InsertTestUser()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	
	firstUser := users[0]

	// Try to register with the same login ID
	createUserReq := request.CreateUserRequest{
		LoginId:  firstUser.LoginId,
		Email:    "different@example.com",
		Password: "password123",
	}

	reqBody, _ := json.Marshal(createUserReq)
	req, _ := http.NewRequest(http.MethodPost, ROUTE_NAME, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response utils.Response
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.False(t, response.Status)
	assert.Equal(t, "入力されたログインIDは既に使用されています", response.Message)
}

func Test_Register_DuplicateEmail(t *testing.T) {
	defer CleanUpTestData()

	route := SetUpRoutes()
	userController := SetUpUserController()

	route.POST(ROUTE_NAME, userController.Register)

	// Insert an initial user
	users, err := InsertTestUser()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	
	firstUser := users[0]

	// Try to register with the same email
	createUserReq := request.CreateUserRequest{
		LoginId:  "differentloginid",
		Email:    firstUser.Email,
		Password: "password123",
	}

	reqBody, _ := json.Marshal(createUserReq)
	req, _ := http.NewRequest(http.MethodPost, ROUTE_NAME, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response utils.Response
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.False(t, response.Status)
	assert.Equal(t, "入力されたメールアドレスは既に使用されています", response.Message)
}

func Test_Register_InvalidData(t *testing.T) {
	defer CleanUpTestData()

	route := SetUpRoutes()
	userController := SetUpUserController()

	route.POST(ROUTE_NAME, userController.Register)

	// Create test request with invalid data
	createUserReq := request.CreateUserRequest{
		LoginId:  "t", // Too short
		Email:    "not-an-email",
		Password: "short", // Too short
	}

	reqBody, _ := json.Marshal(createUserReq)
	req, _ := http.NewRequest(http.MethodPost, ROUTE_NAME, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response utils.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.False(t, response.Status)
	assert.Equal(t, "入力内容に問題があります", response.Message)
}
