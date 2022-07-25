package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/woka20/sirclo-ecommerce-task/auth/db"
	"github.com/woka20/sirclo-ecommerce-task/auth/handler"

	"github.com/woka20/sirclo-ecommerce-task/auth/src/query"
	"github.com/woka20/sirclo-ecommerce-task/auth/src/usecase"

	"github.com/gorilla/mux"
)

// LogRequest function, this function print Request log into console
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	identityQuery := query.NewIdentityQueryInMemory(db.GetInMemoryDb())
	authUseCase := usecase.NewAuthUseCase(identityQuery)

	authHttp := handler.NewHttpAuthHandler(authUseCase)
	r.Handle("/api/login", LogRequest(authHttp.HandlerLogin())).Methods("POST")

	http.ListenAndServe("0.0.0.0:3009", r)

	fmt.Printf("Server Auth Running")

}
