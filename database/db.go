package database

import (
	"log"

	"github.com/screamingarrow/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDb() {
	stringDeConexao := "host=localhost user=postgres password=admin dbname=alunos port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
