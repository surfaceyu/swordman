package middleware

import (
	"swordsman/logger"
	"swordsman/msg"
	"swordsman/services"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	userIdentityKey     = "id"
	userIdentityAccount = "account"
	AuthMiddleware      *jwt.GinJWTMiddleware
)

func init() {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",          //标识
		SigningAlgorithm: "HS256",              //加密算法
		Key:              []byte("secret key"), //密钥
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,       //刷新最大延长时间
		IdentityKey:      userIdentityKey, //指定cookie的id
		PayloadFunc:      payloadFunc,
		IdentityHandler: func(c *gin.Context) interface{} {
			return JwtAccount(c)
		},
		Authenticator: authenticator, //在这里可以写我们的登录验证逻辑
		Authorizator:  authorizator,  //当用户通过token请求受限接口时，会经过这段逻辑
		Unauthorized:  unauthorized,
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

// 生成payload 数据
// 负载，这里可以定义返回jwt中的payload数据
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(msg.Account); ok {
		return jwt.MapClaims{
			userIdentityKey:     v.ID,
			userIdentityAccount: v.Account,
		}
	}
	return jwt.MapClaims{}
}

// 登录验证
func authenticator(c *gin.Context) (interface{}, error) {
	var loginVal msg.Account
	if err := c.ShouldBind(&loginVal); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	userAccount := loginVal.Account
	password := loginVal.Passwd

	var loginAccount = services.UserService.FindUser(userAccount)
	if loginAccount.ID != 0 && loginAccount.Passwd == password {
		return msg.Account{
			ID:      loginAccount.ID,
			Account: loginAccount.Account,
			Passwd:  password,
		}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// 请求接口权限验证
//当用户通过token请求受限接口时，会经过这段逻辑
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(msg.Account); ok && v.ID != 0 {
		return true
	}
	return false
}

// 错误时响应
func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func JwtAccount(c *gin.Context) msg.Account {
	claims := jwt.ExtractClaims(c)
	return msg.Account{
		ID:      uint(claims[userIdentityKey].(float64)),
		Account: claims[userIdentityAccount].(string),
	}
}
