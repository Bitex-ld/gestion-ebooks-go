# Sistema de Gestión de Libros Electrónicos (E-Books) en Go

Proyecto de desarrollo en lenguaje Go enfocado en la gestión de catálogos y colecciones de libros electrónicos. La solución aplica principios de Programación Funcional, Inmutabilidad y patrones de la Programación Orientada a Objetos (POO).

---

## Objetivos e Implementación Técnica

El sistema ha sido diseñado bajo los estándares arquitectónicos del lenguaje Go:

* Encapsulamiento y Ocultamiento de Información: Las estructuras principales (Libro y Coleccion) mantienen sus campos privados. El acceso a los datos se realiza estrictamente a través de métodos de lectura (getters) y copias defensivas para garantizar la integridad del estado.
* Inmutabilidad y Enfoque Funcional: Las operaciones de actualización o adición no modifican las estructuras existentes, sino que retornan nuevas instancias con el estado resultante.
* Abstracción mediante Interfaces: La capa de servicios desacopla la búsqueda e interacción del catálogo a través de la interfaz BuscadorCatalogo.
* Manejo Riguroso de Errores: Control explícito de estados excepcionales (búsquedas sin resultados, entradas no válidas) utilizando la firma idiomática de Go (T, error).

---

## Estructura del Proyecto

gestion-ebooks-go/
├── src/
│   ├── models/
│   │   ├── book.go          # Definición de la estructura Libro y constructores
│   │   └── collection.go    # Gestión inmutable de colecciones de libros
│   └── service/
│       ├── interface.go     # Interfaz BuscadorCatalogo
│       └── catalog.go       # Implementación de los servicios de consulta
├── go.mod                   # Definición del módulo Go
├── main.go                  # Punto de entrada y casos de uso
└── README.md                # Documentación del proyecto                