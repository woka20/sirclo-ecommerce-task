package server

import (
	// "errors"
	"fmt"
	"net"

	// "os"

	// middleware "github.com/wuriyanto48/ecommerce-grpc-microservice/product/grpc/middleware"

	pbProduct "github.com/woka20/sirclo-ecommerce-task/products/protos/product"

	productHandlerPackage "github.com/woka20/sirclo-ecommerce-task/products/src/handler"

	"google.golang.org/grpc"
)

//Server data structure, grpc server model
type Server struct {
	// categoryGrpcHandler *productHandlerPackage.GrpcCategoryHandler
	productGrpcHandler *productHandlerPackage.GrpcProductHandler

	// grpcMiddleware      *middleware.Interceptor
}

// Serve insecure server/ no server side encryption
func (s *Server) Serve(port uint) error {
	address := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	//Register all sub server here
	// pbCategory.RegisterCategoryServiceServer(server, s.categoryGrpcHandler)
	pbProduct.RegisterProductServiceServer(server, s.productGrpcHandler)
	//end register server

	err = server.Serve(l)

	if err != nil {
		return err
	}

	fmt.Sprintf("Product GRPC Server running on PORT %d", port)

	return nil
}

//NewGrpcServer function, return: Pointer GRPC Server, or error otherwise
func NewGrpcServer(productGrpcHandler *productHandlerPackage.GrpcProductHandler) (*Server, error) {
	//init Auth Key

	// grpcAuthKey, ok := os.LookupEnv("GRPC_AUTH_KEY")
	// if !ok {
	// err := errors.New("you need to specify GRPC_AUTH_KEY in the environment variable")
	// return nil, err
	// }

	// grpcMiddleware := middleware.NewInterceptor(grpcAuthKey)

	return &Server{
		// categoryGrpcHandler: categoryGrpcHandler,
		productGrpcHandler: productGrpcHandler,
		// grpcMiddleware:      grpcMiddleware,
	}, nil

}
