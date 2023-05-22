package dto

type GetUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsSeller bool   `json:"is_seller"`
}

type CreateUserResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
