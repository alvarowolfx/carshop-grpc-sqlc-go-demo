package auth

import "github.com/dgrijalva/jwt-go"

type AuthProvider interface {
	ParseToken(tokenString string) (*jwt.Token, error)
	UserClaimFromToken(token *jwt.Token) string
	UserScopesFromToken(token *jwt.Token) []string
}
