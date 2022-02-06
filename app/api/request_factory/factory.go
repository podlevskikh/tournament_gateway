package request_factory

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tournament_gateway/app/api/response_error"
	"tournament_gateway/app/api/response_factory"
)

func ReadJSONRequestBody(c *gin.Context, entity interface{}) bool {
	dec := json.NewDecoder(c.Request.Body)
	if err := dec.Decode(entity); err != nil {
		response_factory.ReturnError(c, response_error.ParseRequest)
		return false
	}
	return true
}
