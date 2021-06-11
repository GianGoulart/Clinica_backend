package model

import (
	"encoding/json"
	"io"
)

type Financeiro struct {
	Id                string  `json:"id" bd:"id"`
	Nome_Paciente     string  `json:"nome_paciente" bd:"-"`
	Nome_Medico       string  `json:"nome_medico" bd:"-"`
	Id_Comercial      string  `json:"id_comercial" bd:"id_comercial"`
	Desc_Comercial    string  `json:"desc_comercial" bd:"-"`
	Data_Pagamento    int64   `json:"data_pagamento" bd:"data_pagamento"`
	Data_Compensacao  int64   `json:"data_compensacao" bd:"data_compensacao"`
	Plano_Contas      int64   `json:"plano_contas" bd:"plano_contas"`
	Plano_Contas_Desc string  `json:"plano_contas_desc" bd:"-"`
	Conta             int64   `json:"conta" bd:"conta"`
	Conta_Desc        string  `json:"conta_desc" bd:"-"`
	Valor_Ajuste      float64 `json:"valor_ajuste" bd:"valor_ajuste"`
	Valor_Liquido     float64 `json:"valor_liquido" bd:"valor_liquido"`
	Obs               string  `json:"obs" bd:"obs"`
}

var MapPlanoContas = map[int64]string{
	1: "01.01.01 - Receitas Consultório",
	2: "01.01.02 - Receitas Hospital",
	3: "01.01.03 - Receitas Materiais",
	4: "01.02.01 - Ajuste Conciliação de Receitas",
	5: "01.02.02 - Faturamento Terceiros",
	6: "01.02.03 - Receitas Taxas de Cirurgia",
	7: "01.01.03 - Alugueis de Sala",
	8: "01.01.03 - Outras Receitas",
}

var MapConta = map[int64]string{
	1: "_Terceiros",
	2: "Din Caixa",
	3: "Din Cofre",
	4: "BB Clinica",
	5: "BB Nucleo",
	6: "Safra Cartões",
}

func (me *Financeiro) PreencheFinanceiro(r *Financeiro) *Financeiro {
	r.Plano_Contas_Desc = MapPlanoContas[r.Plano_Contas]
	r.Conta_Desc = MapConta[r.Conta]

	return r
}

func (me *Financeiro) PreSave() {
	me.Id = NewId()
}

func FinanceiroFromJson(data io.Reader) (*Financeiro, error) {
	decoder := json.NewDecoder(data)
	var o *Financeiro
	err := decoder.Decode(&o)
	return o, err

}
