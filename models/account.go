package models

import "time"

type Account struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Cpf        int       `json:"cpf"`
	Secret     string    `json:"secret"`
	Balance    int64     `json:"balance"`
	Created_at time.Time `json:"created_at"`
}
