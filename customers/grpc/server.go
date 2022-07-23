package server

import (
	// "errors"
	"fmt"
	"net"

	// "os"

	// middleware "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/grpc/middleware"

	pbCustomer "sirclo-ecommerce-task/customers/protos/customer"

	customerHandlerPackage "sirclo-ecommerce-task/customers/src/handler"

	"google.golang.org/grpc"
)

//Server data structure, grpc server model
type Server struct {
	customerGrpcHandler *customerHandlerPackage.GrpcHandler
	// grpcMiddleware        *middleware.Interceptor
}

// Serve insecure server/ no server side encryption
func (s *Server) Serve(port uint) error {
	address := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer(
	//Unary interceptor
	// grpc.UnaryInterceptor(s.grpcMiddleware.Auth),
	//Stream interceptor
	// grpc.StreamInterceptor(s.grpcMiddleware.AuthStream),
	)

	//Register all sub server here
	pbCustomer.RegisterCustomerServiceServer(server, s.customerGrpcHandler)
	//end register server

	err = server.Serve(l)

	if err != nil {
		return err
	}

	fmt.Sprintf("Product GRPC Server running on PORT %d", port)

	return nil
}

//NewGrpcServer function, return: Pointer GRPC Server, or error otherwise
func NewGrpcServer(customerGrpcHandler *customerHandlerPackage.GrpcHandler) (*Server, error) {
	//init Auth Key

	// grpcAuthKey, ok := os.LookupEnv("GRPC_AUTH_KEY")
	// if !ok {
	// err := errors.New("you need to specify GRPC_AUTH_KEY in the environment variable")
	// return nil, err
	// }

	// grpcMiddleware := middleware.NewInterceptor(grpcAuthKey)

	return &Server{
		customerGrpcHandler: customerGrpcHandler,
		// grpcMiddleware:        grpcMiddleware,
	}, nil

}
