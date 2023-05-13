package mysql_orm

import (
	"fmt"

	"github.com/marceloaguero/go-graphql-products/providers/model/provider"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ormRepo struct {
	db *gorm.DB
}

// NewRepo crea un repositorio implementado en ORM (MySQL)
func NewRepo(dsName, dbName string) (provider.Repository, error) {
	db, err := dbConnect(dsName, dbName)
	if err != nil {
		return nil, errors.Wrap(err, "MySQL ORM - Can't connect to DB")
	}

	db.AutoMigrate(&provider.Provider{})

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

func (r *ormRepo) Create(provider *provider.Provider) (*provider.Provider, error) {
	result := r.db.Create(&provider)
	return provider, result.Error
}

func (r *ormRepo) GetByID(id string) (*provider.Provider, error) {
	var provider provider.Provider
	result := r.db.Take(&provider, "id = ?", id)
	return &provider, result.Error
}

func (r *ormRepo) GetByName(name string) (*provider.Provider, error) {
	var provider provider.Provider
	result := r.db.Take(&provider, "name = ?", name)
	return &provider, result.Error
}

func (r *ormRepo) GetAll() ([]*provider.Provider, error) {
	providers := []*provider.Provider{}
	result := r.db.Find(&providers)
	return providers, result.Error
}

func (r *ormRepo) Update(provider *provider.Provider) (*provider.Provider, error) {
	result := r.db.Model(&provider).Updates(provider)
	if result.Error == nil {
		r.db.Model(&provider).Updates(map[string]interface{}{"is_active": provider.IsActive})
	}
	return provider, result.Error
}

func (r *ormRepo) Delete(provider *provider.Provider) error {
	result := r.db.Delete(&provider)
	return result.Error
}
