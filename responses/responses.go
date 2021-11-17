package responses

//Access Forbidden
func UnAuthorized() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Access Forbiddenr",
	}
}

//InternalServerErrorResponse default internal server error response
func InternalServerErrorResponse() interface{} {
	return map[string]interface{}{
		"status": "failed", "message": "Server Internal Error",
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
