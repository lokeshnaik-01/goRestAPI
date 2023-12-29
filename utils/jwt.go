package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"errors"
)
const secretKey = "somesecret123"
func GenerateToken(email string, userId int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": "",
		"userId": "",
		"exp": time.Now().Add(time.Hour*2).Unix(),
	})
	// signedString should be of type byte slice
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string)  error{
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token)(interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signin method")
		}
		return secretKey, nil
	})
	if err!=nil {
		return errors.New("cant parse token")
	}
	if !(parsedToken.Valid) {
		return errors.New("Invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)\
	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }
	// no need to check ok as we know it'll be of type string
	// email, _ := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}