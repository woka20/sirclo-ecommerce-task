package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	configEnv "github.com/joho/godotenv"

	// "github.com/wuriyanto48/ecommerce-grpc-microservice/order/middleware"

	"sirclo-ecommerce-task/order/middle"
	"sirclo-ecommerce-task/order/src/handler"
	grpcService "sirclo-ecommerce-task/order/src/services/grpc"
	"sirclo-ecommerce-task/order/src/usecase"
)

func main() {

	//load environtment variables
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Printf(".env is not loaded: %s", err.Error())
		os.Exit(2)
	}

	confirm, ok := os.LookupEnv("SIGNATURE")

	// publicKey, err := config.InitPublicKey()

	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	publicKey := []byte(confirm)

	// membership GRPC Host
	grpcCustomerHost, ok := os.LookupEnv("CUSTOMER_GRPC_HOST")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	// membership Auth Key
	grpcCustomerAuthKey, ok := os.LookupEnv("CUSTOMER_GRPC_AUTH_KEY")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	// product GRPC Host
	grpcProductHost, ok := os.LookupEnv("PRODUCT_GRPC_HOST")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	// product Auth Key
	grpcProductAuthKey, ok := os.LookupEnv("PRODUCT_GRPC_AUTH_KEY")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	//init order handler

	//init grpc service
	customerGrpcService, err := grpcService.NewCustomerGrpcClient(grpcCustomerHost, grpcCustomerAuthKey)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		os.Exit(1)
	}

	productGrpcService, err := grpcService.NewProductGrpcClient(grpcProductHost, grpcProductAuthKey)
	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	orderUseCase := usecase.NewOrderUseCase(customerGrpcService, productGrpcService)

	orderHttpHandler := handler.NewHttpOrderHandler(orderUseCase)

	//routing

	r := mux.NewRouter()

	r.Handle("/api/me", middle.LogRequest(middle.Bearer(publicKey, orderHttpHandler.Me()))).Methods("GET")

	r.Handle("/api/products", middle.LogRequest(middle.Bearer(publicKey, orderHttpHandler.GetProducts()))).Methods("GET")
	r.Handle("/api/products/{id}", middle.LogRequest(middle.Bearer(publicKey, orderHttpHandler.GetProduct()))).Methods("GET")

	http.ListenAndServe(":3004", r)
}
