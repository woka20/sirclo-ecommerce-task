package main

import (
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"

	inMemDB "sirclo-ecommerce-task/products/db"

	productGrpc "sirclo-ecommerce-task/products/grpc/server"
	productHandlerPackage "sirclo-ecommerce-task/products/src/handler"
	productQueryPackage "sirclo-ecommerce-task/products/src/query"
	productUsecasePackage "sirclo-ecommerce-task/products/src/usecase"
)

func main() {
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	//GrpcPortDefault default port for GRPC Server
	const GrpcPortDefault = 3002

	//init product grpc handler
	productQuery := productQueryPackage.NewProductQuery(inMemDB.GetProductInDb())
	productUseCase := productUsecasePackage.NewProductUseCase(productQuery)

	productGrpcHandler := productHandlerPackage.NewGrpcProductHandler(productUseCase)

	grpcServer, err := productGrpc.NewGrpcServer(productGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}

	err = grpcServer.Serve(GrpcPortDefault)

	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}

}