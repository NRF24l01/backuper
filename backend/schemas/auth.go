package schemas

type AuthRequest struct {
	Username string `json:"nick" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
