package model

import (
	"encoding/json"
	"io"
)

type Acompanhamento struct {
	Id                      string `json:"id" bd:"id"`
	Id_Procedimento         string `json:"id_procedimento" bd:"id_procedimento"`
	Desc_Procedimento       string `json:"desc_procedimento" bd:"-"`
	Envio_Protocolo         int64  `json:"envio_protocolo" bd:"envio_protocolo"`
	Solicitacao_Previa      int64  `json:"solicitacao_previa" bd:"solicitacao_previa"`
	Confirmacao_Solicitacao int64  `json:"confirmacao_solicitacao" bd:"confirmacao_solicitacao"`
	Finalizacao_Previa      int64  `json:"finalizacao_previa" bd:"finalizacao_previa"`
	Status_Previa           int64  `json:"status_previa" bd:"status_previa"`
	DescStatusPrevia        string `json:"desc_status_previa" bd:"-"`
	Envio_Convenio          int64  `json:"envio_convenio" bd:"envio_convenio"`
	Liberacao               int64  `json:"liberacao" bd:"liberacao"`
	Repasse_Paciente        int64  `json:"repasse_paciente" bd:"repasse_paciente"`
	Repasse_Clinica         int64  `json:"repasse_clinica" bd:"repasse_clinica"`
	Status_Reembolso        int64  `json:"status_reembolso" bd:"status_reembolso"`
	DescStatusReembolso     string `json:"desc_status_reembolso" bd:"-"`
	Obs                     string `json:"obs" bd:"obs"`
}

var MapStatusPrevia = map[int64]string{
	1: "01-Pendente envio de protocolo paciente",
	2: "02-Pendente solicitação da prévia",
	3: "03-Confirmação da prévia atrasada",
	4: "04-Finalização da prévia atrasada",
}

var MapStatusReembolso = map[int64]string{
	1: "01-Aguardando finalização de prévia",
	2: "02-Envio para convênio atrasado",
	3: "03-Liberação convênio atrasada",
	4: "04-Repasse paciente atrasado",
	5: "05-Repasse clínica atrasado",
}

func (me *Acompanhamento) PreSave() {
	me.Id = NewId()
}

func (me *Acompanhamento) PreencheAcompanhamento(r *Acompanhamento) *Acompanhamento {
	r.DescStatusPrevia = MapStatusPrevia[r.Status_Previa]
	r.DescStatusReembolso = MapStatusReembolso[r.Status_Reembolso]

	return r
}

func AcompanhamentoFromJson(data io.Reader) (*Acompanhamento, error) {
	decoder := json.NewDecoder(data)
	var o *Acompanhamento
	err := decoder.Decode(&o)
	return o, err

}
