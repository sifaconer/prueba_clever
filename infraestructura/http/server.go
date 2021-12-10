package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"prueba.clever.com/dominio/entidades"
	"prueba.clever.com/interfaces/repository"
)

type Server struct {
	Port       string
	Repository repository.BeerRepository
}

func (s *Server) New() {

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORS())

	e.File("/docs", "infraestructura/http/docs/openapi.json")
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URL = "/docs"
	}))
	e.GET("/beers", s.allBeers)
	e.POST("/beers", s.createBeers)
	e.GET("/beers/:beerID", s.beerByID)
	e.GET("/beers/:beerID/boxprice", s.beerPriceBox)

	e.Logger.Fatal(e.Start(":" + s.Port))
}

func (s *Server) allBeers(c echo.Context) error {
	resp, err := s.Repository.AllBeers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) createBeers(c echo.Context) error {
	body := new(entidades.Beer)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Request invalida")
	}

	err := s.Repository.CreateBeer(*body)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	return c.JSON(http.StatusCreated, entidades.Response{
		Message: "Cerveza creada",
	})
}

func (s *Server) beerByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("beerID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "El parametro {beerID} debe ser un numero entero")
	}
	resp, err := s.Repository.BeerByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) beerPriceBox(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("beerID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "El parametro {beerID} debe ser un numero entero")
	}

	currency := c.QueryParam("currency")
	quantity, err := strconv.Atoi(c.QueryParam("quantity"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "El parametro {beerID} debe ser un numero entero")
	}

	resp, err := s.Repository.BeerPriceBox(id, quantity, currency)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
