package middle

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	Email    string `json:"Email"`
	Password string `json:"Password"`
	ID       string `json:"ID"`
}

// Bearer this middleware function for verifying accessToken from Authorization Header
func Bearer(verifyKey []byte, next http.Handler) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		accessToken := req.Header.Get("Authorization")
		if accessToken == "" {
			http.Error(res, "No Token Provided", http.StatusUnauthorized)
			return
		}
		tokenSlice := strings.Split(accessToken, " ")
		if len(tokenSlice) < 2 {
			fmt.Println("0000")
			http.Error(res, "Token is not valid", http.StatusUnauthorized)
			return
		}

		if tokenSlice[0] != "Bearer" {
			fmt.Println("00001")
			http.Error(res, "Token is not valid", http.StatusUnauthorized)
			return
		}
		tokenString := tokenSlice[1]
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {

			return verifyKey, nil
		})
		fmt.Println(token)

		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			memberID := claims.ID
			req.Header.Add("CustomerId", memberID)
			fmt.Println("OK " + memberID)
			next.ServeHTTP(res, req)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("00002")
				http.Error(res, "Token is not valid", http.StatusUnauthorized)
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				http.Error(res, "Token expired", http.StatusUnauthorized)
				return
			} else {
				fmt.Println("00003")
				http.Error(res, "Token is not valid", http.StatusUnauthorized)
				return
			}

		} else {
			http.Error(res, "Token is not valid", http.StatusUnauthorized)
			return
		}
	})
}
