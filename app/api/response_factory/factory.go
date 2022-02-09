package response_factory

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tournament_gateway/app/api/response_error"
)

func ReturnSuccessList(c *gin.Context, ret interface{}, l int) {
	c.Header("Access-Control-Allow-Origin", "*")
	if l > 0 {
		c.Header("X-Total-Count", strconv.Itoa(l))
		c.Header("Access-Control-Expose-Headers", "X-Total-Count")
	}

	c.JSON(http.StatusOK, ret)
}

func ReturnSuccess(c *gin.Context, ret interface{}) {
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, ret)
}

func ReturnOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "X-PINGOTHER, Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
}

func ReturnError(c *gin.Context, re response_error.Error) {
	var status int
	switch re {
	case response_error.ParseRequest:
		status = http.StatusBadRequest
	case response_error.TournamentNotFound:
		status = http.StatusNotFound
	default:
		if re.Code == response_error.ValidationError.Code {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}
	}
	c.JSON(status, map[string]response_error.Error{"error": re})
}
