package restapi

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/repository/postgresql"
	"github.com/gifff/xenelectronic/service"
)

var (
	CategoryRepository contract.CategoryRepository
	CategoryService    contract.CategoryService

	appConfig AppConfig
)

type AppConfig struct {
	DSN string `long:"dsn" env:"DSN" description:"PostgreSQL Database connection string"`
}

func configureDependencies() {
	postgresDB := postgresql.New(appConfig.DSN)
	CategoryRepository = postgresql.NewCategory(postgresDB)
	CategoryService = service.NewCategory(CategoryRepository)
}
