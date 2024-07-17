package ez_grpc_test

import (
	"errors"
	"testing"

	"github.com/LeonColt/ez"
	"github.com/LeonColt/ez_grpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func getBuilder(code ez.ErrorCode, message string) *ez.Error {
	return &ez.Error{
		Code:    code,
		Message: message,
	}
}

func TestHandleGrpcError(t *testing.T) {
	{
		err := ez_grpc.HandleGrpcError(nil)
		require.Nil(t, err)
	}

	{
		err := getBuilder(ez.ErrorCodeOk, "OK")
		require.Equal(t, codes.OK, ez_grpc.ParseErrorCodeToGrpcCode(err.Code))
		grpcerr := ez_grpc.HandleGrpcError(err)
		require.Nil(t, grpcerr)
	}
	{
		err := getBuilder(ez.ErrorCodeNotFound, "Item was not found")
		require.Equal(t, codes.NotFound, ez_grpc.ParseErrorCodeToGrpcCode(err.Code))
		grpcerr := ez_grpc.HandleGrpcError(err)
		require.EqualError(t, grpcerr, status.Error(codes.NotFound, "5: Item was not found").Error())
	}
	{
		err := getBuilder(ez.ErrorCodeInternal, "error parsing")
		require.Equal(t, codes.Internal, ez_grpc.ParseErrorCodeToGrpcCode(err.Code))
		grpcerr := ez_grpc.HandleGrpcError(err)
		require.EqualError(t, grpcerr, status.Error(codes.Internal, "13: error parsing").Error())
	}
	{
		err := getBuilder(ez.ErrorCodeUnauthenticated, "unauthenticated")
		joinedErr := errors.Join(err, errors.New("additional error"))
		grpcerr := ez_grpc.HandleGrpcError(joinedErr)
		require.EqualError(t, grpcerr, status.Error(codes.Unauthenticated, "16: unauthenticated").Error())
	}
}
