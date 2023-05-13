package provider

type Provider struct {
	ID       string `json:"id" gorm:"type:char(36);primaryKey"`
	Name     string `json:"name" gorm:"size:60" validate:"required,gte=2,lte=60"`
	IsActive bool   `json:"isActive,omitempty"`
}

type Repository interface {
	Create(provider *Provider) (*Provider, error)
	GetByID(id string) (*Provider, error)
	GetByName(name string) (*Provider, error)
	GetAll() ([]*Provider, error)
	Update(provider *Provider) (*Provider, error)
	Delete(provider *Provider) error
}
