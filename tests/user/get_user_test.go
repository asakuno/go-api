package user_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asakuno/go-api/controller"
	"github.com/asakuno/go-api/database/factory"
	"github.com/asakuno/go-api/dto/response"
	"github.com/asakuno/go-api/entity"
	"github.com/asakuno/go-api/repository"
	"github.com/asakuno/go-api/tests"
	user_usecase "github.com/asakuno/go-api/usecase/user"
	"github.com/asakuno/go-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ROUTE_NAME = "/api/v1/users"
)

func SetUpRoutes() *gin.Engine {
	route := gin.Default()
	return route
}

func SetUpUserController() controller.UserController {
	var (
		db             = tests.SetUpDatabaseConnection()
		userRepo       = repository.NewUserRepository(db)
		getUserUsecase = user_usecase.NewGetUserUsecase(userRepo)
		userController = controller.NewUserController(getUserUsecase)
	)

	return userController
}

func InsertTestUser() ([]entity.User, error) {
	db := tests.SetUpDatabaseConnection()
	userFactory := factory.NewUserFactory()

	users, err := userFactory.CreateAndSave(db)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func CleanUpTestData() {
	db := tests.SetUpDatabaseConnection()
	db.Exec("DELETE FROM users")
}

func Test_GetAllUser_OK(t *testing.T) {
	defer CleanUpTestData()

	route := SetUpRoutes()
	userControler := SetUpUserController()

	route.GET(ROUTE_NAME, userControler.GetAllUser)

	expectedUsers, err := InsertTestUser()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	request, _ := http.NewRequest(http.MethodGet, ROUTE_NAME, nil)
	writer := httptest.NewRecorder()

	route.ServeHTTP(writer, request)

	assert.Equal(t, http.StatusOK, writer.Code)

	var responseBody utils.Response
	err = json.Unmarshal(writer.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err.Error())
	}

	assert.True(t, responseBody.Status)

	responseData, err := json.Marshal(responseBody.Data)
	if err != nil {
		t.Fatalf("Failed to marshal response data: %v", err.Error())
	}

	var usersResponse response.GetAllUserResponse
	err = json.Unmarshal(responseData, &usersResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal users response: %v", err.Error())
	}

	assert.Equal(t, len(expectedUsers), usersResponse.Count)
	assert.Equal(t, len(expectedUsers), len(usersResponse.Users))

	expectedLoginIDs := make(map[string]bool)
	for _, user := range expectedUsers {
		expectedLoginIDs[user.LoginId] = true
	}

	for _, user := range usersResponse.Users {
		assert.True(t, expectedLoginIDs[user.LoginId], "User with LoginId %s not found in expected users", user.LoginId)
	}
}

func Test_GetAllUser_EmptyResult(t *testing.T) {
	defer CleanUpTestData()

	route := SetUpRoutes()
	userController := SetUpUserController()
	route.GET(ROUTE_NAME, userController.GetAllUser)

	request, _ := http.NewRequest(http.MethodGet, ROUTE_NAME, nil)
	writer := httptest.NewRecorder()
	route.ServeHTTP(writer, request)

	assert.Equal(t, http.StatusOK, writer.Code)

	var responseBody utils.Response
	err := json.Unmarshal(writer.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err.Error())
	}

	assert.True(t, responseBody.Status)

	responseData, err := json.Marshal(responseBody.Data)
	if err != nil {
		t.Fatalf("Failed to marshal response data: %v", err.Error())
	}

	var usersResponse response.GetAllUserResponse
	err = json.Unmarshal(responseData, &usersResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal users response: %v", err.Error())
	}

	assert.Equal(t, 0, usersResponse.Count)
	assert.Empty(t, usersResponse.Users)
}
