package store

import (
	"github.com/GianGoulart/Clinica_backend/store/acompanhamento"
	"github.com/GianGoulart/Clinica_backend/store/comercial"
	"github.com/GianGoulart/Clinica_backend/store/health"
	"github.com/GianGoulart/Clinica_backend/store/item"
	"github.com/GianGoulart/Clinica_backend/store/medicos"
	"github.com/GianGoulart/Clinica_backend/store/pacientes"
	"github.com/GianGoulart/Clinica_backend/store/procedimentos"
	"github.com/GianGoulart/Clinica_backend/store/user"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	Health         health.Store
	Item           item.Store
	Paciente       pacientes.Store
	Medico         medicos.Store
	Procedimento   procedimentos.Store
	Comercial      comercial.Store
	Acompanhamento acompanhamento.Store
	User           user.Store
}

// Options struct de opções para a criação de uma instancia dos repositórios
type Options struct {
	Writer *sqlx.DB
	Reader *sqlx.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	container := &Container{
		Health:         health.NewStore(opts.Reader),
		Item:           item.NewStore(opts.Reader, opts.Writer),
		Paciente:       pacientes.NewStore(opts.Writer),
		Medico:         medicos.NewStore(opts.Writer),
		Procedimento:   procedimentos.NewStore(opts.Writer),
		Comercial:      comercial.NewStore(opts.Writer),
		Acompanhamento: acompanhamento.NewStore(opts.Writer),
		User:           user.NewStore(opts.Writer),
	}

	logrus.Info("Registered -> Store")

	return container
}
