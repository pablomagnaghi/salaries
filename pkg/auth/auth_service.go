package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"salaries/pkg/domain"
	"salaries/pkg/logger"
	"strconv"
	"strings"
	"time"
)

const (
	AuthorizationHeader = "Authorization"
	TokenTTL            = "3600"
	JwtPrivateKey       = "DUMMY_KEY"
)

var privateKey = []byte(JwtPrivateKey)

type Service interface {
	GenerateJWT(user *domain.User) (string, error)
	VerifyToken(context *gin.Context) error
}

type authServiceImpl struct {
	logger logger.Logger
}

func NewAuthService(logger logger.Logger) Service {
	return &authServiceImpl{
		logger: logger,
	}
}

func (s authServiceImpl) GenerateJWT(user *domain.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(TokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func (s authServiceImpl) VerifyToken(context *gin.Context) error {
	token, err := s.getToken(context)

	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func (s authServiceImpl) getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := s.getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func (s authServiceImpl) getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get(AuthorizationHeader)
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
