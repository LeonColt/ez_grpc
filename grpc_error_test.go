package ez_grpc_test

import (
	"testing"

	"github.com/LeonColt/ez_grpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func getBuilder(code codes.Code, message string) any {
	return &ez_grpc.GrpcErrorBuilder{
		GrpcCode: code,
		Message:  message,
	}
}

func TestCheckType(t *testing.T) {
	{
		err := getBuilder(codes.OK, "OK")
		res, ok := err.(ez_grpc.GrpcException)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, res)
	}

	{
		err := getBuilder(codes.NotFound, "Not Found")
		res, ok := err.(ez_grpc.GrpcException)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, res)
	}

	{
		err := getBuilder(codes.Internal, "Internal Server Error")
		res, ok := err.(ez_grpc.GrpcException)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, res)
	}
}

func TestGetCode(t *testing.T) {
	{
		err := getBuilder(codes.OK, "OK")
		res, ok := err.(ez_grpc.GrpcException)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, res)
		require.Equal(t, codes.OK, res.GetGrpcCode())
	}

	{
		err := getBuilder(codes.NotFound, "Not Found")
		res, ok := err.(ez_grpc.GrpcException)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, res)
		require.Equal(t, codes.NotFound, res.GetGrpcCode())
	}

	{
		err := getBuilder(codes.Internal, "Internal Server Error")
		res, ok := err.(ez_grpc.GrpcException)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, res)
		require.Equal(t, codes.Internal, res.GetGrpcCode())
	}
}

func TestHandleGrpcError(t *testing.T) {
	{
		err := ez_grpc.HandleGrpcError(nil)
		require.Nil(t, err)
	}

	{
		ezerr := getBuilder(codes.OK, "OK")
		err, ok := ezerr.(*ez_grpc.GrpcErrorBuilder)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, err)
		require.Equal(t, codes.OK, err.GetGrpcCode())
		grpcerr := ez_grpc.HandleGrpcError(err)
		require.Nil(t, grpcerr)
	}

	{
		ezerr := getBuilder(codes.NotFound, "Item was not found")
		err, ok := ezerr.(*ez_grpc.GrpcErrorBuilder)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, err)
		require.Equal(t, codes.NotFound, err.GetGrpcCode())
		grpcerr := ez_grpc.HandleGrpcError(err)
		require.EqualError(t, grpcerr, status.Error(codes.NotFound, "Item was not found").Error())
	}

	{
		ezerr := getBuilder(codes.Internal, "error parsing")
		err, ok := ezerr.(*ez_grpc.GrpcErrorBuilder)
		require.True(t, ok)
		require.IsType(t, &ez_grpc.GrpcErrorBuilder{}, err)
		require.Equal(t, codes.Internal, err.GetGrpcCode())
		grpcerr := ez_grpc.HandleGrpcError(err)
		require.EqualError(t, grpcerr, status.Error(codes.Internal, "error parsing").Error())
	}
}
