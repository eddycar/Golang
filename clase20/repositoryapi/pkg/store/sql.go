package store

import (
	"database/sql"
	"repositoryapi/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Read(id int) (*domain.Product, error) {
	var productReturn domain.Product

	query := "SELECT * FROM products WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&productReturn.ID, &productReturn.Name, &productReturn.Quantity, &productReturn.CodeValue, &productReturn.IsPublished, &productReturn.Expiration, &productReturn.Price)
	if err != nil {
		return nil, err
	}
	return &productReturn, nil
}


func (s *SqlStore) Create(product domain.Product) (*domain.Product, error){
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lid, _ := res.LastInsertId()
	product.ID = int(lid)
	return &product, nil
}

func (s *SqlStore) Exist(codeValue string) bool{
	var exist bool
	var id int

	query := "SELECT id FROM products WHERE code_value = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {
		return exist
	}

	if id > 0 {
		exist = true
	}

	return exist

}

