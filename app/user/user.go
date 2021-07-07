package user

import (
	"context"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/GianGoulart/Clinica_backend/store/user"
	"golang.org/x/crypto/bcrypt"
)

type App interface {
	GetUser(ctx context.Context, user *model.User) (*model.User, error)
	Set(ctx context.Context, user *model.User) (*model.User, error)
}

func NewApp(stores *store.Container) App {
	return appImpl{
		store: stores.User,
	}
}

type appImpl struct {
	store user.Store
}

func (s appImpl) GetUser(ctx context.Context, user *model.User) (*model.User, error) {

	res, err := s.store.GetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	plainPassword := []byte(user.Senha)
	// Assumes that u.Password is the actual hash and that you didn't store plain text password.
	err = bcrypt.CompareHashAndPassword([]byte(res.Senha), plainPassword)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s appImpl) Set(ctx context.Context, user *model.User) (*model.User, error) {
	user.Id = model.NewId()

	if err := user.Validate(); err != nil {
		return nil, err

	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Senha = string(hashedPassword)

	res, err := s.store.Set(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
