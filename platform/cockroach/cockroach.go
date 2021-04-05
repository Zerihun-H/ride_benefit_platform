package cockroach

import (
	"log"
	"rideBenefit/internal/constant/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CockroachPlatform interface {
	Open() *gorm.DB
	Migrate() error
}

type cockroachPlatform struct {
	dbURL string
}

func Initialize(dbURL string) CockroachPlatform {
	return &cockroachPlatform{
		dbURL: dbURL}
}

func (cp *cockroachPlatform) Open() *gorm.DB {
	db, err := gorm.Open(postgres.Open(cp.dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (cp *cockroachPlatform) Migrate() error {
	db := cp.Open()
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	if !db.Migrator().HasTable(&model.Driver{}) {
		err := db.Migrator().CreateTable(&model.Driver{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&model.Partner{}) {
		err := db.Migrator().CreateTable(&model.Partner{})
		if err != nil {
			return err
		}
	}
	return nil
}
