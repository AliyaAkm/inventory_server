package usecase

import (
"inventory-service/internal/domain"
"errors"
)

type ProductUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: r}
}

func (uc *ProductUsecase) Create(p *domain.Product) error {
	return uc.repo.Create(p)
}

func (uc *ProductUsecase) GetByID(id uint) (*domain.Product, error) {
	return uc.repo.GetByID(id)
}

func (u *ProductUsecase) Update(id uint, product *domain.Product) error {
	if id == 0 {
		return errors.New("некорректный ID")
	}

	product.ID = id
	return u.repo.Update(product)
}


func (uc *ProductUsecase) Delete(id uint) error {
	return uc.repo.Delete(id)
}

func (uc *ProductUsecase) List(name string, categoryID *uint, limit, offset int) ([]domain.Product, error) {
	return uc.repo.List(name, categoryID, limit, offset)
}
