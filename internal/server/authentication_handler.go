package server

import (
	"go-fiber-postgres-template/internal/dtos"
	"go-fiber-postgres-template/internal/models"
	"go-fiber-postgres-template/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
//
// @Summary		Register a new user
// @Description	This endpoint is use for register a new user
// @Tags			Auth
// @Produce		json
// @Param			data	body		dtos.RegisterInput	true	"The input user struct"
// @Success		200  {object} 		dtos.RegisterResponse
// @Router			/auth/register [post]
func (s *FiberServer) RegisterHandler(c *fiber.Ctx) error {
	var input dtos.RegisterInput
	validate := s.validator
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code)
	}
	err := validate.Struct(input)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}
	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		FullName: input.FullName,
	}
	if err = s.db.GetInstance().Create(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

// LoginHandler godoc
//
// @Summary Login a user
// @Schemes
// @Description This endpoint allows a user to login by providing an email and password.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dtos.LoginInput true "Login Input"
// @Success 200 {object} dtos.LoginResponse
// @Router /auth/login [post]
func (s *FiberServer) LoginHandler(c *fiber.Ctx) error {
	var input dtos.LoginInput
	validate := s.validator
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}
	err := validate.Struct(input)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}
	var user models.User
	if err = s.db.GetInstance().Where("email = ?", input.Email).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	hashedPassword := user.Password
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}
	access, refresh, jti, err := utils.CreateTokenPair(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to create token pair")
	}
	refreshTokenFamily := models.RefreshTokenFamily{
		UserID: user.ID,
	}
	if err = s.db.GetInstance().Create(&refreshTokenFamily).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to create RefreshTokenFamily")
	}

	refreshToken := models.RefreshToken{
		JTI:                  jti,
		RefreshTokenFamilyID: refreshTokenFamily.ID,
	}
	if err = s.db.GetInstance().Create(&refreshToken).Error; err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to create RefreshToken")
	}

	return c.Status(fiber.StatusOK).JSON(dtos.LoginResponse{RefreshToken: refresh, AccessToken: access})
}

// RefreshTokenHandler godoc
// @Summary Refresh token
// @Schemes
// @Description This endpoint allows a user to refresh the access token by providing a refresh token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dtos.RefreshTokenInput true "Refresh Token Input"
// @Success 200 {object} dtos.LoginResponse
// @Router /auth/refresh-token [post]
func (s *FiberServer) RefreshTokenHandler(c *fiber.Ctx) error {
	var input dtos.RefreshTokenInput
	var validate = s.validator
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code)
	}
	if err := validate.Struct(input); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid token")
	}
	claims, err := utils.ValidateToken(input.RefreshToken)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid token")
	}
	if claims["typ"] != "refresh" {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid token type")
	}
	jti := claims["jit"].(string)
	var refreshToken models.RefreshToken

	if err = s.db.GetInstance().Where("jti = ?", jti).Preload("RefreshTokenFamily").
		First(&refreshToken).Error; err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid token")
	}

	if refreshToken.Status == models.Used {
		refreshToken.RefreshTokenFamily.Status = models.Inactive
		if err = s.db.GetInstance().
			Model(&refreshToken.RefreshTokenFamily).
			Updates(refreshToken.RefreshTokenFamily).Error; err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid token")
		}

		return fiber.NewError(fiber.ErrBadRequest.Code, "Invalid token")
	}

	if refreshToken.RefreshTokenFamily.Status == models.Inactive {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Token is inactive")
	}

	userID := claims["sub"].(float64)
	access, refresh, jti, err := utils.CreateTokenPair(uint(userID))
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code, "Failed to create token pair")
	}
	refreshToken.Status = models.Used
	refreshToken.Parent = nil
	if err = s.db.GetInstance().Model(&refreshToken).Updates(refreshToken).Error; err != nil {
		log.Println(err)
		return fiber.NewError(fiber.ErrBadRequest.Code, "Failed to update refresh token")
	}
	refreshTokenNew := models.RefreshToken{
		JTI:                  jti,
		RefreshTokenFamilyID: refreshToken.RefreshTokenFamilyID,
		Parent:               &refreshToken,
	}
	if err = s.db.GetInstance().Create(&refreshTokenNew).Error; err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Failed to create refresh token")
	}
	return c.Status(fiber.StatusOK).JSON(dtos.LoginResponse{RefreshToken: refresh, AccessToken: access})
}
