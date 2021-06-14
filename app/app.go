package app

import (
	"time"

	"github.com/GianGoulart/Clinica_backend/app/acompanhamento"
	"github.com/GianGoulart/Clinica_backend/app/comercial"
	"github.com/GianGoulart/Clinica_backend/app/financeiro"
	"github.com/GianGoulart/Clinica_backend/app/health"
	"github.com/GianGoulart/Clinica_backend/app/item"
	"github.com/GianGoulart/Clinica_backend/app/medicos"
	"github.com/GianGoulart/Clinica_backend/app/pacientes"
	"github.com/GianGoulart/Clinica_backend/app/procedimentos"
	"github.com/GianGoulart/Clinica_backend/app/session"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/sirupsen/logrus"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Health         health.App
	Item           item.App
	Session        session.App
	Paciente       pacientes.App
	Medico         medicos.App
	Procedimento   procedimentos.App
	Comercial      comercial.App
	Financeiro     financeiro.App
	Acompanhamento acompanhamento.App
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores *store.Container

	StartedAt time.Time
	Version   string
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	container := &Container{
		Health:         health.NewApp(opts.Stores, opts.Version, opts.StartedAt),
		Item:           item.NewApp(opts.Stores),
		Session:        session.NewApp(nil),
		Paciente:       pacientes.NewApp(opts.Stores),
		Medico:         medicos.NewApp(opts.Stores),
		Procedimento:   procedimentos.NewApp(opts.Stores),
		Comercial:      comercial.NewApp(opts.Stores),
		Financeiro:     financeiro.NewApp(opts.Stores, comercial.NewApp(opts.Stores)),
		Acompanhamento: acompanhamento.NewApp(opts.Stores),
	}

	logrus.Info("Registered -> App")

	return container

}
