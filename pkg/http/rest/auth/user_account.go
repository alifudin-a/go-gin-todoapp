package rest

import (
	"fmt"
	"net/http"

	models "github.com/alifudin-a/go-todoapp/pkg/domain/models/auth"
	"github.com/alifudin-a/go-todoapp/pkg/http/response"
	services "github.com/alifudin-a/go-todoapp/pkg/services/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type updatUserAccount struct {
}

func NewUpdateUserAccountHandler() *updatUserAccount {
	return &updatUserAccount{}
}

func (*updatUserAccount) UpdateUserAccountHandler(c *gin.Context) {
	var resp response.Response
	var userAccount *models.Auth
	var req = new(models.Auth)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errmsg := fmt.Sprintf("Gagal memvalidasi data! %v ", err)
		resp.Code = http.StatusBadRequest
		resp.Message = errmsg
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	service := services.NewAuthService()

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Gagal generate hash password!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	arg := services.UpdateUserAccountParams{User: models.Auth{
		ID:       req.ID,
		Username: req.Username,
		Password: string(hashedPw),
		Fullname: req.Fullname,
		Email:    req.Email,
	}}

	exist, _ := service.IsNotExist(services.IsNotExistParams{ID: req.ID})
	if !exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Update gagal! ID tidak ada!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	userAccount, err = service.UpdateUserAccount(arg)
	if err != nil {
		errmsg := fmt.Sprintf("Update gagal! Data tidak lengkap! %v", err)
		resp.Code = http.StatusBadRequest
		resp.Message = errmsg
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil mengubah data!"
	resp.Data = map[string]interface{}{
		"user_account": userAccount,
	}

	c.JSON(http.StatusOK, resp)
}
