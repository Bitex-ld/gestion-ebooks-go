package models

import "fmt"

// Libro representa un e-book inmutable
type Libro struct {
	id     string
	titulo string
	autor  string
	genero string
}

// Constructor con validación
func NuevoLibro(id, titulo, autor, genero string) (Libro, error) {
	if id == "" || titulo == "" || autor == "" {
		return Libro{}, fmt.Errorf("error: los campos 'id', 'titulo' y 'autor' son obligatorios")
	}
	return Libro{id: id, titulo: titulo, autor: autor, genero: genero}, nil
}

// Getters (Encapsulamiento)
func (l Libro) ID() string     { return l.id }
func (l Libro) Titulo() string { return l.titulo }
func (l Libro) Autor() string  { return l.autor }
func (l Libro) Genero() string { return l.genero }
