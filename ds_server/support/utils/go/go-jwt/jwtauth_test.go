package main

import (
	"fmt"
	"testing"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

type CustomClaimsTest struct {
	*CustomClaims
	siginKey string
	wanted   string
}
type ExpiredClaimsTest struct {
	CustomClaims
	siginKey string
}

var claims = []CustomClaimsTest{
	{
		&CustomClaims{
			1,
			"awh521",
			"1044176017@qq.com",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
				Issuer:    "test",
			},
		},
		"test",
		"",
	},
}
var expiredClaims = []ExpiredClaimsTest{
	{
		CustomClaims{
			1,
			"awh521",
			"1044176017@qq.com",
			jwt.StandardClaims{
				ExpiresAt: 1500,
				Issuer:    "test",
			},
		},
		"test",
	},
}
var jt *JWT = &JWT{
	[]byte("test"),
}
var foreverClaims CustomClaims = CustomClaims{
	1000,
	"defaul",
	"default@qq.com",
	jwt.StandardClaims{
		ExpiresAt: 0,
		Issuer:    "default",
	},
}

func TestCreateForeverTokens(t *testing.T) {
	token, err := jt.CreateToken(foreverClaims)
	assert.NoError(t, err)
	claims, err := jt.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), claims.StandardClaims.ExpiresAt)
}
func TestJWTCreateToken(t *testing.T) {
	for _, c := range claims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		token, err := j.CreateToken(*c.CustomClaims)
		assert.NoError(t, err)
		assert.IsType(t, "string", token)
		fmt.Println(token)
	}
}
func TestJWTParseToken(t *testing.T) {
	for _, c := range claims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		var err error
		c.wanted, err = j.CreateToken(*c.CustomClaims)
		result, err := j.ParseToken(c.wanted)
		assert.NoError(t, err)
		assert.Equal(t, c.CustomClaims.ID, result.ID)
		assert.Equal(t, c.CustomClaims.Email, result.Email)
		assert.Equal(t, c.CustomClaims.Name, result.Name)
		assert.Equal(t, c.CustomClaims.StandardClaims.ExpiresAt, result.StandardClaims.ExpiresAt)
		assert.Equal(t, c.CustomClaims.StandardClaims.Issuer, result.StandardClaims.Issuer)
	}
}
func TestRefreshToken(t *testing.T) {
	for _, c := range expiredClaims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		token, err := j.CreateToken(c.CustomClaims)
		assert.NoError(t, err)
		claims, err := j.ParseToken(token)
		assert.EqualError(t, err, TokenExpired.Error())
		assert.Nil(t, claims)
		token, err = j.RefreshToken(token)
		assert.NoError(t, err)
		assert.IsType(t, "string", token)
	}
}
