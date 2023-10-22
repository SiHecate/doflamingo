package controllers

import (
	"doflamingo/database"
	"doflamingo/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func MerchantRegister(c *fiber.Ctx) error {
	var RegisterRequest struct {
		Username         string `json:"username"`
		Password         string `json:"password"`
		Password_confirm string `json:"password_confirm"`
	}

	if err := c.BodyParser(&RegisterRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request data",
		})
	}

	if RegisterRequest.Password != RegisterRequest.Password_confirm {
		c.JSON(fiber.Map{
			"Message": "Password do not match",
			"Status":  c.Status(fiber.StatusBadRequest),
		})
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(RegisterRequest.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	db_respone := model.User{
		Username: RegisterRequest.Username,
		Password: string(hashed_password),
		Balance:  0,
		Type:     "merchant",
	}

	if err := database.Conn.Create(&db_respone).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register account in database",
		})
	}

	return c.JSON(db_respone)

}

func MerchantLogin(c *fiber.Ctx) error {
	var LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&LoginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request data",
		})
	}

	var user model.User
	if err := database.Conn.Model(&user).Where("username = ?", LoginRequest.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginRequest.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	now := time.Now()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not log in",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 12),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	type login_response struct {
		Username string
		Type     string
		Balance  float64
	}

	response := login_response{
		Username: LoginRequest.Username,
		Type:     user.Type,
		Balance:  user.Balance,
	}

	return c.JSON(fiber.Map{
		"message":       "Success",
		"user_response": response,
	})
}
