package restapi

import (
	"github.com/gifff/xenelectronic/contract"
	"github.com/gifff/xenelectronic/repository/postgresql"
	"github.com/gifff/xenelectronic/service"
	"github.com/gofrs/uuid"
)

var (
	CategoryRepository contract.CategoryRepository
	CartRepository     contract.CartRepository
	CategoryService    contract.CategoryService
	CartService        contract.CartService

	appConfig AppConfig
)

type AppConfig struct {
	DSN string `long:"dsn" env:"DSN" description:"PostgreSQL Database connection string"`
}

type uuidGenerator struct{}

func (u uuidGenerator) GenerateV4() string {
	return uuid.Must(uuid.NewV4()).String()
}

func configureDependencies() {
	postgresDB := postgresql.New(appConfig.DSN)
	CategoryRepository = postgresql.NewCategory(postgresDB)
	CartRepository = postgresql.NewCart(postgresDB, uuidGenerator{})
	CategoryService = service.NewCategory(CategoryRepository)
	CartService = service.NewCart(CartRepository)
}
