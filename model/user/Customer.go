package model

import (
	mBase "restaurant/model"
)

type (
	Customer struct {
		CustomerID    string    `json:"customer_id"`
		CustomerName  string    `json:"customer_name"`
		CustomerEmail string    `json:"customer_email"`
		CustomerPhone string    `json:"customer_phone"`
		CustomerDOB   string    `json:"customer_dob"`
		CustomerAddress string `json:"customer_address"`
		Base mBase.Base `json:"base"`
		IsDeleted bool `json:"is_deleted"`
	}
)
