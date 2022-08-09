package middleware

import (
	"time"
	"wordGame/logger"
	"wordGame/model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	adminIdentityKey = "admin"
	AdminMiddleware  *jwt.GinJWTMiddleware
)

func init() {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",          //标识
		SigningAlgorithm: "HS256",              //加密算法
		Key:              []byte("secret key"), //密钥
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,        //刷新最大延长时间
		IdentityKey:      adminIdentityKey, //指定cookie的id
		PayloadFunc: func(data interface{}) jwt.MapClaims { //负载，这里可以定义返回jwt中的payload数据
			if v, ok := data.(*model.Account); ok {
				return jwt.MapClaims{
					adminIdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.Account{
				ID: claims[adminIdentityKey].(string),
			}
		},
		Authenticator: authenticator, //在这里可以写我们的登录验证逻辑
		Authorizator: func(data interface{}, c *gin.Context) bool { //当用户通过token请求受限接口时，会经过这段逻辑
			if v, ok := data.(*model.Account); ok && v.ID == "admin" {
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
	adminMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",          //标识
		SigningAlgorithm: "HS256",              //加密算法
		Key:              []byte("secret key"), //密钥
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,        //刷新最大延长时间
		IdentityKey:      adminIdentityKey, //指定cookie的id
		PayloadFunc: func(data interface{}) jwt.MapClaims { //负载，这里可以定义返回jwt中的payload数据
			if v, ok := data.(*model.Account); ok {
				return jwt.MapClaims{
					adminIdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.Account{
				ID: claims[adminIdentityKey].(string),
			}
		},
		Authenticator: adminAuthenticator, //在这里可以写我们的登录验证逻辑
		Authorizator: func(data interface{}, c *gin.Context) bool { //当用户通过token请求受限接口时，会经过这段逻辑
			if v, ok := data.(*model.Account); ok && v.ID == "admin" {
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
	AdminMiddleware = adminMiddleware
}

func adminAuthenticator(c *gin.Context) (interface{}, error) {
	var loginVal model.Account
	if err := c.ShouldBind(&loginVal); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVal.ID
	password := loginVal.Passwd

	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		return &model.Account{
			ID:     "admin",
			Passwd: "admin",
		}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}
