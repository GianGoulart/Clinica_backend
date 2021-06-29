package model

import (
	"encoding/json"
	"io"
)

type Dashboard struct {
	StatusFinanceiro string `json:"status_financeiro"`
	StatusPrevia     string `json:"status_previa"`
	StatusReembolso  string `json:"status_reembolso"`
	NomePaciente     string `json:"nome_paciente"`
	NomeMedico       string `json:"nome_medico"`
	Procedimento     string `json:"procedimento"`
	DataProcedimento int64  `json:"data_procedimento"`
	IdProcedimento   string `json:"id_procedimento"`
}

func DashboardFromJson(data io.Reader) (*Dashboard, error) {
	decoder := json.NewDecoder(data)
	var o *Dashboard
	err := decoder.Decode(&o)
	return o, err

}
