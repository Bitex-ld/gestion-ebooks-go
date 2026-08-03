package service

import "gestion-ebooks-go/src/models"

// BuscadorCatalogo define la interfaz
type BuscadorCatalogo interface {
	BuscarPorTitulo(titulo string) ([]models.Libro, error)
	FiltrarPorGenero(genero string) ([]models.Libro, error)
}
