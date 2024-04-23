package interceptor

import (
	"bytes"
	"github.com/archine/gin-plus/v3/exception"
	"github.com/archine/gin-plus/v3/plugin/logger"
	"github.com/archine/gin-plus/v3/resp"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

func printStack(err error) {
	var buf [2048]byte
	n := runtime.Stack(buf[:], false)
	log.Printf("%s %s", err.Error(), string(buf[:n]))
}

func printSimpleStack(err string) {
	var buf [2048]byte
	n := runtime.Stack(buf[:], false)
	lines := bytes.Split(buf[:n], []byte("\n"))
	log.Printf("%s\n%s", err, string(bytes.Join(lines[9:11], []byte("\n"))))
}

// GlobalExceptionInterceptor gin global exception interceptor
// add via gin middleware.
// thrown when the exception type is string and the BusinessException
func GlobalExceptionInterceptor(context *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case string:
				resp.DirectBadRequest(context, t)
			case *exception.BusinessException:
				printSimpleStack(t.Msg)
				resp.DirectRespWithCode(context, t.Code, t.Msg)
			case error:
				printStack(t)
				resp.SeverError(context, true)
			default:
				logger.Log.Error(r)
				resp.SeverError(context, true)
			}
			context.Abort()
		}
	}()
	context.Next()
}
