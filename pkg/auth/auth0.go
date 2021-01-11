package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Auth0Provider struct {
	Identifier string
	Issuer     string
	JwksURI    string
}

func (ap *Auth0Provider) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		aud := ap.Identifier
		checkAud := token.Claims.(*CustomClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}
		// Verify 'iss' claim
		iss := ap.Issuer
		checkIss := token.Claims.(*CustomClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := ap.getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})

	return token, err
}

func (ap *Auth0Provider) UserClaimFromToken(token *jwt.Token) string {
	user, ok := token.Claims.(*CustomClaims).MapClaims["sub"]
	if ok {
		return string(user.(string))
	}
	return ""
}

func (ap *Auth0Provider) UserScopesFromToken(token *jwt.Token) []string {
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		result := strings.Split(claims.Scope, " ")
		return result
	}

	return []string{}
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type CustomClaims struct {
	Scope string `json:"scope"`
	//jwt.StandardClaims
	jwt.MapClaims
}

func (ap *Auth0Provider) getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(ap.JwksURI)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
