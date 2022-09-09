package database

import (
	"log"
	"os"

	"github.com/screamingarrow/gin-rest-api/env"
	"github.com/screamingarrow/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDb() {
	env.SetEnv()
	stringDeConexao := os.Getenv("CONNECTION_STRING")

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
