package mysql_orm

import (
	"fmt"

	"github.com/marceloaguero/go-graphql-products/products/graph/model/product"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ormRepo struct {
	db *gorm.DB
}

// NewRepo crea un repositorio implementado en ORM (MySQL)
func NewRepo(dsName, dbName string) (product.Repository, error) {
	db, err := dbConnect(dsName, dbName)
	if err != nil {
		return nil, errors.Wrap(err, "MySQL ORM - Can't connect to DB")
	}

	db.AutoMigrate(&product.Product{})

	return &ormRepo{
		db: db,
	}, nil
}

func dbConnect(dsName, dbName string) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s/%s?charset=utf8&parseTime=True&loc=Local", dsName, dbName)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
