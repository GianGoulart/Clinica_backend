package model

import (
	"encoding/json"
	"errors"
	"io"
)

type Paciente struct {
	Id         string `json:"id" bd:"id"`
	Cpf        string `json:"cpf" bd:"cpf"`
	Nome       string `json:"nome" bd:"nome"`
	Telefone   string `json:"telefone" bd:"telefone"`
	Convenio   string `json:"convenio" bd:"convenio"`
	Plano      string `json:"plano" bd:"plano"`
	Acomodacao string `json:"acomodacao" bd:"acomodacao"`
	Status     int64  `json:"status" bd:"status"`
}

func (me *Paciente) PreSave() {
	me.Id = NewId()
}

func (me *Paciente) Validate() error {
	if len(me.Nome) == 0 {
		return errors.New("Necessário informar o nome")
	}
	if len(me.Cpf) == 0 {
		return errors.New("Necessário informar o cpf")
	}

	return nil
}

func PacienteFromJson(data io.Reader) (*Paciente, error) {
	decoder := json.NewDecoder(data)
	var o *Paciente
	err := decoder.Decode(&o)
	return o, err

}
