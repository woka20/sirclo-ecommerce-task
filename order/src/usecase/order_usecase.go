package usecase

import (
	"errors"

	grpcService "github.com/woka20/sirclo-ecommerce-task/order/src/services/grpc"
	serviceModel "github.com/woka20/sirclo-ecommerce-task/order/src/services/model"
)

//orderUseCaseImpl
type orderUseCaseImpl struct {
	customerGrpcService grpcService.CustomerGrpcClient
	productGrpcService  grpcService.ProductGrpcClient
}

//NewOrderUseCase
func NewOrderUseCase(customerGrpcService grpcService.CustomerGrpcClient,
	productGrpcService grpcService.ProductGrpcClient) OrderUseCase {
	return &orderUseCaseImpl{
		customerGrpcService: customerGrpcService,
		productGrpcService:  productGrpcService,
	}
}

//FindMemberByID
func (u *orderUseCaseImpl) FindMemberByID(id string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		memberResult := <-u.customerGrpcService.FindByID(id)

		if memberResult.Error != nil {
			output <- UseCaseResult{Error: memberResult.Error}
			return
		}

		member, ok := memberResult.Result.(serviceModel.Customer)

		if !ok {
			err := errors.New("Result is not member")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: member}
	}()
	return output
}

//FindProductByID
func (u *orderUseCaseImpl) FindProductByID(id int) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		productResult := <-u.productGrpcService.FindByID(id)

		if productResult.Error != nil {
			output <- UseCaseResult{Error: productResult.Error}
			return
		}

		product, ok := productResult.Result.(serviceModel.Product)

		if !ok {
			err := errors.New("Result is not Product")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: product}
	}()
	return output
}

//FindProductAll
func (u *orderUseCaseImpl) FindProductAll() <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		productResult := <-u.productGrpcService.FindAll()

		if productResult.Error != nil {
			output <- UseCaseResult{Error: productResult.Error}
			return
		}

		products, ok := productResult.Result.(serviceModel.Products)

		if !ok {
			err := errors.New("Result is not Products")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: products}
	}()
	return output
}
