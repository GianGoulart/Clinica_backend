package store

import (
	"github.com/GianGoulart/Clinica_backend/store/acompanhamento"
	"github.com/GianGoulart/Clinica_backend/store/comercial"
	"github.com/GianGoulart/Clinica_backend/store/financeiro"
	"github.com/GianGoulart/Clinica_backend/store/health"
	"github.com/GianGoulart/Clinica_backend/store/item"
	"github.com/GianGoulart/Clinica_backend/store/medicos"
	"github.com/GianGoulart/Clinica_backend/store/pacientes"
	"github.com/GianGoulart/Clinica_backend/store/procedimentos"
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
	Financeiro     financeiro.Store
	Acompanhamento acompanhamento.Store
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
		Financeiro:     financeiro.NewStore(opts.Writer),
		Acompanhamento: acompanhamento.NewStore(opts.Writer),
	}

	logrus.Info("Registered -> Store")

	return container
}
