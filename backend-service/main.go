package main

import (
	"github.com/gin-gonic/gin"
	"webcodes.dev/admin/routes"
)

func main() {
	r := gin.Default()
	r = routes.AuthRouter(r)
	r = routes.IndexRouter(r)
	r.Run()
}
