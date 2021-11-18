package user

import (
	"alte/e-commerce/config"
	"alte/e-commerce/constants"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type UsersResponSuccess struct {
	Status  string
	Message string
	Data    []models.User
}

type ResponseFailed struct {
	Status  string
	Message string
}

//without slice
type SingleUserResponseSuccess struct {
	Status  string
	Message string
	Data    models.User
}

var logininfo = LoginUserRequest{
	Email:    "fian@gmail.com",
	Password: "admin",
}
var (
	mock_data_user = models.User{
		Name:        "fian",
		Email:       "fian@gmail.com",
		Password:    "admin",
		PhoneNumber: "081xxx",
		Gender:      "male",
		Birth:       "2021-09-21",
		Role:        "user",
	}
)

func InsertMockDataUserToDB() error {
	var err error
	if err = config.DB.Save(&mock_data_user).Error; err != nil {
		return err
	}
	return nil
}

func TestLoginJWTSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	logininfo, err := json.Marshal(LoginUserRequest{Email: "fian@gmail.com", Password: "admin"})
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	// send data using request body with HTTP method POST
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(logininfo))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	contex := e.NewContext(req, rec)
	contex.SetPath("/login")

	if assert.NoError(t, LoginUsersController(contex)) {
		bodyResponses := rec.Body.String()
		var user SingleUserResponseSuccess
		err := json.Unmarshal([]byte(bodyResponses), &user)
		if err != nil {
			assert.Error(t, err, "error marshal")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success", user.Status)
		assert.Equal(t, "fian@gmail.com", user.Data.Email)
		assert.Equal(t, "admin", user.Data.Password)
	}
}

func TestLoginJWTFailed(t *testing.T) {
	var logininfo = []LoginUserRequest{
		{
			// Email: Benar; Password: Salah
			Email:    "fian@gmail.com",
			Password: "salah",
		}, {
			// Email: Salah; Password: Benar
			Email:    "salah@gmail.com",
			Password: "admin",
		}, {
			// Email: Salah; Password: Salah
			Email:    "salah@gmail.com",
			Password: "salah",
		},
	}
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	for _, loginlist := range logininfo {
		logininfo, err := json.Marshal(loginlist)
		log.Println("logininfo", string(logininfo))
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(logininfo))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		contex := e.NewContext(req, rec)
		contex.SetPath("/login")

		if assert.NoError(t, LoginUsersController(contex)) {
			body := rec.Body.String()
			var respon ResponseFailed
			err := json.Unmarshal([]byte(body), &respon)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			require.Equal(t, http.StatusBadRequest, rec.Code)
			require.Equal(t, "failed", respon.Status)
			require.Equal(t, "login failed", respon.Message)
		}
	}
}

func TestLoginJWTFailedBind(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	logininfo, err := json.Marshal(LoginUserRequestErr{})
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	// send data using request body with HTTP method POST
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(logininfo))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	contex := e.NewContext(req, rec)
	contex.SetPath("/login")

	if assert.NoError(t, LoginUsersController(contex)) {
		bodyResponses := rec.Body.String()
		var user SingleUserResponseSuccess
		err := json.Unmarshal([]byte(bodyResponses), &user)
		if err != nil {
			assert.Error(t, err, "error marshal")
		}
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "failed", user.Status)
		assert.Equal(t, "Bad Request", user.Message)
	}
}

func TestGetUserByIdSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB() //create token
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetUserByIdControllerTesting())(context)

	var user SingleUserResponseSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &user)
	t.Run("GET/jwt/users/:id", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success get user detail", user.Message)
		assert.Equal(t, "success", user.Status)
		assert.Equal(t, "fian", user.Data.Name)
	})
}

func TestGetGetUserByIdFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	//create token
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	t.Run("TestGETUserDetail_InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")
		middleware.JWT([]byte(constants.SECRET_JWT))(GetUserByIdControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, "Access Forbidden", respon.Message)
		assert.Equal(t, "failed", respon.Status)

	})
	t.Run("TestGETUserDetail_InvalidMethod", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("@")
		middleware.JWT([]byte(constants.SECRET_JWT))(GetUserByIdControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)

	})
	t.Run("TestGETUserDetail_ErrorDB", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// Drop table user ini DB test to create failed condition
		config.DB.Migrator().DropTable(&models.User{})
		middleware.JWT([]byte(constants.SECRET_JWT))(GetUserByIdControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Server Internal Error", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
}

func TestCreateUserSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	body, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	// send data using request body with HTTP method POST
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	contex := e.NewContext(req, rec)
	if assert.NoError(t, CreateUserController(contex)) {
		body := rec.Body.String()
		var user SingleUserResponseSuccess
		err := json.Unmarshal([]byte(body), &user)
		if err != nil {
			assert.Error(t, err, "error marshal")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "fian", user.Data.Name)
		assert.Equal(t, "fian@gmail.com", user.Data.Email)
		assert.Equal(t, "admin", user.Data.Password)
	}
}

func TestCreateUserFailed(t *testing.T) {
	e := InitEchoTestAPI()
	body, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	t.Run("TestPOST_InputEmpty", func(t *testing.T) {
		// Try to Give Input is Empty
		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		rec := httptest.NewRecorder()
		contex := e.NewContext(req, rec)
		// Call function on controller
		if assert.NoError(t, CreateUserController(contex)) {
			body := rec.Body.String()
			var respon ResponseFailed
			err := json.Unmarshal([]byte(body), &respon)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
			assert.Equal(t, "failed", respon.Status)
		}
	})
	t.Run("TestPOST_ErrorDB", func(t *testing.T) {
		// Drop table user ini DB test to create failed condition
		config.DB.Migrator().DropTable(&models.User{})
		// send data using request body with HTTP method POST
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		contex := e.NewContext(req, rec)
		// Call function on controller
		if assert.NoError(t, CreateUserController(contex)) {
			body := rec.Body.String()
			var respon ResponseFailed
			err := json.Unmarshal([]byte(body), &respon)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "Server Internal Error", respon.Message)
			assert.Equal(t, "failed", respon.Status)
		}
	})
	t.Run("TestPOST_ErrorBind", func(t *testing.T) {
		body, err := json.Marshal(PostUserRequestErr{})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		rec := httptest.NewRecorder()
		contex := e.NewContext(req, rec)
		// Call function on controller
		if assert.NoError(t, CreateUserController(contex)) {
			body := rec.Body.String()
			var respon ResponseFailed
			err := json.Unmarshal([]byte(body), &respon)
			if err != nil {
				assert.Error(t, err, "error marshal")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Bad Request", respon.Message)
			assert.Equal(t, "failed", respon.Status)
		}
	})
}

func TestUpdateUserSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	var newdata = EditUserRequest{
		Name:        "fianUpdate",
		Email:       "fian@gmail.com",
		Password:    "admin",
		PhoneNumber: "081222333",
		Gender:      "male",
		Birth:       "2021-09-21",
	}
	newbody, err := json.Marshal(newdata)
	if err != nil {
		t.Error(t, err, "error marshal")
	}

	req := httptest.NewRequest(http.MethodPut, "/users/:id", bytes.NewBuffer(newbody))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllerTesting())(context)

	var user SingleUserResponseSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &user)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "fianUpdate", user.Data.Name)
	assert.Equal(t, "fian@gmail.com", user.Data.Email)
	assert.Equal(t, "admin", user.Data.Password)
}

func TestUpdateUserFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	var newdata = EditUserRequest{
		Name:        "fianUpdate",
		Email:       "fian@gmail.com",
		Password:    "admin",
		PhoneNumber: "081222333",
		Gender:      "male",
		Birth:       "2021-09-21",
	}
	newbody, err := json.Marshal(newdata)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	t.Run("TestEditUserDetail_InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, "Access Forbidden", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestEditUserDetail_InvalidMethod", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestEditUserDetail_EmptyInput", func(t *testing.T) {
		newbody, err := json.Marshal(EditUserRequest{Name: "", Email: "", Password: ""})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPut, "/users:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// Drop table user ini DB test to create failed condition
		config.DB.Migrator().DropTable(&models.User{})
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestEditUserDetail_ErrorBind", func(t *testing.T) {
		newbody, err := json.Marshal(EditUserRequestErr{})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPut, "/users:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// Drop table user ini DB test to create failed condition
		config.DB.Migrator().DropTable(&models.User{})
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Bad Request", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestEditUserDetail_ErrorDB", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// Drop table user ini DB test to create failed condition
		config.DB.Migrator().DropTable(&models.User{})
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Server Internal Error", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
}

func TestDeleteUserSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodDelete, "/users:id", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllerTesting())(context)
	var user SingleUserResponseSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &user)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success", user.Status)
	assert.Equal(t, "success deleted user", user.Message)
}

func TestDeleteUserFailed(t *testing.T) {
	e := InitEchoTestAPI()

	t.Run("TestDEL_InvalidMethod", func(t *testing.T) {
		InsertMockDataUserToDB()
		var userDetail models.User
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/users:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestDEL_InvalidID", func(t *testing.T) {
		InsertMockDataUserToDB()
		var userDetail models.User
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/users:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("3")
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, "Access Forbidden", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestDEL_ErrorDB", func(t *testing.T) {
		InsertMockDataUserToDB()
		var userDetail models.User
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, logininfo.Password).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/users:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// Drop table user ini DB test to create failed condition
		config.DB.Migrator().DropTable(&models.User{})
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteUserControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Server Internal Error", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
}
