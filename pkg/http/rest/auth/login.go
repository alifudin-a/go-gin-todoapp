package rest

import (
	"net/http"
	"os"
	"time"

	models "github.com/alifudin-a/go-todoapp/pkg/domain/models/auth"
	"github.com/alifudin-a/go-todoapp/pkg/http/response"
	services "github.com/alifudin-a/go-todoapp/pkg/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type login struct{}

func NewLoginHandler() *login {
	return &login{}
}

type AuthTokenClaims struct {
	*jwt.StandardClaims
	models.Auth
}

func (*login) LoginHandler(c *gin.Context) {
	var resp response.Response
	var login *models.Auth
	var req = new(models.Auth)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Gagal memvalidasi data!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	service := services.NewAuthService()

	arg := services.LoginParams{Username: req.Username}

	login, err = service.Login(arg)
	if err != nil {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login gagal! Periksa kembali username anda!"
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(req.Password))
	if err != nil {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login gagal! Periksa kembali password anda!"
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &AuthTokenClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
		models.Auth{
			ID:       login.ID,
			Username: login.Username,
			Fullname: login.Fullname,
			Email:    login.Email,
		},
	}

	stringJwt, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "Gagal generate jwt"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Login Berhasil!"
	resp.Data = map[string]interface{}{
		"token": stringJwt,
	}

	c.JSON(http.StatusOK, resp)
}
