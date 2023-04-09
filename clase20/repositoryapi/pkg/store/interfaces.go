package store

import "repositoryapi/internal/domain"

type StoreInterface interface {
	Read(id int) (*domain.Product, error)
	Create(product domain.Product) (*domain.Product, error)
	Exist(codeValue string) bool
}

