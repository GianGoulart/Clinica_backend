package model

import (
	"encoding/json"
	"errors"
	"io"
)

type Medico struct {
	Id           string `json:"id" bd:"id"`
	Nome         string `json:"nome" bd:"nome"`
	Cpf          string `json:"cpf" bd:"cpf"`
	Banco_pf     string `json:"banco_pf" bd:"banco_pf"`
	Agencia_pf   string `json:"agencia_pf" bd:"agencia_pf"`
	Conta_pf     string `json:"conta_pf" bd:"conta_pf"`
	Razao_social string `json:"razao_social" bd:"razao_social"`
	Banco_pj     string `json:"banco_pj" bd:"banco_pj"`
	Agencia_pj   string `json:"agencia_pj" bd:"agencia_pj"`
	Conta_pj     string `json:"conta_pj" bd:"conta_pj"`
	Cnpj         string `json:"cnpj" bd:"cnpj"`
}

func (me *Medico) PreSave() {
	me.Id = NewId()
}

func (me *Medico) Validate() error {
	if len(me.Nome) == 0 {
		return errors.New("Necessário informar o nome")
	}
	if len(me.Cpf) == 0 {
		return errors.New("Necessário informar o cpf")
	}

	return nil
}

func MedicosFromJson(data io.Reader) (*Medico, error) {
	decoder := json.NewDecoder(data)
	var o *Medico
	err := decoder.Decode(&o)
	return o, err

}
