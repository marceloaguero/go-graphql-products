package product

type Product struct {
	ID          string  `json:"id" gorm:"type:char(36);primaryKey"`
	Name        string  `json:"name" gorm:"size:60" validate:"required,gte=2,lte=60"`
	Description string  `json:"description,omitempty" gorm:"size:250" validate:"lte=250"`
	Unit        string  `json:"unit" gorm:"size=32" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       float64 `json:"stock,omitempty"`
	IsActive    bool    `json:"isActive,omitempty"`
}

// Repository representa el repositorio permanente de los productos.
// Se utiliza el concepto de interface para desacoplar la implementación específica del repositorio.
// Los métodos son los básicos de un ABM. Luego, en los usecases, quizás aparezcan otros métodos que se agregan y "extienden" esta interface.
// En general, los métodos devuelven los datos del producto afectado, en los casos de alta, consulta y actualización exitosos. Y un error en caso de falla.
type Repository interface {
	Create(product *Product) (*Product, error) // Create permite agregar un producto nuevo al repositorio
	GetByID(id string) (*Product, error)       // GetByID permite recuperar un único producto, si existe, del repositorio
	GetByName(name string) (*Product, error)   // GetByName permite recuperar un único producto por nombre
	GetAll() ([]*Product, error)               // GetAll permite recuperar, en un slice, todos los productos existentes en el repositorio
	Update(product *Product) (*Product, error) // Update permite actualizar los datos de un producto
	Delete(product *Product) error             // Delete elmimina un producto del repositorio
}
