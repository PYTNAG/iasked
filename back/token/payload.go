package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID
	UserId    int32
	Username  string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewPayload(userId int32, username string, duration time.Duration) (*Payload, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	return &Payload{
		ID:        uuid,
		UserId:    userId,
		Username:  username,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}, nil
}
