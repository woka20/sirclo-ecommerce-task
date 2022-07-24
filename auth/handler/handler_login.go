package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"os"
	"time"

	"sirclo-ecommerce-task/auth/src/model"
	"sirclo-ecommerce-task/auth/src/usecase"

	// "github.com/golang-jwt/jwt"
	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	Email    string `json:"Email"`
	Password string `json:"Password"`
	ID       string `json:"ID"`
}

type HttpAuthHandler struct {
	authUseCase usecase.AuthUseCase
}

// NewHttpOrderHandler for initialise HttpOrderHandler model
func NewHttpAuthHandler(authUseCase usecase.AuthUseCase) *HttpAuthHandler {
	return &HttpAuthHandler{authUseCase: authUseCase}
}

var _testing = false

func (a *HttpAuthHandler) HandlerLogin() http.Handler {

	var identityLogin *model.Identity
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&identityLogin)

		new := <-a.authUseCase.AuthUser(identityLogin.Email, identityLogin.Password)
		if new.Result == nil {
			http.Error(w, "Invalid email or password", http.StatusBadRequest)
			return
		}

		// err = configEnv.Load(".env")
		// if err != nil {
		// 	fmt.Printf(".env is not loaded: %s", err.Error())
		// 	os.Exit(2)
		// }

		appName, ok := os.LookupEnv("APPLICATION_NAME")
		fmt.Println(appName)
		if _testing {
			appName = "TEST"
		} else if !ok {
			// fmt.Printf("Error %e", err.Error())
			os.Exit(1)
		}

		expired := time.Duration(1) * time.Hour

		// userInfo := <-a.authUseCase.AuthUser(email, password)
		// if userInfo.Result == nil {
		// 	http.Error(w, "Invalid username or password", http.StatusBadRequest)
		// 	return
		// }
		fmt.Println("JLA3")
		claims := MyClaims{
			StandardClaims: jwt.StandardClaims{
				Issuer:    appName,
				ExpiresAt: time.Now().Add(expired).Unix(),
				Subject:   identityLogin.ID,
			},
			Email:    identityLogin.Email,
			Password: identityLogin.Password,
			ID:       identityLogin.ID,
		}

		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			claims,
		)

		signed, ok := os.LookupEnv("SIGNATURE")

		if _testing {
			signed = "TEST"
		} else if !ok {
			// fmt.Printf("Error %e", err.Error())
			os.Exit(1)
		}

		signedToken, err := token.SignedString([]byte(signed))
		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tokenString, _ := json.Marshal(map[string]string{"token": signedToken})
		w.Write([]byte(tokenString))
	})
}
