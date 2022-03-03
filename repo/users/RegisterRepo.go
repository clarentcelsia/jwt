package repo

import (
	mUser "restaurant/model/user"
	"restaurant/util"
	"strings"

	"github.com/google/uuid"
)

type IRegisterRepo interface {
	CreateUser(user mUser.User) (mUser.User, error)
	GetAllUser() ([]mUser.User, error)
}

//class RegisterRepo() *implement IRegisterUseCase{}
type RegisterRepo struct {
}

//override method [class RegisterRepo() implement IRegisterRepo{}]
func (r RegisterRepo) CreateUser(user mUser.User) (mUser.UserResponse, error) {
	db := util.CheckConnection()
	user.UserID = strings.Replace(uuid.Must(uuid.NewRandom()).String(), "-", "", -1)
	user.Base.PrePersist()
	encodedPass := util.PasswordEncoder(user.Password)
	// if char := user.Username[0:1]; char == "a" || char == "A" {
	// 	role := "Admin"
	// 	user.Role = role
	// } else {
	// 	role := "User"
	// 	user.Role = role
	// }

	var userResponse mUser.UserResponse
	query := `INSERT INTO users VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING *`
	if errrows := db.QueryRow(query,
		user.UserID,
		user.Fullname,
		user.Email,
		user.PhoneNumber,
		user.DOB,
		user.Address1,
		user.Address2,
		user.Username,
		encodedPass,
		user.Base.CreatedAt,
		user.Base.UpdatedAt,
		user.Role,
	).Scan(
		&userResponse.UserID,
		&userResponse.Fullname,
		&userResponse.Email,
		&userResponse.PhoneNumber,
		&userResponse.DOB,
		&userResponse.Address1,
		&userResponse.Address2,
		&userResponse.Username,
		&userResponse.Password,
		&userResponse.Usrcrt,
		&userResponse.Usrupd,
		&userResponse.Role,
	); errrows != nil {
		return mUser.UserResponse{}, errrows
	}
	return userResponse, nil
}

func (r RegisterRepo) GetAllUser() ([]mUser.UserResponse, error) {
	var (
		users []mUser.UserResponse
		user  mUser.UserResponse
	)

	db := util.CheckConnection()
	query := `SELECT * FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if errscan := rows.Scan(&user.UserID,
			&user.Fullname,
			&user.Email,
			&user.PhoneNumber,
			&user.DOB,
			&user.Address1,
			&user.Address2,
			&user.Username,
			&user.Password,
			&user.Usrcrt,
			&user.Usrupd,
			&user.Role,
		); errscan != nil {
			return nil, errscan
		}
		users = append(users, user)
	}
	defer rows.Close()
	return users, nil
}
