package logincontroller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vanneeza/e-commerce-lite/internal/domain/web"
	loginweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/login"
	customerservice "github.com/vanneeza/e-commerce-lite/internal/service/customer"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type tokenData struct {
	Token string `json:"token"`
}

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	customerService customerservice.CustomerService
}

func NewLoginController(customerService customerservice.CustomerService) LoginController {
	return &loginController{
		customerService: customerService,
	}
}

func (l *loginController) Login(ctx *gin.Context) {
	var login loginweb.LoginRequest
	jwtKey := os.Getenv("JWT_KEY")

	err := ctx.ShouldBindJSON(&login)
	helper.PanicError(err)

	getCustomer, err := l.customerService.GetPassword(ctx.Request.Context(), login.Password, login.NoHp)
	helper.PanicError(err)

	if login.NoHp == "" && login.Password == "" {
		result := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "user not found",
			Data:    "NULL",
		}
		ctx.JSON(http.StatusNotFound, result)
		return
	} else {
		match := helper.CheckPasswordHash(login.Password, getCustomer.Password)
		if !match {
			result := web.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: "wrong password",
				Data:    "NULL",
			}
			ctx.JSON(http.StatusBadRequest, result)
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = getCustomer.Id
		claims["name"] = getCustomer.Name
		claims["number_handphone"] = getCustomer.NoHp
		claims["address"] = getCustomer.Address
		claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

		var jwtKeyByte = []byte(jwtKey)
		tokenString, err := token.SignedString(jwtKeyByte)
		helper.PanicError(err)

		result := web.WebResponse{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "The customer has successfully logged in, hello " + getCustomer.Name,
			Data: tokenData{
				Token: tokenString,
			},
		}
		ctx.JSON(http.StatusOK, result)
	}

}
