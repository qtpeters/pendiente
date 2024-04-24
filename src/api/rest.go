package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func NewRestApi() RestApi {
	return RestApi{
		logger: zap.Must(zap.NewDevelopment()),
		engine: gin.Default(),
	}
}

type RestApi struct {
	logger *zap.Logger
	engine *gin.Engine
}

func (api *RestApi) EnableAddTodo() {
	api.engine.GET("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Working": "true",
		})
	})
}

func (api *RestApi) Run() {
	err := api.engine.Run(":8080")
	if err != nil {
		api.logger.Fatal("Unable to start Web Engine")
	}
}
