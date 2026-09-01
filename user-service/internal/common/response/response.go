package response

import (
	"user-service/internal/constants"
	errConstant "user-service/internal/constants/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data"`
	Token   *string     `json:"token"`
}

type ParamHTTPResponse struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    interface{}
	Token   *string
}

func HTTPResponse(param ParamHTTPResponse) {
	if param.Err == nil {
		param.Gin.JSON(param.Code, Response{
			Status:  constants.Success,
			Message: param.Message,
			Data:    param.Data,
			Token:   param.Token,
		})
		return
	}
	message := errConstant.ErrInternalServerError.Error()
	if param.Message != nil {
		message = *param.Message
	} else if param.Err != nil {
		message = param.Err.Error()
	}

	param.Gin.JSON(param.Code, Response{
		Status:  constants.Error,
		Message: message,
		Data:    param.Data,
	})
	return
}
