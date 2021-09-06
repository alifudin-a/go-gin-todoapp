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
	auth := v1.Group("/auth")

	auth.POST("/login", rest.NewLoginHandler().LoginHandler)
	auth.POST("/register", rest.NewRegisterHandler().RegisterHandler)
	auth.PUT("/user_account", rest.NewUpdateUserAccountHandler().UpdateUserAccountHandler)

	r.Run(":" + os.Getenv("APP_PORT"))
}
