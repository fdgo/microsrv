package main

import (
	"fmt"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
	r := gin.Default()
	r.GET("/jwt", func(c *gin.Context) {
		j := &JWT{
			[]byte("test"),
		}
		claims := CustomClaims{
			1,
			"awh521",
			"1044176017@qq.com",
			jwt.StandardClaims{
				ExpiresAt: 15000, //time.Now().Add(24 * time.Hour).Unix()
				Issuer:    "test",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			c.Abort()
		}
		c.String(http.StatusOK, token+"---------------<br>")
		fmt.Println(token,"222222222")
		res, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				newToken, err := j.RefreshToken(token)
				if err != nil {
					c.String(http.StatusOK, err.Error())
				} else {
					c.String(http.StatusOK, newToken)
				}
			} else {
				c.String(http.StatusOK, err.Error())
			}
		} else {
			c.JSON(http.StatusOK, res)
		}
	})
	authorize := r.Group("/", JWTAuth())
	{
		authorize.GET("user", func(c *gin.Context) {
			claims := c.MustGet("claims").(*CustomClaims)
			fmt.Println(claims.ID)
			c.String(http.StatusOK, claims.Name)
		})
	}
	return r
}
