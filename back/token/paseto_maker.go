package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

const randomSymmetricKey = "\x00"

const (
	usernameKey = "username"
	userIdKey   = "user_id"
	uuidKey     = "uuid"
)

type PasetoMaker struct {
	token paseto.Token
	key   paseto.V4SymmetricKey
}

func NewMaker(symmetricKey string) (*PasetoMaker, error) {
	maker := &PasetoMaker{
		token: paseto.NewToken(),
	}

	if symmetricKey == randomSymmetricKey {
		maker.key = paseto.NewV4SymmetricKey()
	} else {
		var err error
		maker.key, err = paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
		if err != nil {
			return nil, err
		}
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(userId int32, username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userId, username, duration)
	if err != nil {
		return "", payload, err
	}

	maker.token.SetIssuedAt(payload.IssuedAt)
	maker.token.SetExpiration(payload.ExpiredAt)
	maker.token.SetString(usernameKey, username)
	maker.token.SetString(uuidKey, payload.ID.String())
	maker.token.Set(userIdKey, userId)

	return maker.token.V4Encrypt(maker.key, nil), payload, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.NewParser()

	parser.AddRule(paseto.NotExpired())

	parsedToken, err := parser.ParseV4Local(maker.key, token, nil)
	if err != nil {
		return nil, err
	}

	payload := &Payload{}

	uuidString, err := parsedToken.GetString(uuidKey)
	if err != nil {
		return nil, err
	}

	payload.ID, err = uuid.Parse(uuidString)
	if err != nil {
		return nil, err
	}

	err = parsedToken.Get(userIdKey, &payload.UserId)
	if err != nil {
		return nil, err
	}

	payload.Username, err = parsedToken.GetString(usernameKey)
	if err != nil {
		return nil, err
	}

	payload.IssuedAt, err = parsedToken.GetIssuedAt()
	if err != nil {
		return nil, err
	}

	payload.ExpiredAt, err = parsedToken.GetExpiration()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
