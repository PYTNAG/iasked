package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/PYTNAG/iasked/token"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func GetPayload(ctx context.Context, maker *token.PasetoMaker) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("cannot get metadata from context")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	headerFields := strings.Fields(values[0])
	if len(headerFields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authorizationType := strings.ToLower(headerFields[0])
	if authorizationType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type %s", authorizationType)
	}

	accessToken := headerFields[1]
	payload, err := maker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}
