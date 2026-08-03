package service

import (
	"fmt"
	"strings"

	"gestion-ebooks-go/src/models"
)

type CatalogoServicio struct {
	libros []models.Libro
}

func NuevoCatalogoServicio(libros []models.Libro) *CatalogoServicio {
	return &CatalogoServicio{libros: libros}
}

func (cs *CatalogoServicio) BuscarPorTitulo(titulo string) ([]models.Libro, error) {
	if strings.TrimSpace(titulo) == "" {
		return nil, fmt.Errorf("error: el título de búsqueda no puede estar vacío")
	}

	var resultados []models.Libro
	for _, b := range cs.libros {
		if strings.Contains(strings.ToLower(b.Titulo()), strings.ToLower(titulo)) {
			resultados = append(resultados, b)
		}
	}

	if len(resultados) == 0 {
		return nil, fmt.Errorf("no se encontraron libros con el título: '%s'", titulo)
	}

	return resultados, nil
}

func (cs *CatalogoServicio) FiltrarPorGenero(genero string) ([]models.Libro, error) {
	if strings.TrimSpace(genero) == "" {
		return nil, fmt.Errorf("error: el género no es válido")
	}

	var resultados []models.Libro
	for _, b := range cs.libros {
		if strings.EqualFold(b.Genero(), genero) {
			resultados = append(resultados, b)
		}
	}

	if len(resultados) == 0 {
		return nil, fmt.Errorf("no se encontraron libros del género: '%s'", genero)
	}

	return resultados, nil
}
