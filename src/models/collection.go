package models

import "fmt"

// Coleccion representa una lista de lectura
type Coleccion struct {
	nombre string
	libros []Libro
}

func NuevaColeccion(nombre string) (Coleccion, error) {
	if nombre == "" {
		return Coleccion{}, fmt.Errorf("error: el nombre no puede estar vacío")
	}
	return Coleccion{nombre: nombre, libros: []Libro{}}, nil
}

func (c Coleccion) Nombre() string { return c.nombre }

func (c Coleccion) Libros() []Libro {
	copia := make([]Libro, len(c.libros))
	copy(copia, c.libros)
	return copia
}

// AgregarLibro inmutable: crea una nueva lista sin modificar la original
func (c Coleccion) AgregarLibro(l Libro) Coleccion {
	nuevaLista := make([]Libro, len(c.libros), len(c.libros)+1)
	copy(nuevaLista, c.libros)
	nuevaLista = append(nuevaLista, l)

	return Coleccion{nombre: c.nombre, libros: nuevaLista}
}
