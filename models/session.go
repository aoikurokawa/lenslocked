package models

import (
	"database/sql"
	"fmt"

	"github.com/Aoi1011/lenslocked/rand"
)

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

const (
	MinBytesPerToken = 32
)

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		UserID: userID,
		Token:  token,
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}