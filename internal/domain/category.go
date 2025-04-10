package domain

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	Create(*Category) error
	GetByID(id uint) (*Category, error)
	Update(category *Category) error
	Delete(id uint) error
	List(name string, limit, offset int) ([]Category, error)
}
