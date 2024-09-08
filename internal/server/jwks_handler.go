package server

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"go-fiber-postgres-template/internal/utils"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/jwk"
)

func (s *FiberServer) loadPublicKey() (interface{}, error) {
	pubKeyFile, err := os.Open(utils.PublicKeyPath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to open public key file")
	}
	defer pubKeyFile.Close()

	pubKeyBytes, err := io.ReadAll(pubKeyFile)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to read public key file")
	}

	block, _ := pem.Decode(pubKeyBytes)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to parse public key")
	}

	return pubKey, nil
}
func (s *FiberServer) JwksHandler(c *fiber.Ctx) error {
	pubKey, err := s.loadPublicKey()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	jwkKey, err := jwk.New(pubKey)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("failed to create JWK: %v", err))
	}

	jwkSet := jwk.NewSet()
	jwkSet.Add(jwkKey)
	jwkKey.Set(jwk.AlgorithmKey, "RS256")
	jwkKey.Set(jwk.KeyUsageKey, "sig")

	return c.Status(http.StatusOK).JSON(jwkSet)
}
