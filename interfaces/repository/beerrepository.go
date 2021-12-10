package repository

import "prueba.clever.com/dominio/entidades"

type BeerRepository interface {
	AllBeers() ([]entidades.Beer, error)
	CreateBeer(entidades.Beer) error
	BeerByID(id int) (entidades.Beer, error)
	BeerPriceBox(id, quantity int, currency string) (entidades.BeerBox, error)
}
