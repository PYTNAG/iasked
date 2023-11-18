package api

import (
	"context"
	"database/sql"

	db "github.com/PYTNAG/iasked/db/sqlc"
	pb "github.com/PYTNAG/iasked/pb/user"
	"github.com/PYTNAG/iasked/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// get payload

	// validate request

	// authorize request

	nullableUsername := sql.NullString{
		Valid: req.Username != nil,
	}

	if nullableUsername.Valid {
		nullableUsername.String = req.GetUsername()
	}

	hash, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	params := db.CreateUserParams{
		Email:    req.GetEmail(),
		Username: nullableUsername,
		Hash:     hash,
	}

	var userInfo db.CreateUserRow
	userInfo, err = s.store.CreateUser(ctx, params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user with email %s already exists", params.Email)
			}
		}

		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	response := &pb.CreateUserResponse{
		Id: userInfo.ID,
	}

	return response, nil
}
