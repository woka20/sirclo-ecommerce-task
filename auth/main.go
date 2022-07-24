package main

import (
	"net/http"
	"sirclo-ecommerce-task/auth/db"
	"sirclo-ecommerce-task/auth/handler"

	"sirclo-ecommerce-task/auth/src/query"
	"sirclo-ecommerce-task/auth/src/usecase"
	"sirclo-ecommerce-task/order/middle"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	identityQuery := query.NewIdentityQueryInMemory(db.GetInMemoryDb())
	authUseCase := usecase.NewAuthUseCase(identityQuery)

	authHttp := handler.NewHttpAuthHandler(authUseCase)
	r.Handle("/api/login", middle.LogRequest(authHttp.HandlerLogin()))

	http.ListenAndServe(":3009", r)

}
