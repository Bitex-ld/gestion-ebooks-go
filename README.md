# Sistema de Gestion de Libros Electronicos

Este proyecto es la planificacion de un sistema para organizar y buscar libros electronicos (E-books) usando el lenguaje Go.

---

## Como funciona el sistema
El proyecto esta dividido en dos partes principales:

1. Catalogo de libros: Permite buscar libros por su titulo, autor o filtrar por su genero.
2. Coleccion del usuario: Permite al usuario guardar sus libros favoritos en listas de lectura.

Para evitar errores, el sistema se diseño usando programacion funcional. Esto significa que cuando buscamos un libro o agregamos uno a la lista, nunca modificamos los datos originales, sino que el sistema crea copias nuevas con los cambios.

---

## Carpetas del proyecto
Hemos organizado el repositorio de la siguiente manera:

* docs: Aqui guardaremos el documento PDF con la explicacion del proyecto.
* src/domain: Aqui se define como es un "Libro" y una "Coleccion" (sus datos).
* src/catalog: Aqui va la logica para buscar y filtrar los libros.
* src/collection: Aqui va la logica para que el usuario guarde sus libros.

---

## Tecnologias utilizadas
Para este proyecto no necesitamos instalar programas o librerias de otras personas. Usamos unicamente las herramientas que vienen por defecto en Go (como "strings" para buscar textos y "slices" para ordenar listas).
