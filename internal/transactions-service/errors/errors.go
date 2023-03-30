package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCreateDeposit     = status.Error(codes.Internal, "create deposit transaction error")
	ErrCreateWithdrawal  = status.Error(codes.Internal, "create withdrawal transaction error")
	ErrGetTransaction    = status.Error(codes.Internal, "get transaction error")
	ErrDeleteTransaction = status.Error(codes.Internal, "delete transaction error")
)
