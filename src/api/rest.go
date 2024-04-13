package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


type RestApi struct {
	logger *zap.logger
	engine *gin.Engine
}

func newRestApi() *RestApi {
	return &RestApi{
		logger: zap.NewDevelopment(),
		engine: gin.Default()
	}
}