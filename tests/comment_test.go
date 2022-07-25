//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("someTestKey"))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}

func TestCreateComment(t *testing.T) {
	t.Run("should create a comment", func(t *testing.T) {
		client := resty.New()
		validToken := createToken()
		resp, err := client.R().
			SetHeader("Authorization", "bearer "+validToken).
			SetBody(`{"slug": "teste-slug", "author": "Vinny", "body": "Hello World"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("should not create comment without an valid token", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug": "teste-slug", "author": "Vinny", "body": "Hello World"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})
}
