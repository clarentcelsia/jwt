package controller

import (
	"net/http"
	"reflect"
	middleware "restaurant/authorization"
	"restaurant/config"
	"restaurant/util"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	//get token
	strAuthorization := "Authorization"
	strSpace := " "

	paramHeader := c.Request.Header[strAuthorization][0]
	tokenString := strings.Split(paramHeader, strSpace)[1]

	_, authMiddleware := middleware.JWTMiddleware("HS256")
	jwtToken, errParse := authMiddleware.ParseTokenString(tokenString)
	claims := jwt.ExtractClaimsFromToken(jwtToken)
	//if interface empty
	if reflect.TypeOf(claims["username"]) == reflect.TypeOf(nil) {
		util.HandleError(c, gin.H{
			"error message": errParse,
		}, "Invalid token", "Invalid token", http.StatusUnauthorized)
		return
	}

	// username := claims["username"].(string)
	// email := claims["email"].(string)

	strBearerToken := "Bearer " + tokenString
	urlRefreshToken := config.Hostname() + config.ListenAndServeServerPort() + "/api/v1/private/restaurant/refreshtoken"
	queryParamsRefresh := map[string]string{}
	headersRefresh := map[string]string{
		strAuthorization: strBearerToken}

	resp, _, respCodeRefresh, errHttpReqRefresh := util.HttpGetReq(c, urlRefreshToken, headersRefresh, queryParamsRefresh, map[string]string{}, map[string]interface{}{})
	if errHttpReqRefresh != nil {
		c.JSON(respCodeRefresh, gin.H{
			"result": errHttpReqRefresh,
		})
	} else {
		c.JSON(respCodeRefresh, gin.H{
			"result": resp,
		})
	}
}
