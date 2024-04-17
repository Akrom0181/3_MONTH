package handler

import (
	"clone/3_exam/api/models"
	"clone/3_exam/config"
	"clone/3_exam/pkg/jwt"
	"clone/3_exam/pkg/logger"
	"clone/3_exam/service"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services service.IServiceManager
	Log      logger.ILogger
}

func NewStrg(services service.IServiceManager, log logger.ILogger) Handler {
	return Handler{
		Services: services,
		Log:      log,
	}
}

func handleResponse(c *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	if statusCode >= 100 && statusCode <= 199 {
		resp.Description = config.ERR_INFORMATION
	} else if statusCode >= 200 && statusCode <= 299 {
		resp.Description = config.SUCCESS
		log.Info("REQUEST SUCCEEDED", logger.Any("msg: ", msg), logger.Int("status: ", statusCode))
	} else if statusCode >= 300 && statusCode <= 399 {
		resp.Description = config.ERR_REDIRECTION
	} else if statusCode >= 400 && statusCode <= 499 {
		resp.Description = config.ERR_BADREQUEST
		log.Error("!!!!!!!! BAD REQUEST !!!!!!!!", logger.Any("error: ", msg), logger.Int("status: ", statusCode))
		fmt.Println("BAD REQUEST: "+msg, "reason: ", data)
	} else {
		resp.Description = config.ERR_INTERNAL_SERVER
		fmt.Println("INTERNAL SERVER ERROR: "+msg, "reason: ", data)
		log.Error("!!!!!!!! ERR_INTERNAL_SERVER !!!!!!!!", logger.Any("error: ", msg), logger.Int("status: ", statusCode))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.ParseUint(limitStr, 10, 30)
	if err != nil {
		return 0, err
	}

	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

func getAuthInfo(c *gin.Context) (models.AuthInfo, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models.AuthInfo{}, errors.New("unauthorized")
	}

	m, err := jwt.ExtractClaims(accessToken)
	if err != nil {
		return models.AuthInfo{}, err
	}

	role := m["user_role"].(string)
	if !(role == config.User_ROLE) {
		return models.AuthInfo{}, errors.New("unauthorized")
	}

	return models.AuthInfo{
		UserID:   m["user_id"].(string),
		UserRole: role,
	}, nil
}
