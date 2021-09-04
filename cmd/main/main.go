package main

import (
	database "github.com/alifudin-a/go-todoapp/pkg/database/postgres"
	"github.com/alifudin-a/go-todoapp/pkg/routes"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	database.OpenPG()
	routes.InitRoutes()
}
