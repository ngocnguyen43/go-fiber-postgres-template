package dtos

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
