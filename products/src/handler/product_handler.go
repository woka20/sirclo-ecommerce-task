package handler

import (
	"errors"

	"sirclo-ecommerce-task/products/src/model"
	"sirclo-ecommerce-task/products/src/usecase"

	pb "sirclo-ecommerce-task/products/protos/product"

	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GrpcProductHandler
type GrpcProductHandler struct {
	productUseCase usecase.ProductUseCase
}

// NewGrpcProductHandler
func NewGrpcProductHandler(productUseCase usecase.ProductUseCase) *GrpcProductHandler {
	return &GrpcProductHandler{productUseCase}
}

// FindByID
func (h *GrpcProductHandler) FindByID(ctx context.Context, arg *pb.ProductQueryRequest) (*pb.ProductResponse, error) {

	id := arg.ID

	productResult := <-h.productUseCase.FindByID(int(id))

	if productResult.Error != nil {
		return nil, status.Error(codes.InvalidArgument, productResult.Error.Error())
	}

	product, ok := productResult.Result.(*model.Product)

	if !ok {
		err := errors.New("Result is Not Product")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	productResponse := &pb.ProductResponse{
		ID:          int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Stock:       int32(product.Stock),
		Price:       product.Price,
	}

	return productResponse, nil
}

//FindAll
func (h *GrpcProductHandler) FindAll(arg *pb.ProductQueryRequest, stream pb.ProductService_FindAllServer) error {

	productResult := <-h.productUseCase.FindAll()

	if productResult.Error != nil {
		return status.Error(codes.InvalidArgument, productResult.Error.Error())
	}

	products, ok := productResult.Result.(model.Products)

	if !ok {
		err := errors.New("Result is Not Products")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	for _, product := range products {

		productResponse := &pb.ProductResponse{
			ID:          int32(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Stock:       int32(product.Stock),
			Price:       product.Price,
		}

		if err := stream.Send(productResponse); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
