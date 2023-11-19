package api

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func InvalidArgument(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	invalidArgument := status.New(codes.InvalidArgument, "invalid argument")

	statusDetails, err := invalidArgument.WithDetails(badRequest)
	if err != nil {
		return invalidArgument.Err()
	}

	return statusDetails.Err()
}
