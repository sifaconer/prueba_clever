package entidades

type Beer struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type BeerBox struct {
	Price float64 `json:"price_total"`
}

type Response struct {
	Message string `json:"message"`
}
