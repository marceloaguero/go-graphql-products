package provider

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	validate *validator.Validate
)

type Usecase interface {
	Repository
}

type usecase struct {
	repository Repository
}

func NewUsecase(repo Repository) Usecase {
	return &usecase{
		repository: repo,
	}
}

func (u *usecase) Create(provider *Provider) (*Provider, error) {
	// Verify name uniqueness
	provider.Name = strings.TrimSpace(provider.Name)
	_, err := u.GetByName(provider.Name)
	if err == nil {
		return nil, errors.Errorf("UC - Create - Provider with name %s already exists", provider.Name)
	}

	validate := validator.New()
	err = validate.Struct(provider)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return nil, errors.Wrap(validationErrors, "UC - Create - Error during provider data validation")
	}

	provider.ID = uuid.NewString()
	provider, err = u.repository.Create(provider)
	if err != nil {
		return nil, errors.Wrap(err, "UC - Create - Error creating a new provider")
	}

	return provider, nil
}

func (u *usecase) GetByID(id string) (*Provider, error) {
	provider, err := u.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetByID - Error fetching a provider")
	}

	return provider, nil
}

func (u *usecase) GetByName(name string) (*Provider, error) {
	provider, err := u.repository.GetByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetByName - Error fetching a provider")
	}

	return provider, nil
}

func (u *usecase) GetAll() ([]*Provider, error) {
	providers, err := u.repository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "UC - GetAll - Error fetching all providers")
	}

	return providers, nil
}

func (u *usecase) Update(provider *Provider) (*Provider, error) {
	// Trim spaces
	provider.Name = strings.TrimSpace(provider.Name)

	formerProvider := &Provider{}

	// Verificar la unicidad del nombre
	formerProvider, err := u.GetByName(provider.Name)
	if (err == nil) && (formerProvider.ID != provider.ID) {
		return nil, errors.Errorf("UC - Update - Provider with name %s already exists", provider.Name)
	}

	validate = validator.New()
	if err := validate.Struct(provider); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return nil, errors.Wrap(validationErrors, "UC - Update - Error during provider data validation")
	}

	_, err = u.GetByID(provider.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "UC - Update - Provider with id %d does not exist", provider.ID)
	}

	provider, err = u.repository.Update(provider)
	if err != nil {
		return nil, errors.Wrapf(err, "UC - Update - Error updating provider with id %d", provider.ID)
	}

	return provider, nil
}

func (u *usecase) Delete(provider *Provider) error {
	_, err := u.GetByID(provider.ID)
	if err != nil {
		return errors.Wrapf(err, "UC - Delete - Provider with id %d does not exist", provider.ID)
	}

	if err := u.repository.Delete(provider); err != nil {
		return errors.Wrapf(err, "UC - Delete - Error deleting provider with id %d", provider.ID)
	}

	return nil
}
