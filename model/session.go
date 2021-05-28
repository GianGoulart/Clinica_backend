package model

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Session modelo padrão de sessão
type Session struct {
	ID             string            `json:"id"`
	Token          string            `json:"token"`
	CreateAt       int64             `json:"create_at"`
	ExpiresAt      int64             `json:"expires_at"`
	LastActivityAt int64             `json:"last_activity_at"`
	UserID         string            `json:"user_id"`
	DeviceID       string            `json:"device_id"`
	Roles          string            `json:"roles"`
	IsOAuth        bool              `json:"is_oauth"`
	Props          map[string]string `json:"props"`
}

// Cache é a interface do pacote de cache
type Cache interface {
	Get(ctx context.Context, key string, v interface{}) error
	Set(ctx context.Context, key string, v interface{}) error

	Del(ctx context.Context, key string) error

	WithExpiration(d time.Duration) Cache
}

// Options struct de opções para a criação de uma instancia do cache
type Options struct {
	Expiration time.Duration
	URL        string
	Password   string
	Timeout    time.Duration
}

// Is verifica se possui alguma das roles
func (s *Session) Is(roles ...string) bool {
	pattern := "(" + strings.Join(roles, "|") + ")"

	ok, err := regexp.MatchString(pattern, s.Roles)
	if err != nil {
		logrus.Error("model.session.is", s, roles)

		return false
	}

	return ok
}

type contextKey string

const sessionKey = contextKey("session")

// SetSession sobrescreve o contexto com a sessão
func SetSession(ctx context.Context, sess *Session) context.Context {
	return context.WithValue(ctx, sessionKey, sess)
}

// GetSession pega a sessão do context
func GetSession(ctx context.Context) Session {
	sess, ok := ctx.Value(sessionKey).(*Session)
	if !ok {
		return Session{}
	}
	return *sess
}
