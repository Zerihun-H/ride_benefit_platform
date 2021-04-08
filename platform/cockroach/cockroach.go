package cockroach

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CockroachPlatform interface {
	Open() (*gorm.DB, error)
	Migrate() error
}

type cockroachPlatform struct {
	dbURL string
}

func Initialize(dbURL string) CockroachPlatform {
	return &cockroachPlatform{
		dbURL: dbURL}
}

func (cp *cockroachPlatform) Open() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(cp.dbURL), &gorm.Config{})
}

func (cp *cockroachPlatform) Migrate() error {
	db, err := cp.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	if !db.Migrator().HasTable(&model.Employee{}) {
		err := db.Migrator().CreateTable(&model.Employee{})
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

	if !db.Migrator().HasTable(&model.User{}) {
		err := db.Migrator().CreateTable(&model.User{})
		if err != nil {
			return err
		}
	}
	return nil
}
