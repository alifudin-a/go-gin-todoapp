package routes

import (
	"os"

	rest "github.com/alifudin-a/go-todoapp/pkg/http/rest/auth"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/auth", rest.NewLoginHandler().LoginHandler)

	r.Run(":" + os.Getenv("APP_PORT"))
}
