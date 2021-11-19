package product

import (
	"alte/e-commerce/config"
	"alte/e-commerce/constants"
	"alte/e-commerce/controllers/user"
	"alte/e-commerce/lib/database"
	"alte/e-commerce/middlewares"
	"alte/e-commerce/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type ProductResponSuccess struct {
	Status  string
	Message string
	Data    []models.Product
}

type ResponseFailed struct {
	Status  string
	Message string
}

//without slice
type SingleProductResponseSuccess struct {
	Status  string
	Message string
	Data    models.Product
}

var logininfo = user.LoginUserRequest{
	Email:    "fian@gmail.com",
	Password: "admin",
}
var (
	mock_data_product = models.Product{
		Title:       "Jaket Hoodie ERIGO",
		Desc:        "size M",
		Price:       50000,
		Status:      "ready",
		Category_ID: 1,
		User_ID:     1,
	}
)

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

func InsertMockDataProductToDB() error {
	var err error
	if err = config.DB.Save(&mock_data_product).Error; err != nil {
		return err
	}
	return nil
}

var xpass string

func InsertMockDataUserToDB() error {
	xpass, _ = database.EncryptPassword(mock_data_user.Password)
	mock_data_user.Password = xpass
	var err error
	if err = config.DB.Save(&mock_data_user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetAllProductSuccess(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataProductToDB()
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/products")
	if assert.NoError(t, GetAllProductsController(context)) {
		body := rec.Body.String()
		var product ProductResponSuccess
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get all product", product.Message)
		assert.Equal(t, "success", product.Status)
		assert.Equal(t, 1, len(product.Data))
		assert.Equal(t, "Jaket Hoodie ERIGO", product.Data[0].Title)
	}
}

func TestGetAllProductFailed(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataProductToDB()
	config.DB.Migrator().DropTable(&models.Product{})
	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/products")
	if assert.NoError(t, GetAllProductsController(context)) {
		body := rec.Body.String()
		var product ProductResponSuccess
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Server Internal Error", product.Message)
		assert.Equal(t, "failed", product.Status)
	}
}

func TestGetProductByIdSuccess(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataProductToDB()
	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	if assert.NoError(t, GetProductController(context)) {
		body := rec.Body.String()
		var product SingleProductResponseSuccess
		err := json.Unmarshal([]byte(body), &product)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get product", product.Message)
		assert.Equal(t, "success", product.Status)
		assert.Equal(t, "Jaket Hoodie ERIGO", product.Data.Title)
	}
}

func TestGetProductByIdFailed(t *testing.T) {
	// create database connection and create controller
	e := InitEchoTestAPI()
	InsertMockDataProductToDB()
	// setting controller
	t.Run("TestGETUserDetail_InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		if assert.NoError(t, GetProductController(context)) {
			body := rec.Body.String()
			var respon ResponseFailed
			err := json.Unmarshal([]byte(body), &respon)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
			assert.Equal(t, "failed", respon.Status)
		}
	})
	t.Run("TestGETUserDetail_NotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("100")
		if assert.NoError(t, GetProductController(context)) {
			body := rec.Body.String()
			var respon ResponseFailed
			err := json.Unmarshal([]byte(body), &respon)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "Data Not Found or Data Doesn't Exist", respon.Message)
			assert.Equal(t, "not found", respon.Status)
		}
	})
	t.Run("TestGETUserDetail_ErrorDB", func(t *testing.T) {
		config.DB.Migrator().DropTable(&models.Product{})
		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("100")
		if assert.NoError(t, GetProductController(context)) {
			body := rec.Body.String()
			var respon ResponseFailed
			err := json.Unmarshal([]byte(body), &respon)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "Server Internal Error", respon.Message)
			assert.Equal(t, "failed", respon.Status)
		}
	})
}

func TestCreateProductSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	body, err := json.Marshal(mock_data_product)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products")
	middleware.JWT([]byte(constants.SECRET_JWT))(CreateProductsControllerTesting())(context)
	var product SingleProductResponseSuccess
	bodyRes := res.Body.String()
	json.Unmarshal([]byte(bodyRes), &product)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success create new product", product.Message)
	assert.Equal(t, "success", product.Status)
}

func TestCreateProductFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	t.Run("TestPOST_InputEmpty", func(t *testing.T) {
		newbody, err := json.Marshal(EditProduct{Title: "", Desc: "", Price: 0})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		middleware.JWT([]byte(constants.SECRET_JWT))(CreateProductsControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestPOST_ErrorBind", func(t *testing.T) {
		newbody, err := json.Marshal(EditProductErr{})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		middleware.JWT([]byte(constants.SECRET_JWT))(CreateProductsControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Bad Request", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestPOST_ErrorDB", func(t *testing.T) {
		newbody, err := json.Marshal(mock_data_product)
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		config.DB.Migrator().DropTable(&models.Product{})
		middleware.JWT([]byte(constants.SECRET_JWT))(CreateProductsControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Server Internal Error", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})

}

func TestUpdateProductSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataProductToDB()
	var newdata = EditProduct{
		Title:       "Jaket Hoodie Update",
		Desc:        "size L",
		Price:       1000,
		Status:      "ready",
		Category_ID: 1,
	}
	newbody, err := json.Marshal(newdata)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/products/:id", bytes.NewBuffer(newbody))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllerTesting())(context)
	var product SingleProductResponseSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &product)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "product update successful", product.Message)
	assert.Equal(t, "success", product.Status)
}

func TestUpdateProductFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataProductToDB()
	var newdata = EditProduct{
		Title:       "Jaket Hoodie Update",
		Desc:        "size L",
		Price:       1000,
		Status:      "ready",
		Category_ID: 1,
	}
	newbody, err := json.Marshal(newdata)
	if err != nil {
		t.Error(t, err, "error marshal")
	}
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	t.Run("TestEditUserDetail_InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/products/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "Data Not Found or Data Doesn't Exist", respon.Message)
		assert.Equal(t, "not found", respon.Status)
	})
	t.Run("TestEditUserDetail_InvalidIMethod", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/products/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestEditUserDetail_EmptyInput", func(t *testing.T) {
		newbody, err := json.Marshal(EditProduct{Title: "", Desc: "", Price: 0})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPut, "/products/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestEditUserDetail_ErrorBind", func(t *testing.T) {
		newbody, err := json.Marshal(EditProductErr{})
		if err != nil {
			t.Error(t, err, "error marshal")
		}
		req := httptest.NewRequest(http.MethodPut, "/products/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Bad Request", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestEditUserDetail_ErrorDB", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/products/:id", bytes.NewBuffer(newbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// Drop table user ini DB test to create failed condition
		config.DB.Migrator().DropTable(&models.Product{})
		middleware.JWT([]byte(constants.SECRET_JWT))(UpdateProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Server Internal Error", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
}

func TestDeleteProductSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataProductToDB()
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductControllerTesting())(context)
	var product SingleProductResponseSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &product)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success", product.Status)
	assert.Equal(t, "product deleted successfully", product.Message)
}

func TestDeleteProductFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	t.Run("TestDeleteProductDetail_NotFound", func(t *testing.T) {
		InsertMockDataProductToDB()
		var userDetail models.User
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("100")
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "Data Not Found or Data Doesn't Exist", respon.Message)
		assert.Equal(t, "not found", respon.Status)
	})
	t.Run("TestDeleteProductDetail_InvalidMethod", func(t *testing.T) {
		InsertMockDataProductToDB()
		var userDetail models.User
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("#")
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Invalid Format Data or Invalid Request Method", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
	t.Run("TestDeleteProductDetail_ErrorDB", func(t *testing.T) {
		InsertMockDataUserToDB()
		InsertMockDataProductToDB()
		var userDetail models.User
		tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
		if tx.Error != nil {
			t.Error(tx.Error)
		}
		token, err := middlewares.CreateToken(int(userDetail.ID))
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		config.DB.Migrator().DropTable(&models.Product{})
		middleware.JWT([]byte(constants.SECRET_JWT))(DeleteProductControllerTesting())(context)
		var respon ResponseFailed
		body := res.Body.String()
		json.Unmarshal([]byte(body), &respon)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Server Internal Error", respon.Message)
		assert.Equal(t, "failed", respon.Status)
	})
}

func TestGetMyProductSuccess(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataProductToDB()
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/products/my", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products/my")
	middleware.JWT([]byte(constants.SECRET_JWT))(GetMyProductControllerTesting())(context)
	var product ProductResponSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &product)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "success get all product", product.Message)
	assert.Equal(t, "success", product.Status)
	assert.Equal(t, 1, len(product.Data))
	assert.Equal(t, "Jaket Hoodie ERIGO", product.Data[0].Title)
}

func TestGetMyProductFailed(t *testing.T) {
	e := InitEchoTestAPI()
	InsertMockDataUserToDB()
	InsertMockDataProductToDB()
	var userDetail models.User
	tx := config.DB.Where("email = ? AND password = ?", logininfo.Email, xpass).First(&userDetail)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
	token, err := middlewares.CreateToken(int(userDetail.ID))
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/products/my", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products/my")
	config.DB.Migrator().DropTable(&models.Product{})
	middleware.JWT([]byte(constants.SECRET_JWT))(GetMyProductControllerTesting())(context)
	var product ProductResponSuccess
	body := res.Body.String()
	json.Unmarshal([]byte(body), &product)
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, "Server Internal Error", product.Message)
	assert.Equal(t, "failed", product.Status)
}
