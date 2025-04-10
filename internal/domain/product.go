package domain

type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	CategoryID uint    `json:"category_id"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
}

type ProductRepository interface {
	Create(*Product) error
	GetByID(id uint) (*Product, error)
	Update(product *Product) error
	Delete(id uint) error
	List(name string, categoryID *uint, limit, offset int) ([]Product, error)
}
