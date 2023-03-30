package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCreateUser = status.Error(codes.Internal, "create user error")
	ErrGetUser    = status.Error(codes.Internal, "get user error")
	ErrDeleteUser = status.Error(codes.Internal, "delete user error")
)
