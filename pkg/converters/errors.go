package converters

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidUID = status.Error(codes.InvalidArgument, "invalid user uid")
)
