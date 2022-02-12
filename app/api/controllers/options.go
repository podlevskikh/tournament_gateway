package controllers

import (
	"github.com/gin-gonic/gin"
	"tournament_gateway/app/api/response_factory"
)

func OptionsHTTPHandler(c *gin.Context) { response_factory.ReturnOptions(c) }
