package dto

type GetUserPublicResponse struct {
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

type GetUserDetailResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsSeller bool   `json:"is_seller"`
	WalletID int    `json:"wallet_id"`
}
