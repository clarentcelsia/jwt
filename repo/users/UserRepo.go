package repo

import (
	"net/http"
	mUser "restaurant/model/user"
	"restaurant/util"
)

type IUserRepo interface {
	CheckUserByUsername(username string)
	UpdateUserPassword(user mUser.UserUpdatePassword)
}

type UserRepo struct{}

func (u UserRepo) CheckUserByUsername(username string) mUser.UserResponse {
	// db := util.CheckConnection()

	// var user mUser.UserResponse
	// query := `SELECT * FROM users WHERE username = $1`
	// rows := db.QueryRow(query, username)
	// if rows.Err() != nil {
	// rows.Scan(&)
	// }
	return mUser.UserResponse{}
}

func (u UserRepo) UpdateUserPassword(user mUser.UserUpdatePassword) mUser.UserUpdateResponse {
	var userUpdatePass mUser.UserUpdateResponse
	db := util.CheckConnection()

	query := `UPDATE users SET password = $1 WHERE username = $2 RETURNING username, email`
	newEncodedPass := util.PasswordEncoder(user.NewPassword)
	if errrows := db.QueryRow(query, newEncodedPass, user.Username).Scan(&userUpdatePass.Username, &userUpdatePass.Email); errrows != nil {
		return mUser.UserUpdateResponse{
			Status:   http.StatusInternalServerError,
			Username: "",
			Email:    "",
			Message:  "Can't update data",
		}
	}
	userUpdatePass.Status = http.StatusOK
	userUpdatePass.Message = "Update successfully"
	return userUpdatePass
}
