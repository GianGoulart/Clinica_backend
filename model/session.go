package model

import (
	"context"
	"regexp"
	"strings"

	"github.com/tradersclub/TCUtils/logger"
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

// Is verifica se possui alguma das roles
func (s *Session) Is(roles ...string) bool {
	pattern := "(" + strings.Join(roles, "|") + ")"

	ok, err := regexp.MatchString(pattern, s.Roles)
	if err != nil {
		logger.Error("model.session.is", s, roles)

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
