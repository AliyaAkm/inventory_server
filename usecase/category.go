package usecase

import (
"inventory-service/internal/domain"
"errors"
)

type CategoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(r domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: r}
}

func (uc *CategoryUsecase) Create(c *domain.Category) error {
	return uc.repo.Create(c)
}

func (uc *CategoryUsecase) GetByID(id uint) (*domain.Category, error) {
	return uc.repo.GetByID(id)
}

func (u *CategoryUsecase) Update(id uint, category *domain.Category) error {
	if id == 0 {
		return errors.New("некорректный ID")
	}

	category.ID = id
	return u.repo.Update(category)
}


func (uc *CategoryUsecase) Delete(id uint) error {
	return uc.repo.Delete(id)
}

func (uc *CategoryUsecase) List(name string, limit, offset int) ([]domain.Category, error) {
	return uc.repo.List(name, limit, offset)
}
