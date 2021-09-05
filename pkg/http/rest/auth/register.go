package rest

import (
	"net/http"

	models "github.com/alifudin-a/go-todoapp/pkg/domain/models/auth"
	"github.com/alifudin-a/go-todoapp/pkg/http/response"
	services "github.com/alifudin-a/go-todoapp/pkg/services/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type register struct{}

func NewRegisterHandler() *register {
	return &register{}
}

func (*register) RegisterHandler(c *gin.Context) {
	var resp response.Response
	var register *models.Auth
	var req = new(models.Auth)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Gagal memvalidasi data!"
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

	arg := services.RegisterParams{Reg: models.Auth{
		Username: req.Username,
		Password: string(hashedPw),
		Fullname: req.Fullname,
		Email:    req.Email,
	}}

	exist := service.IsExist(services.IsExistParams{Username: req.Username})
	if exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Register gagal! Username sudah ada!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	register, err = service.Register(arg)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Register gagal! Data tidak lengkap!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Register berhasil!"
	resp.Data = map[string]interface{}{
		"resgister": register,
	}

	c.JSON(http.StatusOK, resp)
}
