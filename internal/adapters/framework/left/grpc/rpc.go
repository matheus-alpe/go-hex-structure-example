package rpc

import (
	"context"

	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetAddition(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetDivision(req.A, req.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}
