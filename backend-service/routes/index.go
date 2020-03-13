package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webkodes.com/admin/controllers"
)

// IndexRouter router for handling index routes
func IndexRouter(r *gin.Engine) *gin.Engine {

	r.GET("/api/v1/", controllers.Index)

	r.GET("/api/v1/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	return r
}
