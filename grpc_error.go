package ez_grpc

import (
	"errors"

	"github.com/LeonColt/ez"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func ParseErrorCodeToGrpcCode(in ez.ErrorCode) codes.Code {
	switch in {
	case ez.ErrorCodeOk:
		return codes.OK
	case ez.ErrorCodeCancelled:
		return codes.Canceled
	case ez.ErrorCodeUnknown:
		return codes.Unknown
	case ez.ErrorCodeInvalidArgument:
		return codes.InvalidArgument
	case ez.ErrorCodeDeadlineExceeded:
		return codes.DeadlineExceeded
	case ez.ErrorCodeNotFound:
		return codes.NotFound
	case ez.ErrorCodeConflict:
		return codes.AlreadyExists
	case ez.ErrorCodeNotAuthorized:
		return codes.PermissionDenied
	case ez.ErrorCodeResourceExhausted:
		return codes.ResourceExhausted
	case ez.ErrorCodeFailedPrecondition:
		return codes.FailedPrecondition
	case ez.ErrorCodeAborted:
		return codes.Aborted
	case ez.ErrorCodeOutOfRange:
		return codes.OutOfRange
	case ez.ErrorCodeUnimplemented:
		return codes.Unimplemented
	case ez.ErrorCodeInternal:
		return codes.Internal
	case ez.ErrorCodeUnavailable:
		return codes.Unavailable
	case ez.ErrorCodeDataLoss:
		return codes.DataLoss
	case ez.ErrorCodeUnauthenticated:
		return codes.Unauthenticated
	}
	return codes.Unknown
}

func HandleGrpcError(err error) error {
	if err == nil {
		return nil
	}
	var grpcErr *ez.Error
	if errors.As(err, &grpcErr) {
		return status.Error(ParseErrorCodeToGrpcCode(grpcErr.Code), grpcErr.Error())
	} else {
		return status.Errorf(codes.Internal, "unknown error: %#v", err)
	}
}
