package grpc

//ServiceResult
type ServiceResult struct {
	Result interface{}
	Error  error
}

type CustomerGrpcClient interface {
	FindByID(id string) <-chan ServiceResult
	FindByEmail(email string) <-chan ServiceResult
}

//ProductGrpcClient interface
type ProductGrpcClient interface {
	FindByID(id int) <-chan ServiceResult
	FindAll() <-chan ServiceResult
}
