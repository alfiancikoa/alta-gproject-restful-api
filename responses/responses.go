package responses

//Access Forbidden
func UnAuthorized() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Access Forbidden",
	}
}

//Invalid Email or Password
func InvalidEmailPassword() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Incorrect Email or Password",
	}
}

//Invalid password
func LoginFailed() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Fail to Login",
	}
}

//InternalServerErrorResponse default internal server error response
func InternalServerErrorResponse() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Server Internal Error",
	}
}

//Data Already Exist response
func DataAlreadyExist() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Data Already Exist",
	}
}

//BadRequestResponse default not found error response
func BadRequestResponse() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Bad Request",
	}
}

//InvalidFormatMethodInputdefault Invalid Format or Method
func InvalidFormatMethodInput() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Invalid Format Data or Invalid Request Method",
	}
}

//NewConflictResponse default not found error response
func DataNotExist() interface{} {
	return map[string]interface{}{
		"status": "not found", "message": "Data Not Found or Data Doesn't Exist",
	}
}
