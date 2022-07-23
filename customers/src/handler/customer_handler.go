package handler

import (
	"errors"

	pb "sirclo-ecommerce-task/customers/protos/customer"
	"sirclo-ecommerce-task/customers/src/model"
	"sirclo-ecommerce-task/customers/src/usecase"

	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GrpcHandler
type GrpcHandler struct {
	customerUseCase usecase.CustomerUseCase
}

// NewGrpcHandler
func NewGrpcHandler(customerUseCase usecase.CustomerUseCase) *GrpcHandler {
	return &GrpcHandler{customerUseCase}
}

// Save
func (h *GrpcHandler) Save(ctx context.Context, data *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	return nil, nil
}

// FindByID
func (h *GrpcHandler) FindByID(arg *pb.QueryRequest, stream pb.CustomerService_FindByIDServer) error {
	id := arg.ID

	cstResult := <-h.customerUseCase.FindByID(id)

	if cstResult.Error != nil {
		return status.Error(codes.InvalidArgument, cstResult.Error.Error())
	}

	member, ok := cstResult.Result.(*model.Customer)

	if !ok {
		err := errors.New("Result is Not Member")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	response := &pb.CustomerResponse{
		ID:        member.ID,
		FirstName: member.FirstName,
		LastName:  member.LastName,
		Email:     member.Email,
		Password:  member.Password,

		BirthDate: member.BirthDate.String(),

		CreatedAt: member.CreatedAt.String(),
		UpdatedAt: member.UpdatedAt.String(),
	}

	if err := stream.Send(response); err != nil {
		status.Error(codes.Internal, err.Error())
	}

	return nil
}

// FindByEmail
func (h *GrpcHandler) FindByEmail(arg *pb.QueryRequest, stream pb.CustomerService_FindByEmailServer) error {
	email := arg.Email

	cstResult := <-h.customerUseCase.FindByEmail(email)

	if cstResult.Error != nil {
		return status.Error(codes.InvalidArgument, cstResult.Error.Error())
	}

	member, ok := cstResult.Result.(*model.Customer)

	if !ok {
		err := errors.New("Result is Not Member")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	response := &pb.CustomerResponse{
		ID:        member.ID,
		FirstName: member.FirstName,
		LastName:  member.LastName,
		Email:     member.Email,
		Password:  member.Password,

		BirthDate: member.BirthDate.String(),

		CreatedAt: member.CreatedAt.String(),
		UpdatedAt: member.UpdatedAt.String(),
	}

	if err := stream.Send(response); err != nil {
		status.Error(codes.Internal, err.Error())
	}

	return nil
}
