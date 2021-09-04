package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.New()

	r.Run(":" + os.Getenv("APP_PORT"))
}
