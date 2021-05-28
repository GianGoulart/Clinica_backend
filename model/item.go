package model

import (
	"encoding/json"
	"errors"
	"io"
)

// Item de exemplo
type Item struct {
	ID             string            `json:"id"`
	Token          string            `json:"token"`
	CreateAt       int64             `json:"create_at"`
	ExpiresAt      int64             `json:"expires_at"`
	LastActivityAt int64             `json:"last_activity_at"`
	UserID         string            `json:"user_id"`
	DeviceID       string            `json:"device_id"`
	Roles          string            `json:"roles"`
	IsOAuth        bool              `json:"is_oauth"`
	Props          map[string]string `json:"props"`
}

// ToItem converte uma interface{} para *Item
func ToItem(data interface{}) (*Item, error) {
	value, ok := data.(*Item)
	if !ok {
		return nil, errors.New("não foi possível converter interface{} para *Item")
	}
	return value, nil
}

// ItemFromJson converte um io.Reader para *Item
func ItemFromJson(data io.Reader) *Item {
	decoder := json.NewDecoder(data)
	var o Item
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	}
	return nil

}
