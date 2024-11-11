package res

import "github.com/gin-gonic/gin"

type ResponseCapsule struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GlobalResponseHandler(ctx *gin.Context, code int, data any) {
	var (
		message ResMapKV
	)

	if code >= 200 && code < 400 {
		message = ResMapKV{
			HttpCode: code,
			Message:  "success",
		}
	} else {
		message = ResMapKV{
			HttpCode: code,
			Message:  "fail",
		}
	}
	if msg, exist := ResMap[code]; exist {
		message = msg
	}
	ctx.JSON(message.HttpCode, ResponseCapsule{
		Code: message.HttpCode,
		Msg:  message.Message,
		Data: data,
	})
	return
}
