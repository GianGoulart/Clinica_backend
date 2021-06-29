package model

import (
	"encoding/json"
	"errors"
	"io"
)

type Medico struct {
	Id            string `json:"id" bd:"id"`
	Nome          string `json:"nome" bd:"nome"`
	Especialidade string `json:"especialidade" bd:"especialidade"`
}

func (me *Medico) PreSave() {
	me.Id = NewId()
}

func (me *Medico) Validate() error {
	if len(me.Nome) == 0 {
		return errors.New("Necessário informar o nome")
	}
	if len(me.Especialidade) == 0 {
		return errors.New("Necessário informar a especialidade")
	}

	return nil
}

func MedicosFromJson(data io.Reader) (*Medico, error) {
	decoder := json.NewDecoder(data)
	var o *Medico
	err := decoder.Decode(&o)
	return o, err

}
