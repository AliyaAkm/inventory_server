package handlers

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	"inventory-service/proto/inventorypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InventoryHandler struct {
	inventorypb.UnimplementedInventoryServiceServer
	categoryUC *usecase.CategoryUsecase
	productUC  *usecase.ProductUsecase
}

func NewInventoryHandler(catUC *usecase.CategoryUsecase, prodUC *usecase.ProductUsecase) *InventoryHandler {
	return &InventoryHandler{
		categoryUC: catUC,
		productUC:  prodUC,
	}
}

func (h *InventoryHandler) CreateCategory(ctx context.Context, req *inventorypb.CreateCategoryRequest) (*inventorypb.CategoryResponse, error) {
	category := &domain.Category{Name: req.GetName()}
	if err := h.categoryUC.Create(category); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create category: %v", err)
	}

	return &inventorypb.CategoryResponse{
		Id:   uint32(category.ID),
		Name: category.Name,
	}, nil
}

func (h *InventoryHandler) GetCategoryByID(ctx context.Context, req *inventorypb.GetCategoryRequest) (*inventorypb.CategoryResponse, error) {
	category, err := h.categoryUC.GetByID(uint(req.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "category not found: %v", err)
	}

	return &inventorypb.CategoryResponse{
		Id:   uint32(category.ID),
		Name: category.Name,
	}, nil
}

func (h *InventoryHandler) UpdateCategory(ctx context.Context, req *inventorypb.UpdateCategoryRequest) (*inventorypb.CategoryResponse, error) {
	category := &domain.Category{Name: req.GetName()}
	if err := h.categoryUC.Update(uint(req.GetId()), category); err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot update category: %v", err)
	}

	return &inventorypb.CategoryResponse{
		Id:   req.GetId(),
		Name: category.Name,
	}, nil
}

func (h *InventoryHandler) DeleteCategory(ctx context.Context, req *inventorypb.DeleteCategoryRequest) (*inventorypb.DeleteResponse, error) {
	if err := h.categoryUC.Delete(uint(req.GetId())); err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot delete category: %v", err)
	}

	return &inventorypb.DeleteResponse{
		Message: "Category deleted",
	}, nil
}

func (h *InventoryHandler) ListCategories(ctx context.Context, req *inventorypb.ListCategoriesRequest) (*inventorypb.CategoryListResponse, error) {
	categories, err := h.categoryUC.List(req.GetName(), int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list categories: %v", err)
	}

	var resp []*inventorypb.CategoryResponse
	for _, c := range categories {
		resp = append(resp, &inventorypb.CategoryResponse{
			Id:   uint32(c.ID),
			Name: c.Name,
		})
	}

	return &inventorypb.CategoryListResponse{Categories: resp}, nil
}

func (h *InventoryHandler) CreateProduct(ctx context.Context, req *inventorypb.CreateProductRequest) (*inventorypb.ProductResponse, error) {
	product := &domain.Product{
		Name:       req.GetName(),
		Price:      req.GetPrice(),
		Stock:      int(req.GetStock()),
		CategoryID: uint(req.GetCategoryId()),
	}

	if err := h.productUC.Create(product); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	return &inventorypb.ProductResponse{
		Id:         uint32(product.ID),
		Name:       product.Name,
		Price:      product.Price,
		Stock:      int32(product.Stock),
		CategoryId: uint32(product.CategoryID),
	}, nil
}

func (h *InventoryHandler) GetProductByID(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.ProductResponse, error) {
	product, err := h.productUC.GetByID(uint(req.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	return &inventorypb.ProductResponse{
		Id:         uint32(product.ID),
		Name:       product.Name,
		Price:      product.Price,
		Stock:      int32(product.Stock),
		CategoryId: uint32(product.CategoryID),
	}, nil
}

func (h *InventoryHandler) UpdateProduct(ctx context.Context, req *inventorypb.UpdateProductRequest) (*inventorypb.ProductResponse, error) {
	product := &domain.Product{
		Name:       req.GetName(),
		Price:      req.GetPrice(),
		Stock:      int(req.GetStock()),
		CategoryID: uint(req.GetCategoryId()),
	}

	if err := h.productUC.Update(uint(req.GetId()), product); err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to update product: %v", err)
	}

	return &inventorypb.ProductResponse{
		Id:         req.GetId(),
		Name:       product.Name,
		Price:      product.Price,
		Stock:      int32(product.Stock),
		CategoryId: req.GetCategoryId(),
	}, nil
}

func (h *InventoryHandler) DeleteProduct(ctx context.Context, req *inventorypb.DeleteProductRequest) (*inventorypb.DeleteResponse, error) {
	if err := h.productUC.Delete(uint(req.GetId())); err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot delete product: %v", err)
	}

	return &inventorypb.DeleteResponse{
		Message: "Product deleted",
	}, nil
}

func (h *InventoryHandler) ListProducts(ctx context.Context, req *inventorypb.ListProductsRequest) (*inventorypb.ProductListResponse, error) {
	var categoryID *uint
	if req.CategoryId != 0 {
		tmp := uint(req.GetCategoryId())
		categoryID = &tmp
	}

	products, err := h.productUC.List(req.GetName(), categoryID, int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}

	var resp []*inventorypb.ProductResponse
	for _, p := range products {
		resp = append(resp, &inventorypb.ProductResponse{
			Id:         uint32(p.ID),
			Name:       p.Name,
			Price:      p.Price,
			Stock:      int32(p.Stock),
			CategoryId: uint32(p.CategoryID),
		})
	}

	return &inventorypb.ProductListResponse{Products: resp}, nil
}
