package auth

import (
	"restaurant/config"
	mUser "restaurant/model/user"
	"restaurant/util"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var user mUser.UserResponse

func JWTMiddleware(alg string) (gin.HandlerFunc, *jwt.GinJWTMiddleware) {
	/** GIN-JWT
		#Signing Algorithm : [signature]. Algorithm used to sign token issued.
		#PayloadFunc : A set of fields that will be included in token being generated.
			usually used to include user identification detail.
			# note: payload won't encrypt the given data, make sure never include the password or such things.
			# note2: the output token can be decode to look what included data[payload]
		#IdentityHandler :

		#Authenticator: Give token for authenticator... [i.e admin and user]
			#note: loginhandler calls Authenticator.
		#Authorizator: Authorized for [i.e username = admin]..
			#To make this function works, make sure to define
				auth.Use(authMiddleware.MiddlewareFunc()){
	    			auth.GET("/hello", helloHandler)} -> admin can access this method
			)
	*/
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "RestaurantJWT",
		Key:              []byte(config.ProjectSecret()),
		Timeout:          time.Hour,
		MaxRefresh:       2 * time.Hour,
		IdentityKey:      "username",
		SigningAlgorithm: alg,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if model, ok := data.(mUser.UserResponse); ok {
				return jwt.MapClaims{
					"username": model.Username,
					"email":    model.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c) //extract jwt claims from mapclaims
			return &mUser.UserResponse{
				Username: claims["username"].(string),
				Email:    claims["email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login mUser.UserLoginReq
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userdb := CheckProfile(login.Username)
			isVerified := util.PasswordComparator(userdb.Password, login.Password)
			if login.Username == userdb.Username && isVerified {
				user.Username = userdb.Username
				user.Email = userdb.Email
				user.Role = userdb.Role
				return user, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			/* #Authorizator: Success[authorized] if..*/
			claims := jwt.ExtractClaims(c) //get attr value from jwt claims (payload)
			username := claims["username"].(string)
			response := CheckProfile(username)

			return strings.ToLower(response.Role) == "admin"
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			util.HandleError(c, gin.H{
				"code":    code,
				"message": message,
			}, "Unauthorized", "Unauthorized", code)
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			util.HandleSuccess(c, gin.H{
				"code":   code,
				"token":  message,
				"result": user,
			}, "Login Successfully", code)
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	// }, authMiddleware
	return authMiddleware.MiddlewareFunc(), authMiddleware
}

func CheckProfile(username string) mUser.UserLoginResponse {
	db := util.CheckConnection()

	var userLoginResp mUser.UserLoginResponse
	query := `SELECT user_id, username, password, email, role FROM users WHERE username = $1`
	rows := db.QueryRow(query, username)
	if rows != nil {
		if err := rows.Scan(&userLoginResp.UserID, &userLoginResp.Username, &userLoginResp.Password, &userLoginResp.Email, &userLoginResp.Role); err != nil {
			return mUser.UserLoginResponse{}
		}
		return mUser.UserLoginResponse{
			UserID:   userLoginResp.UserID,
			Username: userLoginResp.Username,
			Email:    userLoginResp.Email,
			Password: userLoginResp.Password, //hashed
			Role:     userLoginResp.Role,
		}
	}
	return mUser.UserLoginResponse{}
}
