package user

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	errWrap "user-service/internal/common/error"
	"user-service/internal/common/response"
	errConstant "user-service/internal/constants/error"
	"user-service/internal/domain/dto/request"
	"user-service/internal/services/user"

	"github.com/gin-gonic/gin"
	customValidator "github.com/go-playground/validator/v10"
)

type Handler struct {
	userService user.Service
	validate    *customValidator.Validate
}

func NewHandler(userService user.Service, validate *customValidator.Validate) *Handler {
	return &Handler{
		userService: userService,
		validate:    validate,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.HTTPResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  errConstant.ErrBadRequest,
			Gin:  c,
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		response.HTTPResponse(response.ParamHTTPResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errWrap.ErrValidationResponse(err),
			Err:     err,
			Gin:     c,
		})
		return
	}

	if err := h.userService.Register(c.Request.Context(), user.CreateInput{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}); err != nil {
		if errors.Is(err, errConstant.ErrAlreadyExists) {
			response.HTTPResponse(response.ParamHTTPResponse{
				Code: http.StatusConflict,
				Err:  err,
				Gin:  c,
			})
			return
		}
		slog.Error("register user failed", "email", req.Email, "error", err)
		response.HTTPResponse(response.ParamHTTPResponse{
			Code: http.StatusInternalServerError,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HTTPResponse(response.ParamHTTPResponse{
		Code:    http.StatusCreated,
		Message: new("Register success"),
		Gin:     c,
	})
}

func (h *Handler) VerifyAccount(c *gin.Context) {
	token := strings.TrimSpace(c.Param("token"))
	if token == "" {
		response.HTTPResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  errConstant.ErrBadRequest,
			Gin:  c,
		})
		return
	}

	if err := h.userService.VerifyAccount(c.Request.Context(), token); err != nil {
		switch {
		case errors.Is(err, errConstant.ErrInvalidToken):
			response.HTTPResponse(response.ParamHTTPResponse{
				Code: http.StatusBadRequest,
				Err:  err,
				Gin:  c,
			})
			return
		case errors.Is(err, errConstant.ErrNotFound):
			response.HTTPResponse(response.ParamHTTPResponse{
				Code: http.StatusNotFound,
				Err:  err,
				Gin:  c,
			})
			return
		default:
			slog.Error("verify account failed", "token", token, "error", err)
			response.HTTPResponse(response.ParamHTTPResponse{
				Code: http.StatusInternalServerError,
				Err:  err,
				Gin:  c,
			})
			return
		}
	}
	response.HTTPResponse(response.ParamHTTPResponse{
		Code:    http.StatusOK,
		Message: new("Verify account successfully"),
		Gin:     c,
	})
}
