package sqlite

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
	"prueba.clever.com/dominio/casosdeuso"
	"prueba.clever.com/dominio/entidades"
)

type BeersImpl struct {
	DB *sql.DB
}

func (b *BeersImpl) AllBeers() ([]entidades.Beer, error) {
	row, err := b.DB.Query("SELECT id, name, brewery, country, price, currency FROM beer")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	resp := []entidades.Beer{}
	for row.Next() {
		beer := entidades.Beer{}
		row.Scan(
			&beer.ID,
			&beer.Name,
			&beer.Brewery,
			&beer.Country,
			&beer.Price,
			&beer.Currency,
		)
		resp = append(resp, beer)
	}
	if err := row.Err(); err != nil {
		return resp, err
	}
	return resp, nil
}

func (b *BeersImpl) CreateBeer(body entidades.Beer) error {
	sql := "INSERT INTO beer(id, name, brewery, country, price, currency) values (?, ?, ?, ?, ?, ?)"
	stmt, err := b.DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		body.ID,
		body.Name,
		body.Brewery,
		body.Country,
		body.Price,
		body.Currency)
	if typeErr, ok := err.(sqlite3.Error); ok {
		if typeErr.Code == sqlite3.ErrNo(sqlite3.ErrConstraint) {
			return errors.New("el id de la cerveza ya existe")
		}
		return err
	}

	return nil
}

func (b *BeersImpl) BeerByID(id int) (entidades.Beer, error) {
	beer := entidades.Beer{}
	stmt, err := b.DB.Prepare("SELECT id, name, brewery, country, price, currency FROM beer where id = ?")
	if err != nil {
		return beer, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&beer.ID,
		&beer.Name,
		&beer.Brewery,
		&beer.Country,
		&beer.Price,
		&beer.Currency,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return beer, errors.New("el id de la cerveza no existe")
		}
		return beer, nil
	}
	return beer, nil
}

func (b *BeersImpl) BeerPriceBox(id, quantity int, currency string) (entidades.BeerBox, error) {
	box := entidades.BeerBox{}
	beer, err := b.BeerByID(id)
	if err != nil {
		return box, err
	}

	total := beer.Price * float64(quantity)
	result, err := casosdeuso.HomologarMoneda(beer.Currency, currency, total)
	if err != nil {
		return box, err
	}

	box.Price = *result
	return box, err
}
