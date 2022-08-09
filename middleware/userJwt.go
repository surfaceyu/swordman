package middleware

import (
	"swordsman/logger"
	"swordsman/model"
	"swordsman/services"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	userIdentityKey = "id"
	AuthMiddleware  *jwt.GinJWTMiddleware
)

func init() {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",          //标识
		SigningAlgorithm: "HS256",              //加密算法
		Key:              []byte("secret key"), //密钥
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,       //刷新最大延长时间
		IdentityKey:      userIdentityKey, //指定cookie的id
		PayloadFunc: func(data interface{}) jwt.MapClaims { //负载，这里可以定义返回jwt中的payload数据
			if v, ok := data.(*model.Account); ok {
				return jwt.MapClaims{
					userIdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.Account{
				ID: claims[userIdentityKey].(string),
			}
		},
		Authenticator: authenticator, //在这里可以写我们的登录验证逻辑
		Authorizator: func(data interface{}, c *gin.Context) bool { //当用户通过token请求受限接口时，会经过这段逻辑
			if v, ok := data.(*model.Account); ok && v.ID != "" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) { //错误时响应
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// 指定从哪里获取token 其格式为："<source>:<name>" 如有多个，用逗号隔开
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		logger.Fatal("JWT Error:" + err.Error())
	}
	AuthMiddleware = authMiddleware
}

func authenticator(c *gin.Context) (interface{}, error) {
	var loginVal model.Account
	if err := c.ShouldBind(&loginVal); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	userID := loginVal.ID
	password := loginVal.Passwd

	var loginAccount = services.UserService.FindUser(userID)
	if loginAccount.ID == userID && loginAccount.Passwd == password {
		return &loginAccount, nil
	}
	return nil, jwt.ErrFailedAuthentication
}
