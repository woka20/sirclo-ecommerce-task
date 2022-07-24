package main

import (
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"

	"sirclo-ecommerce-task/customers/db"

	customerGrpc "sirclo-ecommerce-task/customers/grpc"
	"sirclo-ecommerce-task/customers/src/handler"
	"sirclo-ecommerce-task/customers/src/query"
	"sirclo-ecommerce-task/customers/src/repo"
	"sirclo-ecommerce-task/customers/src/usecase"
)

const GrpcPortDefault = 3001

func main() {
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	//init member handler
	memberQuery := query.NewCustomerQueryInMemory(db.GetInMemoryDb())
	memberRepository := repo.NewCustomerRepositoryInMemory(db.GetInMemoryDb())

	memberUseCase := usecase.NewCustomerUseCase(memberRepository, memberQuery)

	memberGrpcHandler := handler.NewGrpcHandler(memberUseCase)
	//end init member handler

	grpcServer, err := customerGrpc.NewGrpcServer(memberGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Server Customer Running")

	err = grpcServer.Serve(GrpcPortDefault)

	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}

}
