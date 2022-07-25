package main

import (
	"fmt"
	"net/http"

	"github.com/woka20/sirclo-ecommerce-task/auth/db"
	"github.com/woka20/sirclo-ecommerce-task/auth/handler"

	"github.com/woka20/sirclo-ecommerce-task/auth/src/query"
	"github.com/woka20/sirclo-ecommerce-task/auth/src/usecase"
	"github.com/woka20/sirclo-ecommerce-task/order/middle"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	identityQuery := query.NewIdentityQueryInMemory(db.GetInMemoryDb())
	authUseCase := usecase.NewAuthUseCase(identityQuery)

	authHttp := handler.NewHttpAuthHandler(authUseCase)
	r.Handle("/api/login", middle.LogRequest(authHttp.HandlerLogin())).Methods("POST")

	http.ListenAndServe(":3009", r)

	fmt.Printf("Server Auth Running")

}
