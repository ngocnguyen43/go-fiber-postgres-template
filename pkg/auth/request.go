package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"go-fiber-postgres-template/pkg/user"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/pbkdf2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Names    string `json:"names"`
}

func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	hash := pbkdf2.Key([]byte(password+os.Getenv("APP_SECRET")), salt, 4096, 32, sha256.New)
	return fmt.Sprintf("%s:%s", hex.EncodeToString(salt), hex.EncodeToString(hash))
}

func (r *RegisterRequest) bind(c *fiber.Ctx, u *user.User) error {
	if err := c.BodyParser(r); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	salt, err := generateSalt(14)

	if err != nil {
		log.Fatal(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	hash := hashPassword(r.Password, salt)
	u.Username = r.Username
	u.Password = hash
	u.Email = r.Email
	u.Names = r.Names
	return nil
}

func (r *RegisterRequest) Validate() {}

func VerifyPassword(storedPassword, inputPassword string) bool {
	parts := strings.Split(storedPassword, ":")
	if len(parts) != 2 {
		return false
	}

	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return false
	}

	storedHash, err := hex.DecodeString(parts[1])
	if err != nil {
		return false
	}

	inputHash := pbkdf2.Key([]byte(inputPassword+os.Getenv("APP_SECRET")), salt, 4096, 32, sha256.New)
	return subtle.ConstantTimeCompare(storedHash, inputHash) == 1
}
