package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type authContextKey string

const (
	tokenInfoContextKey = authContextKey("token")
	userIDContextKey    = authContextKey("userID")
	scopesContextKey    = authContextKey("scopes")
)

func VerifyFuncForProvider(ap AuthProvider) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		tokenInfo, err := ap.ParseToken(token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		userID := ap.UserClaimFromToken(tokenInfo)
		scopes := ap.UserScopesFromToken(tokenInfo)

		ctx = context.WithValue(ctx, tokenInfoContextKey, token)
		ctx = context.WithValue(ctx, userIDContextKey, userID)
		ctx = context.WithValue(ctx, scopesContextKey, scopes)

		return ctx, nil
	}
}

func GetTokenFromContext(ctx context.Context) *jwt.Token {
	token, ok := ctx.Value(tokenInfoContextKey).(*jwt.Token)
	if !ok {
		return nil
	}
	return token
}

func GetUserIDFromContext(ctx context.Context) *jwt.Token {
	token, ok := ctx.Value(userIDContextKey).(*jwt.Token)
	if !ok {
		return nil
	}
	return token
}

func CheckScopeFromContext(ctx context.Context, scope string) bool {
	scopes, ok := ctx.Value(scopesContextKey).([]string)
	if !ok {
		return false
	}
	if !ok {
		return false
	}
	hasScope := false
	for i := range scopes {
		if scopes[i] == scope {
			hasScope = true
		}
	}
	return hasScope
}
