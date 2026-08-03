package main

import (
	"fmt"
	"log"

	"gestion-ebooks-go/src/models"
	"gestion-ebooks-go/src/service"
)

func main() {
	fmt.Println("=== SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS ===")

	// 1. Datos iniciales con manejo de errores
	l1, err := models.NuevoLibro("1", "El Hobbit", "J.R.R. Tolkien", "Fantasía")
	if err != nil {
		log.Fatal(err)
	}

	l2, err := models.NuevoLibro("2", "1984", "George Orwell", "Ciencia Ficción")
	if err != nil {
		log.Fatal(err)
	}

	l3, err := models.NuevoLibro("3", "Fahrenheit 451", "Ray Bradbury", "Ciencia Ficción")
	if err != nil {
		log.Fatal(err)
	}

	catalogoInicial := []models.Libro{l1, l2, l3}

	// 2. Uso de Interfaz
	var buscador service.BuscadorCatalogo = service.NuevoCatalogoServicio(catalogoInicial)

	fmt.Println("\n--- Buscando libros del género 'Ciencia Ficción' ---")
	librosCF, err := buscador.FiltrarPorGenero("Ciencia Ficción")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, libro := range librosCF {
			fmt.Printf("• %s - %s\n", libro.Titulo(), libro.Autor())
		}
	}

	// 3. Prueba de Manejo de Errores
	fmt.Println("\n--- Probando Manejo de Errores (Búsqueda inexistente) ---")
	_, err = buscador.BuscarPorTitulo("Programación en Go")
	if err != nil {
		fmt.Println("Error capturado adecuadamente:", err)
	}

	// 4. Inmutabilidad
	fmt.Println("\n--- Demostración de Inmutabilidad ---")
	coleccionOriginal, _ := models.NuevaColeccion("Mis Favoritos")
	coleccionActualizada := coleccionOriginal.AgregarLibro(l1)

	fmt.Printf("Cantidad de libros en Colección Original: %d\n", len(coleccionOriginal.Libros()))
	fmt.Printf("Cantidad de libros en Nueva Colección: %d\n", len(coleccionActualizada.Libros()))
}
