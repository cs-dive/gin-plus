package interceptor

import (
	"github.com/archine/gin-plus/v3/resp"
	"github.com/gin-gonic/gin"
)

// GlobalExceptionInterceptor gin global exception interceptor
// add via gin middleware.
// thrown when the exception type is string and the BusinessException
func GlobalExceptionInterceptor(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			resp.DirectRespErr(context, r)
			context.Abort()
		}
	}()
	context.Next()
}
