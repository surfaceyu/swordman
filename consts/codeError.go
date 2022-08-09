package consts

import "github.com/gin-gonic/gin"

type CodeError struct {
	Code    int
	Message string
}

var (
	CodeOk          = CodeError{200, "成功"}
	CodeErrorNoReg  = CodeError{1001, "账号未注册"}
	CodeErrorPasswd = CodeError{1002, "密码错误"}
	CodeErrorParam  = CodeError{1003, "参数错误"}
	CodeErrorHasReg = CodeError{1001, "账号已注册"}
)

func Code(code CodeError) gin.H {
	return gin.H{
		"code":    code.Code,
		"message": code.Message,
	}
}
