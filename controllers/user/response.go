package user

type GetUserResponse struct {
	ID          int
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
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
	Role        string `json:"role" form:"role"`
}
type PostUserRequestErr struct {
	Name int
}

type EditUserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phonenumber" form:"phonenumber"`
	Gender      string `json:"gender" form:"gender"`
	Birth       string `json:"birth" form:"birth"`
}
type EditUserRequestErr struct {
	Name int
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type LoginUserRequestErr struct {
	Email    int
	Password int
}
