package postgres

import (
	"gorm.io/gorm"
	"inventory-service/internal/domain"
	"errors"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) Create(c *domain.Category) error {
	return r.db.Create(c).Error
}

func (r *CategoryRepo) GetByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepo) Update(category *domain.Category) error {
	var existing domain.Category
	if err := r.db.First(&existing, category.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("категория с таким ID не найдена")
		}
		return err
	}
	return r.db.Save(category).Error
}


func (r *CategoryRepo) Delete(id uint) error {
	result := r.db.Delete(&domain.Category{}, id)
	if result.RowsAffected == 0 {
		return errors.New("категория с таким ID не найдена")
	}
	return result.Error
}

func (r *CategoryRepo) List(name string, limit, offset int) ([]domain.Category, error) {
	var categories []domain.Category
	query := r.db

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Limit(limit).Offset(offset).Find(&categories).Error
	return categories, err
}
