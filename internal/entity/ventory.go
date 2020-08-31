package entity

type Response struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
		ID    int64  `json:"id"`
	}

	RegisterRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   int64  `json:"roleId"`
	}

	VerifySession struct {
		UserID int64  `json:"id"`
		Role   string `json:"role"`
	}
)

type (
	GetItemResponse struct {
		Items          []Item `json:"items"`
		TotalPage      int64  `json:"totalPage"`
		MaxDataPerPage int64  `json:"maxDataPerPage"`
	}
)

type (
	GetItemCategoryResponse struct {
		Category       []ItemCategory `json:"category"`
		TotalPage      int64          `json:"totalPage"`
		MaxDataPerPage int64          `json:"maxDataPerPage"`
	}
)
