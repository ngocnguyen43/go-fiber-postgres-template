package auth

import (
	"go-fiber-postgres-template/internal/database"
	"go-fiber-postgres-template/pkg/user"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"

	"github.com/gofiber/fiber/v2"
)

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Register a new user
//	@Tags			Auth
//	@Produce		json
//	@Param			data	body		RegisterRequest	true	"The input user struct"
//	@Success		200		{object}	AuthResponse
//	@Router			/register [post]
func Register(c *fiber.Ctx) error {
	db := database.New().GetInstance()

	user := new(user.User)

	register := &RegisterRequest{}

	if err := register.bind(c, user); err != nil {
		return err
	}
	err := db.Create(&user).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	refreshToken := RefreshToken{
		Jti:    uuid.NewString(),
		Parent: nil,
		Status: New,
	}
	err = db.Create(&refreshToken).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	at, rt, err := refreshToken.GenerateTokenPairs(*user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(AuthResponse{
		RefreshToken: rt,
		AccessToken:  at,
	})
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			Auth
//	@Produce		json
//	@Param			data	body		LoginRequest	true	"The input user struct"
//	@Success		200		{object}	AuthResponse
//	@Router			/login [post]
func Login(c *fiber.Ctx) error {
	db := database.New().GetInstance()

	loginRequest := &LoginRequest{}
	if err := c.BodyParser(&loginRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result := user.User{}
	db.Where(user.User{
		Email: loginRequest.Email,
	}).First(&result)
	refreshToken := RefreshToken{
		Jti:    uuid.NewString(),
		Parent: nil,
		Status: New,
	}
	verifyPassword := VerifyPassword(result.Password, loginRequest.Password)
	if !verifyPassword {
		return fiber.NewError(fiber.StatusBadRequest, "authenticate failed ")
	}
	err := db.Create(&refreshToken).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	at, rt, err := refreshToken.GenerateTokenPairs(result)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(AuthResponse{
		RefreshToken: rt,
		AccessToken:  at,
	})
}
