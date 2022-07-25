package main

import (
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"

	inMemDB "github.com/woka20/sirclo-ecommerce-task/products/db"

	productGrpc "github.com/woka20/sirclo-ecommerce-task/products/grpc/server"
	productHandlerPackage "github.com/woka20/sirclo-ecommerce-task/products/src/handler"
	productQueryPackage "github.com/woka20/sirclo-ecommerce-task/products/src/query"
	productUsecasePackage "github.com/woka20/sirclo-ecommerce-task/products/src/usecase"
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
	fmt.Printf("Server  Product Running.....")

	err = grpcServer.Serve(GrpcPortDefault)

	if err == nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}

}
