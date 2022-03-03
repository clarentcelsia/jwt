package controller

import (
	"net/http"
	"restaurant/config"
	mUser "restaurant/model/user"
	mUserRepo "restaurant/repo/users"
	"restaurant/util"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var repo mUserRepo.RegisterRepo

func RegisterUser(c *gin.Context) {
	var userRegister mUser.User
	if errbind := c.ShouldBindJSON(&userRegister); errbind != nil {
		util.HandleError(c, nil, "Can't bind data, please make sure the data has bound correctly", "bind error", http.StatusBadRequest)
		log.Error("Can't bind data, please make sure the data has bound correctly")
		log.Fatal(errbind)
	}

	repo = mUserRepo.RegisterRepo{}
	response, err := repo.CreateUser(userRegister)
	if err != nil {
		log.Error(util.INSERT_RES_FAILED)
		util.HandleError(c, nil, util.INSERT_RES_FAILED, "Insert failed", http.StatusInternalServerError)
	} else {
		util.HandleSuccess(c, response, util.INSERT_RES_SUCCEESS, http.StatusCreated)
	}
}

func LoginUser(c *gin.Context) {
	var login mUser.UserLoginReq
	if errbind := c.ShouldBindJSON(&login); errbind != nil {
		log.Error("Can't bind data, please make sure the data has bound correctly")
		log.Fatal(errbind)
	}

	url := config.Hostname() + config.ListenAndServeServerPort() + "/auth/loginhandler"

	println("url: " + url)
	loginMaps := structs.Map(login)
	if response, _, status, err := util.HttpPostReq(c, url, nil, nil, nil, loginMaps); err != nil {
		util.HandleError(c, response["Items"], "Login Failed", "Unauthorized", status)
	} else {
		util.HandleSuccess(c, response["Items"].(map[string]interface{}), "Login Successfully", status)
	}
}

func UpdatePassword(c *gin.Context) {
	var upd mUser.UserUpdatePassword
	if errbind := c.BindJSON(&upd); errbind != nil {
		util.HandleError(c, gin.H{
			"message": "Can't bind data, make sure data has bound correctly",
		}, "Can't bind data, make sure data has bound correctly", "BoundException", http.StatusBadRequest)
		panic(errbind)
	}
	claims := jwt.ExtractClaims(c)
	username := claims["username"].(string)
	//for other case
	//email := claims["email"].(string)
	println("username from claims : " + username)

	upd.Username = username
	repo := mUserRepo.UserRepo{}
	if user := repo.UpdateUserPassword(upd); (mUser.UserUpdateResponse{}) == user {
		log.Error(util.UPDATE_FAILED)
		util.HandleError(c, nil, util.UPDATE_FAILED, "Update failed", http.StatusInternalServerError)
	} else {
		util.HandleSuccess(c, user, util.UPDATE_SUCCESS, http.StatusOK)
	}
}

func GetAllUser(c *gin.Context) {
	repo = mUserRepo.RegisterRepo{}
	response, err := repo.GetAllUser()
	if err != nil {
		log.Error(util.GET_REQUEST_FAILED)
		util.HandleError(c, nil, util.GET_REQUEST_FAILED, "Fetch data failed", http.StatusInternalServerError)
	} else {
		util.HandleSuccess(c, response, util.GET_REQUEST_SUCCESS, http.StatusOK)
	}
}
