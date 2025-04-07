package handlers

import (
	"api-gateway/Storage"
	"api-gateway/Storage/Sqlite"
	"api-gateway/config"
	"api-gateway/domain"
	"api-gateway/gateway/Response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func SetupAuth(group *gin.RouterGroup, cfg *config.Config, storage Storage.Repo) {

	group.POST("/register", func(c *gin.Context) {
		RegisterUser(c, storage)
	})

	group.POST("/login", func(c *gin.Context) {
		LoginUser(c, cfg, storage)
	})
}

func RegisterUser(c *gin.Context, storage Storage.Repo) {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, Response.Err{Error: fmt.Sprintf("invalid credentials: %s", err)})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response.Err{
			Error: fmt.Sprintf("failed to hash password: %s", err),
		})
		return
	}

	user := domain.User{
		Email:    userInput.Email,
		Password: string(hash),
	}

	val := validator.New()
	err = val.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response.Err{Error: "invalid credentials"})
		return
	}

	user, err = storage.CreateUser(user)
	if err != nil {
		if errors.Is(err, Sqlite.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, Response.Err{Error: fmt.Sprintf("user already exists: %s", err)})
			return
		}
		c.JSON(http.StatusInternalServerError, Response.Err{Error: fmt.Sprintf("failed to create user: %s", err)})
		return
	}

	c.JSON(http.StatusOK, Response.Normal{Message: "user registered"})
}

func LoginUser(c *gin.Context, cfg *config.Config, storage Storage.Repo) {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, Response.Err{Error: fmt.Sprintf("invalid credentials: %s", err)})
		return
	}

	user, err := storage.GetUserByEmail(userInput.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response.Err{Error: fmt.Sprintf("failed to get user: %s", err)})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response.Err{Error: fmt.Sprintf("invalid credentials: %s", err)})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["ID"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	secretKey := []byte(cfg.JWTSecret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response.Err{Error: fmt.Sprintf("failed to sign token: %s", err)})
		return
	}

	c.JSON(http.StatusOK, Response.JWTToken{
		Token: tokenString,
	})
}
