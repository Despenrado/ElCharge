package v1

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func testHelper() *AuthController {
	return &AuthController{
		jwtKey: "testksey",
	}
}
func TestCreateTokenString(t *testing.T) {
	ac := testHelper()
	token, err := ac.createTokenString("testID")
	assert.Nil(t, err)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ac.jwtKey), nil
	})
	if !tkn.Valid {
		t.Error("invalid token")
	}
}

func TestRenewToken(t *testing.T) {
	ac := testHelper()
	token, err := ac.createTokenString("testID")
	assert.Nil(t, err)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ac.jwtKey), nil
	})
	token, err = ac.RenewToken(claims)
	assert.Nil(t, err)
	claims = &Claims{}
	tkn, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ac.jwtKey), nil
	})
	if !tkn.Valid {
		t.Error("invalid token")
	}
}
