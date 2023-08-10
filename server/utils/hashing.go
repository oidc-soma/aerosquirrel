// Package utils provides utilities for the AeroSquirrel API.
package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashedPassword hashes a password.
func HashedPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash password: %w", err)
	}

	return string(hashed), nil
}

// CompareHashAndPassword compares a hashed password with a password.
func CompareHashAndPassword(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return fmt.Errorf("mismatched password: %w", err)
	}

	return nil
}
