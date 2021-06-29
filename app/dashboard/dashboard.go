package dashboard

import (
	"context"
	"time"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
)

type App interface {
	GetDash(ctx context.Context) (*[]model.Dashboard, error)
}

func NewApp(stores *store.Container) App {
	return appImpl{
		store: stores,
	}
}

type appImpl struct {
	store *store.Container
}

func (s appImpl) GetDash(ctx context.Context) (*[]model.Dashboard, error) {
	dash := new([]model.Dashboard)

	todayTimeDay := time.Now().UTC()
	todayTime := time.Date(todayTimeDay.Year(), todayTimeDay.Month(), todayTimeDay.Day(), 3, 0, 0, 0, time.UTC)

	mapAcompanhamento := make(map[string]model.Acompanhamento)

	mapComercial := make(map[string]model.Comercial)
	mapComercialValor := make(map[string]float64)

	acompanhamentos, err := s.store.Acompanhamento.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	comercial, err := s.store.Comercial.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range *acompanhamentos {
		mapAcompanhamento[v.Id_Procedimento] = v
	}

	for _, v := range *comercial {
		mapComercial[v.Id_Procedimento] = v

		mapComercialValor[v.Id_Procedimento] += v.Valor_Parcelas
	}

	procedimentos, err := s.store.Procedimento.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range *procedimentos {
		d := new(model.Dashboard)
		d.NomeMedico = p.Nome_Medico
		d.NomePaciente = p.Nome_Paciente
		d.Procedimento = p.NomeProcedimento
		d.DataProcedimento = p.Data
		d.IdProcedimento = p.Id

		//Status Previa
		if p.Status == 4 || p.Esteira == 2 || p.Esteira == 0 {
			d.StatusPrevia = " - "
		} else if p.Status == 3 && mapAcompanhamento[p.Id].Envio_Protocolo == 0 {
			d.StatusPrevia = "01-Pendente envio de protocolo paciente"
		} else if p.Status != 4 && mapAcompanhamento[p.Id].Solicitacao_Previa == 0 {
			d.StatusPrevia = "02-Pendente solicitação da prévia"
		} else if mapAcompanhamento[p.Id].Confirmacao_Solicitacao == 0 && time.Unix(mapAcompanhamento[p.Id].Solicitacao_Previa, 0).AddDate(0, 0, 2).Before(todayTime) {
			d.StatusPrevia = "03-Confirmação da prévia atrasada"
		} else if mapAcompanhamento[p.Id].Finalizacao_Previa == 0 && time.Unix(mapAcompanhamento[p.Id].Confirmacao_Solicitacao, 0).AddDate(0, 0, 7).Before(todayTime) {
			d.StatusPrevia = "04-Finalização da prévia atrasada"
		} else {
			d.StatusPrevia = "OK"
		}

		//Status Reembolso
		if p.Status != 3 || p.Esteira == 2 || p.Esteira == 0 {
			d.StatusReembolso = " - "
		} else if mapAcompanhamento[p.Id].Finalizacao_Previa == 0 {
			d.StatusReembolso = "01-Aguardando finalização de prévia"
		} else if mapAcompanhamento[p.Id].Finalizacao_Previa > 0 && mapAcompanhamento[p.Id].Envio_Convenio == 0 && time.Unix(p.Data, 0).AddDate(0, 0, 7).Before(todayTime) {
			d.StatusReembolso = "02-Envio para convênio atrasado"
		} else if p.Procedimento == 3 && mapAcompanhamento[p.Id].Envio_Convenio > 0 && mapAcompanhamento[p.Id].Liberacao == 0 && time.Unix(mapAcompanhamento[p.Id].Envio_Convenio, 0).AddDate(0, 0, 10).Before(todayTime) {
			d.StatusReembolso = "03-Liberação convênio atrasada"
		} else if p.Procedimento != 3 && mapAcompanhamento[p.Id].Envio_Convenio > 0 && mapAcompanhamento[p.Id].Liberacao == 0 && time.Unix(mapAcompanhamento[p.Id].Envio_Convenio, 0).AddDate(0, 0, 5).Before(todayTime) {
			d.StatusReembolso = "03-Liberação convênio atrasada"
		} else if mapAcompanhamento[p.Id].Liberacao > 0 && mapAcompanhamento[p.Id].Repasse_Paciente == 0 && time.Unix(mapAcompanhamento[p.Id].Liberacao, 0).AddDate(0, 0, 7).Before(todayTime) {
			d.StatusReembolso = "04-Repasse paciente atrasado"
		} else if mapAcompanhamento[p.Id].Repasse_Paciente > 0 && mapAcompanhamento[p.Id].Repasse_Clinica == 0 && time.Unix(mapAcompanhamento[p.Id].Repasse_Paciente, 0).AddDate(0, 0, 3).Before(todayTime) {
			d.StatusReembolso = "05-Repasse clínica atrasado"
		} else {
			d.StatusReembolso = "OK"
		}

		//Status Financeiro
		if (p.Status == 1 || p.Status == 2) && p.Data > 0 && p.Data < todayTime.Unix() {
			d.StatusFinanceiro = "Erro: Data de procedimento JÁ PASSOU, porém ainda consta como Agendado ou A Agendar"
		} else if p.Status != 3 && mapComercialValor[p.Id] > 0 {
			d.StatusFinanceiro = "OK"
		} else if p.Status != 3 && mapComercialValor[p.Id] > 0 {
			d.StatusFinanceiro = "Erro: Produção NÃO REALIZADA, porém constam valores no Financeiro"
		} else if p.Esteira == 0 {
			d.StatusFinanceiro = "Erro: Procedimento Realizado sem Esteira definida"
		} else if p.Valor != mapComercialValor[p.Id] {
			d.StatusFinanceiro = "Erro: Valor total do procedimento não bate com valores lançados no Financeiro"
		} else if p.Esteira == 1 && (mapComercial[p.Id].Tipo_Pagamento == 2 || mapComercial[p.Id].Tipo_Pagamento == 3 && mapComercialValor[p.Id] > 0) {
			d.StatusFinanceiro = "Erro: 100% Convenio COM valores Extra ou Particular no Financeiro"
		} else if p.Esteira == 2 && (mapComercial[p.Id].Tipo_Pagamento == 1 || mapComercial[p.Id].Tipo_Pagamento == 2 && mapComercialValor[p.Id] > 0) {
			d.StatusFinanceiro = "Erro: 100% Particular COM valores Extra ou Reembolso no Financeiro"
		} else if p.Esteira == 3 && (mapComercial[p.Id].Tipo_Pagamento == 1 && mapComercialValor[p.Id] == 0) {
			d.StatusFinanceiro = "Erro: Convenio+Extra SEM valores de Reembolso no Financeiro"
		} else if p.Esteira == 3 && (mapComercial[p.Id].Tipo_Pagamento == 2 && mapComercialValor[p.Id] == 0) {
			d.StatusFinanceiro = "Erro: Convenio+Extra SEM valores de Extra no Financeiro"
		} else if p.Esteira == 3 && (mapComercial[p.Id].Tipo_Pagamento == 3 && mapComercialValor[p.Id] > 0) {
			d.StatusFinanceiro = "Erro: Convenio+Extra COM valores de Particular no Financeiro"
		} else if (mapComercial[p.Id].Data_Emissao_NF == 0 && mapComercial[p.Id].Tipo_Pagamento == 1 && mapComercialValor[p.Id] > 0) ||
			(mapComercial[p.Id].Data_Emissao_NF == 0 && mapComercial[p.Id].Tipo_Pagamento == 2 && mapComercialValor[p.Id] > 0) ||
			(mapComercial[p.Id].Data_Emissao_NF == 0 && mapComercial[p.Id].Tipo_Pagamento == 3 && mapComercialValor[p.Id] > 0) {
			d.StatusFinanceiro = "Aviso: Existem NFs pendentes de Emissão"
		} else {
			d.StatusFinanceiro = "OK"
		}

		*(dash) = append(*(dash), *d)
	}

	return dash, nil
}
