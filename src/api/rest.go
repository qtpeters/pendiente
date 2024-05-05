package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qtpeters/pendiente/data"
	"go.uber.org/zap"
	"net/http"
)

type TodoManager interface {
	CreateTodo(todo data.Todo) error
	UpdateTodo(todo data.Todo) error
	DeleteTodo(id string) error
}

func NewRestApi(manager TodoManager) RestApi {
	return RestApi{
		mgr:    manager,
		logger: zap.Must(zap.NewDevelopment()),
		engine: gin.Default(),
	}
}

type RestApi struct {
	mgr    TodoManager
	logger *zap.Logger
	engine *gin.Engine
}

func (api *RestApi) EnableCreateTodo() {
	api.engine.GET("/create", func(c *gin.Context) {

		var todo data.Todo
		if err := c.BindJSON(todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Unable to use the data sent to this endpoint",
			})
		} else {
			if err = api.mgr.CreateTodo(todo); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to write to data store",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "Success",
				})
			}
		}
	})
}

func (api *RestApi) EnableUpdateTodo() {
	api.engine.GET("/update", func(c *gin.Context) {

		var todo data.Todo
		if err := c.BindJSON(todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Unable to use the data sent to this endpoint",
			})
		} else {
			if err = api.mgr.UpdateTodo(todo); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to write to data store",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "Success",
				})
			}
		}
	})
}

func (api *RestApi) EnableDeleteTodo() {
	api.engine.DELETE("/delete", func(c *gin.Context) {
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
