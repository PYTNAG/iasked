package api

import (
	shared_api "github.com/PYTNAG/iasked/api"
	pb "github.com/PYTNAG/iasked/pb/user"
	"github.com/PYTNAG/iasked/validation"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validation.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, shared_api.FieldViolation("email", err))
	}

	if err := validation.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, shared_api.FieldViolation("password", err))
	}

	if err := validation.ValidateUsername(req.GetUsername()); req.Username != nil && err != nil {
		violations = append(violations, shared_api.FieldViolation("username", err))
	}

	return
}
