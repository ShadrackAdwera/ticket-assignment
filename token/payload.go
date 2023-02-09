package token

import (
	"fmt"
	"time"

	uuid "github.com/google/uuid"
)

var (
	ErrorExpiredToken = fmt.Errorf("your token is expired")
)

type TokenPayload struct {
	TokenId   uuid.UUID `json:"token_id"`
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(id int64, username string, duration time.Duration) (*TokenPayload, error) {
	tokenId, err := uuid.NewUUID()

	if err != nil {
		return &TokenPayload{}, nil
	}

	payload := &TokenPayload{
		TokenId:   tokenId,
		ID:        id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *TokenPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrorExpiredToken
	}
	return nil
}
