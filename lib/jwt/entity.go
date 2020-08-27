package jwt

//JWTValidateResponse ...
type JWTValidateResponse struct {
	ExpiresIn int64  `json:"expiresIn"`
	Role      string `json:"role"`
	UserID    int64  `json:"userID"`
	Email     string `json:"email"`
}
