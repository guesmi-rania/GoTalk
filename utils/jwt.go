package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key pour signer les JWT (à changer en production)
var jwtSecret = []byte("TonSecretTrèsSecret123!")

// Claims personnalisés
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// Générer un JWT pour un utilisateur
func GenerateToken(userID uint, expirationHours int) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "chatterly",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Vérifier et parser un JWT
func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token invalide")
	}

	return claims, nil
}
