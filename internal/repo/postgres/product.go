package postgres

import (
	"gorm.io/gorm"
	"inventory-service/internal/domain"
	"errors"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(p *domain.Product) error {
	return r.db.Create(p).Error
}

func (r *ProductRepo) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepo) Update(product *domain.Product) error {
	var existing domain.Product
	if err := r.db.First(&existing, product.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("продукт с таким ID не найден")
		}
		return err
	}
	return r.db.Save(product).Error
}


func (r *ProductRepo) Delete(id uint) error {
	result := r.db.Delete(&domain.Product{}, id)
	if result.RowsAffected == 0 {
		return errors.New("продукт с таким ID не найден")
	}
	return result.Error
}

func (r *ProductRepo) List(name string, categoryID *uint, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product
	query := r.db

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	err := query.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}
