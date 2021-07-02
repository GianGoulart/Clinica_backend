package model

import (
	"encoding/json"
	"errors"
	"io"
)

type User struct {
	Id    string `json:"id" bd:"id"`
	Nome  string `json:"nome" bd:"nome"`
	Email string `json:"email" bd:"email"`
	Senha string `json:"senha" bd:"senha"`
	Roles string `json:"roles" bd:"roles"`
}

func (me *User) PreSave() {
	me.Id = NewId()
}

func (me *User) Validate() error {
	if len(me.Nome) == 0 {
		return errors.New("Necessário informar o médico")
	}
	if len(me.Senha) == 0 {
		return errors.New("Necessário informar o paciente")
	}
	if len(me.Roles) == 0 {
		return errors.New("Necessário informar o paciente")
	}

	return nil
}

func UserFromJson(data io.Reader) (*User, error) {
	decoder := json.NewDecoder(data)
	var o *User
	err := decoder.Decode(&o)
	return o, err

}
