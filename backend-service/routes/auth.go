package routes

import (
	"github.com/gin-gonic/gin"
	"webcodes.dev/admin/controllers"
)

// AuthRouter router for handling auth routes
func AuthRouter(r *gin.Engine) *gin.Engine {

	r.POST("/api/v1/user/login", controllers.Login)
	r.POST("/api/v1/user/signup", controllers.Signup)
	r.GET("/api/v1/user/getusers", controllers.Users)
	r.POST("/api/v1/user/getuser", controllers.User)
	r.POST("/api/v1/user/updateuser", controllers.UpdateUser)
	r.POST("/api/v1/user/deleteuser", controllers.DeleteUser)
	r.POST("/api/v1/user/refreshtoken", controllers.RefreshToken)

	return r
}
