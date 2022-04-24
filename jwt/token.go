package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Gen 生成Token
func Gen(claims map[string]interface{}, salt string, exp time.Duration) string {
	if exp != 0 {
		claims["exp"] = time.Now().Add(exp).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	signed, err := token.SignedString([]byte(salt))
	if err != nil {
		fmt.Println("Token General Failed:", err)
	}
	return signed
}

// Parse 解析Token
func Parse(otk, salt string) (claims map[string]interface{}) {
	token, err := jwt.Parse(otk, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("That's Not A Token:", err)
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println("Token Has Expired:", err)
				return
			} else {
				fmt.Println("Invalid Token:", err)
				return
			}
		} else {
			fmt.Println("Token Parse Failed:", err)
			return
		}
	}
	if !token.Valid {
		fmt.Println("Invalid Token:", err)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Parse Token Format Error:", err)
		return
	}
	return
}
