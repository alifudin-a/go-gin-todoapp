package services

import (
	"log"

	database "github.com/alifudin-a/go-todoapp/pkg/database/postgres"
	models "github.com/alifudin-a/go-todoapp/pkg/domain/models/auth"
	query "github.com/alifudin-a/go-todoapp/pkg/domain/query/auth"
)

type AuthService interface {
	Login(arg LoginParams) (*models.Auth, error)
	Register(arg RegisterParams) (*models.Auth, error)
	IsExist(arg IsExistParams) bool
	UpdateUserAccount(arg UpdateUserAccountParams) (*models.Auth, error)
	IsNotExist(arg IsNotExistParams) (bool, error)
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

type RegisterParams struct {
	Reg models.Auth
}

func (*service) Register(arg RegisterParams) (*models.Auth, error) {
	var db = database.PG
	var register models.Auth

	tx := db.MustBegin()
	err := tx.QueryRowx(query.Register, arg.Reg.Username, arg.Reg.Password, arg.Reg.Fullname, arg.Reg.Email).StructScan(&register)
	if err != nil {
		tx.Rollback()
		log.Println("[Register] An error occured: ", err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &register, nil
}

type IsExistParams struct {
	Username string
}

func (*service) IsExist(arg IsExistParams) bool {
	var db = database.PG
	var exist int

	err := db.Get(&exist, query.IsExist, arg.Username)
	if err != nil {
		log.Println("[IsExist] An error occured: ", err)
		return true
	}

	if exist == 1 {
		return true
	}

	return false
}

type IsNotExistParams struct {
	ID int
}

func (*service) IsNotExist(arg IsNotExistParams) (bool, error) {
	var db = database.PG
	var total int

	err := db.Get(&total, query.IsNotExist, arg.ID)
	if err != nil {
		log.Println("[IsNotExist] An error occured: ", err)
		return false, err
	}

	if total == 0 {
		return false, err
	}

	return true, nil
}

type UpdateUserAccountParams struct {
	User models.Auth
}

func (*service) UpdateUserAccount(arg UpdateUserAccountParams) (*models.Auth, error) {
	var db = database.PG
	var user models.Auth

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateUserAccount, arg.User.Username, arg.User.Password, arg.User.Fullname, arg.User.Email, arg.User.ID).StructScan(&user)
	if err != nil {
		tx.Rollback()
		log.Println("[Update] An error occured: ", err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
