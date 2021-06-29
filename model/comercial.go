package model

import (
	"encoding/json"
	"io"
)

type Comercial struct {
	Id                      string       `json:"id" bd:"id"`
	Id_Procedimento         string       `json:"id_procedimento" bd:"id_procedimento"`
	Procedimento            Procedimento `json:"procedimento" bd:"-"`
	Nome_Paciente           string       `json:"nome_paciente" bd:"-"`
	Nome_Medico             string       `json:"nome_medico" bd:"-"`
	Id_Medico_Part          string       `json:"id_medico_part" bd:"id_medico_part"`
	Nome_Medico_Part        string       `json:"nome_medico_part" bd:"-"`
	Funcao_Medico_Part      int64        `json:"funcao_medico_part" bd:"funcao_medico_part"`
	Funcao_Medico_Part_Desc string       `json:"funcao_medico_part_desc" bd:"-"`
	Qtd_Parcelas            int64        `json:"qtd_parcelas" bd:"qtd_parcelas"`
	Valor_Parcelas          float64      `json:"valor_parcelas" bd:"valor_parcelas"`
	Tipo_Pagamento          int64        `json:"tipo_pagamento" bd:"tipo_pagamento"`
	Tipo_Pagamento_Desc     string       `json:"tipo_pagamento_desc" bd:"-"`
	Forma_Pagamento         int64        `json:"forma_pagamento" bd:"forma_pagamento"`
	Forma_Pagamento_Desc    string       `json:"forma_pagamento_desc" bd:"-"`
	Data_Emissao_NF         int64        `json:"data_emissao_nf" bd:"data_emissao_nf"`
	Data_Vencimento         int64        `json:"data_vencimento" bd:"data_vencimento"`
	Data_Pagamento          int64        `json:"data_pagamento" bd:"data_pagamento"`
	Data_Compensacao        int64        `json:"data_compensacao" bd:"data_compensacao"`
	Plano_Contas            int64        `json:"plano_contas" bd:"plano_contas"`
	Plano_Contas_Desc       string       `json:"plano_contas_desc" bd:"-"`
	Conta                   int64        `json:"conta" bd:"conta"`
	Conta_Desc              string       `json:"conta_desc" bd:"-"`
	Valor_Ajuste            float64      `json:"valor_ajuste" bd:"valor_ajuste"`
	Valor_Liquido           float64      `json:"valor_liquido" bd:"valor_liquido"`
	Obs                     string       `json:"obs" bd:"obs"`
}

var MapTipoPagto = map[int64]string{
	1: "Reembolso",
	2: "Extra",
	3: "Particular",
}

var MapFormaPagto = map[int64]string{
	1: "Reembolso",
	2: "Dinheiro",
	3: "Transferencia",
	4: "Boleto",
	5: "Cheque",
	6: "Débito",
	7: "Crédito a Vista",
	8: "Crédito Parcelado",
}

var MapFuncao = map[int64]string{
	1: "CP",
	2: "AN",
	3: "IN",
	4: "1AX",
	5: "2AX",
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

func (me *Comercial) PreSave() {
	me.Id = NewId()
}

func (me *Comercial) PreencheComercial(r *Comercial) *Comercial {
	r.Tipo_Pagamento_Desc = MapTipoPagto[r.Tipo_Pagamento]
	r.Forma_Pagamento_Desc = MapFormaPagto[r.Forma_Pagamento]
	r.Funcao_Medico_Part_Desc = MapFuncao[r.Funcao_Medico_Part]
	r.Plano_Contas_Desc = MapPlanoContas[r.Plano_Contas]
	r.Conta_Desc = MapConta[r.Conta]

	return r
}

func ComercialFromJson(data io.Reader) (*Comercial, error) {
	decoder := json.NewDecoder(data)
	var o *Comercial
	err := decoder.Decode(&o)
	return o, err

}
