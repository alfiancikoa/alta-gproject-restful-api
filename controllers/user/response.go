package user

type GetUserResponse struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phonenumber" form:"phonenumber"`
	Gender      string `json:"gender" form:"gender"`
	Birth       string `json:"birth" form:"birth"`
}

type PostUserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phonenumber" form:"phonenumber"`
	Gender      string `json:"gender" form:"gender"`
	Birth       string `json:"birth" form:"birth"`
}

type EditUserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phonenumber" form:"phonenumber"`
	Gender      string `json:"gender" form:"gender"`
	Birth       string `json:"birth" form:"birth"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}