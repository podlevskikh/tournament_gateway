package response_factory

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vollyemsk_tournament_gateway/app/api/response_error"
)

func ReturnSuccess(c *gin.Context, ret interface{}) {
	c.JSON(http.StatusOK, ret)
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
