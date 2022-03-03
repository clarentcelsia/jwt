package model

import (
	"restaurant/model"
)

type (
	User struct {
		UserID      string     `json:"userId"`
		Fullname    string     `json:"fullname" binding:"required"`
		Email       string     `json:"email" binding:"required"`
		PhoneNumber string     `json:"phoneNumber" binding:"required"`
		DOB         string     `json:"dob" binding:"required"`
		Address1    string     `json:"address1" binding:"required"`
		Address2    *string    `json:"address2"`
		Username    string     `json:"username" binding:"required"`
		Password    string     `json:"password" binding:"required"`
		Role        string     `json:"role"`
		Base        model.Base `json:"-"`
	}

	UserResponse struct {
		UserID      string  `json:"userId"`
		Fullname    string  `json:"fullname"`
		Email       string  `json:"email"`
		PhoneNumber string  `json:"phoneNumber"`
		DOB         string  `json:"dob"`
		Address1    string  `json:"address1"`
		Address2    *string `json:"address2"`
		Username    string  `json:"username"`
		Password    string  `json:"-"`
		Usrcrt      string  `json:"usrcrt"`
		Usrupd      string  `json:"usrupd"`
		Role        string  `json:"role"`
	}

	UserLoginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	UserLoginResponse struct {
		UserID   string
		Username string
		Email    string
		Password string
		Role     string
	}

	UserUpdatePassword struct {
		Username    string `json:"username"`
		NewPassword string `json:"newPassword"`
		Email       string `json:"email"`
	}

	UserUpdateResponse struct {
		Status   int    `json:"status"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Message  string `json:"newPassword"`
	}
)
