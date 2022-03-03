package model

type (
	TokenClaims struct {
		UserID       string
		Username     string
		Email        string
		IdentityUser string
		Alg          string
		Iat          float64 //issuedAt
		Exp          float64 //expires
	}

	ForgotPassword struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	ResetPassword struct {
		Username    string `json:"username" binding:"required"`
		OldPassword string
		NewPassword string
	}
)
