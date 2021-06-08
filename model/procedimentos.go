package model

import (
	"encoding/json"
	"errors"
	"io"
)

type Procedimento struct {
	Id                 string  `json:"id" bd:"id"`
	Id_Paciente        string  `json:"id_paciente" bd:"id_paciente"`
	Nome_Paciente      string  `json:"nome_paciente" bd:"-"`
	Id_Medico          string  `json:"id_medico" bd:"id_medico"`
	Nome_Medico        string  `json:"nome_medico" bd:"-"`
	Desc_Procedimento  string  `json:"desc_procedimento" bd:"desc_procedimento"`
	Procedimento       int64   `json:"procedimento" bd:"procedimento"`
	NomeProcedimento   string  `json:"nome_procedimento" bd:"-"`
	Local_Procedimento int64   `json:"local_procedimento" bd:"local_procedimento"`
	NomeLocal          string  `json:"nome_local" bd:"-"`
	Status             int64   `json:"status" bd:"status"`
	NomeStatus         string  `json:"nome_status" bd:"-"`
	Data               int64   `json:"data" bd:"data"`
	Valor              float64 `json:"valor" bd:"valor"`
	Esteira            int64   `json:"esteira" bd:"esteira"`
	DescEsteira        string  `json:"desc_esteira" bd:"-"`
}

var MapProcedimento = map[int64]string{
	1: "CONSULTA",
	2: "EXAMES",
	3: "CIRURGIA",
	4: "RETORNO",
	5: "ESCLERO",
	6: "CURATIVO",
	7: "COMBOS",
	8: "VENDAS DE PRODUTOS",
	9: "TAXAS E LOCAÇÃO",
}

var MapLocal = map[int64]string{
	1: "Clinica Abrão",
	2: "Albert Einstein",
	3: "Lefort Liberdade",
	4: "Lefort Morumbi",
	5: "São Luiz Jabaquara",
	6: "São Luiz Morumbi",
	7: "São Luiz Itaim",
	8: "Vila Nova Star",
}

var MapStatusProcedimento = map[int64]string{
	1: "A Agendar",
	2: "Agendado",
	3: "Realizado",
	4: "Cancelado",
}

var MapStatusEsteira = map[int64]string{
	1: "Convenio",
	2: "Particular",
	3: "Misto",
}

func (me *Procedimento) PreSave() {
	me.Id = NewId()
}

func (me *Procedimento) Validate() error {
	if len(me.Id_Medico) == 0 {
		return errors.New("Necessário informar o médico")
	}
	if len(me.Id_Paciente) == 0 {
		return errors.New("Necessário informar o paciente")
	}
	if me.Procedimento == 0 {
		return errors.New("Necessário informar o procedimento")
	}
	if me.Local_Procedimento == 0 {
		return errors.New("Necessário informar o local")
	}

	return nil
}

func ProcedimentosFromJson(data io.Reader) (*Procedimento, error) {
	decoder := json.NewDecoder(data)
	var o *Procedimento
	err := decoder.Decode(&o)
	return o, err

}

func (me *Procedimento) PreencheProcedimentos(r *Procedimento) *Procedimento {
	r.NomeProcedimento = MapProcedimento[r.Procedimento]
	r.NomeLocal = MapLocal[r.Local_Procedimento]
	r.DescEsteira = MapStatusEsteira[r.Esteira]
	r.NomeStatus = MapStatusProcedimento[r.Status]

	return r
}
