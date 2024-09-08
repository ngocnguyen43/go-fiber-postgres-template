package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GenerateRSAKeys generates a new RSA private and public key pair
func GenerateRSAKeys(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	return privateKey, &privateKey.PublicKey, nil
}

// SavePEMKey saves a PEM-encoded RSA private key to a file
func SavePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer outFile.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	if err = pem.Encode(outFile, privateKeyPEM); err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

// SavePublicPEMKey saves a PEM-encoded RSA public key to a file
func SavePublicPEMKey(fileName string, pubkey *rsa.PublicKey) error {
	pubASN1, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %w", err)
	}

	pubPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	}

	outFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer outFile.Close()

	if err = pem.Encode(outFile, pubPEM); err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

// EnsureRSAKeys ensures that RSA keys exist in the specified directory
func EnsureRSAKeys(dir string, bits int) (*rsa.PrivateKey, error) {
	privateKeyPath := filepath.Join(dir, "private_key.pem")
	publicKeyPath := filepath.Join(dir, "public_key.pem")

	if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
		log.Println("Private key not found, generating new keys...")

		privateKey, publicKey, generationError := GenerateRSAKeys(bits)
		if generationError != nil {
			return nil, errors.New("failed to generate RSA keys")
		}

		if err = SavePEMKey(privateKeyPath, privateKey); err != nil {
			return nil, errors.New("failed to save private key")
		}

		if err = SavePublicPEMKey(publicKeyPath, publicKey); err != nil {
			return nil, errors.New("failed to save public key")
		}

		log.Println("Keys generated and saved successfully.")
		return privateKey, nil
	}
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, errors.New("failed to read private key file")
	}

	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key")
	}

	log.Println("Keys already exist.")
	return privateKey, nil
}
func DeleteRSAKeys(dir string) error {
	if err := os.Remove(dir); err != nil {
		return fmt.Errorf("failed to delete directory: %w", err)
	}

	return nil
}
