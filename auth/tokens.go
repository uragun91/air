package auth

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(userId string, role string) (string, error) {

	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"));

	if err != nil {
		return "", err;
	}

	claims := jwt.MapClaims{};
	claims["userId"] = userId;
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix();
	claims["role"] = role;
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);

	return token.SignedString([]byte(os.Getenv("API_SECRET")));
}

func TokenValid(c *gin.Context) error {
	tokenString := extractToken(c);
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil;
	});

	if err != nil {
		return err;
	}

	return nil;
}

func extractToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}