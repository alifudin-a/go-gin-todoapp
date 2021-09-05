package services

import (
	"log"

	database "github.com/alifudin-a/go-todoapp/pkg/database/postgres"
	models "github.com/alifudin-a/go-todoapp/pkg/domain/models/auth"
	query "github.com/alifudin-a/go-todoapp/pkg/domain/query/auth"
)

type AuthService interface {
	Login(arg LoginParams) (*models.Auth, error)
}

type service struct{}

func NewAuthService() AuthService {
	return &service{}
}

type LoginParams struct {
	Username string
}

func (*service) Login(arg LoginParams) (*models.Auth, error) {

	var db = database.PG
	var login models.Auth

	err := db.Get(&login, query.Login, arg.Username)
	if err != nil {
		log.Println("[Login] An error occured: ", err)
		return nil, err
	}

	return &login, nil
}
