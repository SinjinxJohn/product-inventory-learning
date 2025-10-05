package auth

import (
	"fmt"
	"time"

	"example.com/event-app/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":      userId,
		"exirationAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseAndValidateToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		//callback function which checks if the token was signed with the correct signature HMAC(HS256/HS512)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing metho")
		}
		return []byte(config.Envs.JWTSecret), nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid or expired token:%v", err)
	}

	//jwt consists of three parts header,payload(claims),signature
	//the below code extracts the claims and checks if it can be type asserted to a map
	//since the payload was used as a map {"userID":1}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token claims")
	}

	//here the userID is checked if it is in float64 format since jwt library converts ints to float
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user id in token")
	}

	//then here we return the userID after type conversion

	return int(userIDFloat), nil
}
