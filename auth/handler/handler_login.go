package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"os"
	"time"

	"sirclo-ecommerce-task/auth/src/usecase"

	configEnv "github.com/joho/godotenv"

	// "github.com/golang-jwt/jwt"
	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	// Email    string `json:"Email"`
	// Password string `json:"Password"`
}

type HttpAuthHandler struct {
	authUseCase usecase.AuthUseCase
}

// NewHttpOrderHandler for initialise HttpOrderHandler model
func NewHttpAuthHandler(authUseCase usecase.AuthUseCase) *HttpAuthHandler {
	return &HttpAuthHandler{authUseCase: authUseCase}
}

func (a *HttpAuthHandler) HandlerLogin() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		email, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Invalid email or password", http.StatusBadRequest)
			return
		}

		err := configEnv.Load(".env")
		if err != nil {
			fmt.Printf(".env is not loaded: %s", err.Error())
			os.Exit(2)
		}

		appName, ok := os.LookupEnv("APPLICATION_NAME")

		if !ok {
			fmt.Printf("Error %e", err.Error())
			os.Exit(1)
		}

		expired := time.Duration(1) * time.Hour

		userInfo := <-a.authUseCase.AuthUser(email, password)
		if userInfo.Result == nil {
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}
		claims := MyClaims{
			StandardClaims: jwt.StandardClaims{
				Issuer:    appName,
				ExpiresAt: time.Now().Add(expired).Unix(),
			},
			// Email:    userInfo["email"].(string),
			// Password: userInfo["password"].(string),
		}

		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			claims,
		)

		signed, ok := os.LookupEnv("SIGNATURE")

		if !ok {
			fmt.Printf("Error %e", err.Error())
			os.Exit(1)
		}

		signedToken, err := token.SignedString(signed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tokenString, _ := json.Marshal(map[string]string{"token": signedToken})
		w.Write([]byte(tokenString))
	})
}
