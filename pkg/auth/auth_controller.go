package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"salaries/pkg/domain"
)

const (
	dummyUsername = "pmagnaghi"
	dummyPassword = "123456"
)

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type Controller interface {
	Login(context *gin.Context)
}

type authControllerImpl struct {
	authService Service
}

func NewAuthController(authService Service) Controller {
	return &authControllerImpl{
		authService: authService,
	}
}

func (c authControllerImpl) Login(context *gin.Context) {
	var input AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	// TODO: Here is necessary find user in db and validate user and password carefully
	// TODO: For the exercise only validate a dummy user
	user, err := c.validate(input.Username, input.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, err)
		return
	}

	jwt, err := c.authService.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, Token{jwt})
}

func (c authControllerImpl) validate(username, password string) (*domain.User, error) {
	if username == dummyUsername && password == dummyPassword {
		return &domain.User{
			ID:       1,
			Username: dummyUsername,
			Password: dummyPassword,
		}, nil
	}
	return nil, errors.New("invalid user")
}
