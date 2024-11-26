package controller

import (
	"echo_jwt/config"
	"echo_jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Struct for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}
type UserController interface {
	GetUserData(c echo.Context) error
	AddUserData(c echo.Context) error
	UpdateUserData(c echo.Context) error
	DeleteUserData(c echo.Context) error
	GetUser(c echo.Context) error
	AuthenticateUser(c echo.Context) error
}
type Usercontrollerimplementation struct{}

func NewUserController() UserController {
	return &Usercontrollerimplementation{}
}

func (u *Usercontrollerimplementation) GetUserData(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}

func (u *Usercontrollerimplementation) AddUserData(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

func (u *Usercontrollerimplementation) UpdateUserData(c echo.Context) error {
	var user models.User
	UserID := c.Param("id")

	if err := config.DB.First(&user, UserID).Error; err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"User not found"})
	}

	var updatedData models.User
	if err := c.Bind(&updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input"})
	}

	if updatedData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to hash password"})
		}
		updatedData.Password = string(hashedPassword)
	}

	config.DB.Model(&user).Updates(updatedData)
	return c.JSON(http.StatusOK, user)
}

func (u *Usercontrollerimplementation) DeleteUserData(c echo.Context) error {
	var user models.User
	UserID := c.Param("id")

	if err := config.DB.First(&user, UserID).Error; err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"User not found"})
	}

	config.DB.Delete(&user)
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (u *Usercontrollerimplementation) GetUser(c echo.Context) error {
	UserID := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, UserID).Error; err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *Usercontrollerimplementation) AuthenticateUser(c echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"Invalid input"})
	}

	var user models.User
	if err := config.DB.First(&user, "email = ?", credentials.Email).Error; err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{"Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	secret := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{"Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
