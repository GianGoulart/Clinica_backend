package user

import (
	"context"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/GianGoulart/Clinica_backend/store/user"
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

	return res, nil
}

func (s appImpl) Set(ctx context.Context, user *model.User) (*model.User, error) {
	user.Id = model.NewId()

	if err := user.Validate(); err != nil {
		return nil, err

	}
	res, err := s.store.Set(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
