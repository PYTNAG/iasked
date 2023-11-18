package api

import (
	db "github.com/PYTNAG/iasked/db/sqlc"
	pb "github.com/PYTNAG/iasked/pb/user"
	"github.com/PYTNAG/iasked/token"
	"github.com/PYTNAG/iasked/util"
)

type UserServer struct {
	pb.UnimplementedUserAPIServer

	config      util.Config
	store       db.Querier
	pasetoMaker *token.PasetoMaker
}

func NewUserServer(config util.Config, store db.Querier) (*UserServer, error) {
	maker, err := token.NewMaker(config.PasetoSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &UserServer{
		config:      config,
		store:       store,
		pasetoMaker: maker,
	}

	return server, nil
}
