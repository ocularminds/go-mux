package mux

import (
	"database/sql"
	"errors"
)

func (p *Product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) getProducts(db *sql.DB) ([]Product, error) {
	return nil, errors.New("Not implemented")
}

func (p *Product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}
