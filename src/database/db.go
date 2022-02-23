package database

import (
	"time"

	"github.com/felipe-seo/br-api/src/models"
)

var db1 []models.Account

func Database() {
	db1 = append(db1, models.Account{Id: "1", Name: "Boris Fausto", Cpf: 12355567812, Secret: "LesPassword", Balance: 100000, Created_at: time.Now().UTC()})
}
