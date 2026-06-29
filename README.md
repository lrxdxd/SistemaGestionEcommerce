# Sistema de Gestión Ecommerce

**Materia:** Programación Orientada a Objetos

**Docente:** Milton Ricardo Palacios Morocho

**Integrante:** Lirio David Villón Pérez

**Fecha:** 28 de Junio de 2026

## Objetivo del programa
El objetivo de este proyecto es desarrollar un sistema de gestión para un ecommerce utilizando el lenguaje de programación Go, aplicando los conocimientos adquiridos durante la asignatura de Programación Orientada a Objetos. El sistema permite administrar usuarios, productos y ventas, además de generar reportes y ofrecer servicios web que facilitan la consulta de la información almacenada en la base de datos.
## Descripción del proyecto
El Sistema de Gestión Ecommerce es una aplicación web que fue desarrollada con Go y MySQL. Su finalidad es facilitar la administración básica de un negocio mediante el registro de usuarios, productos y ventas.

Toda la información se almacena en una base de datos MySQL, permitiendo que los datos permanezcan disponibles incluso después de cerrar la aplicación. Además, el sistema incorpora un dashboard con estadísticas generales, un módulo de reportes y servicios web que entregan información en formato JSON.

## Principales funcionalidades
El sistema desarrollado permite realizar las siguientes operaciones:
* Registro y consulta de usuarios.
* Registro y consulta de productos.
* Registro de ventas.
* Actualización automática del stock al realizar una venta.
* Dashboard con información estadística del sistema.
* Reportes generales del ecommerce.
* Persistencia de la información mediante MySQL.
* Consulta de información mediante servicios web en formato JSON.
## Servicios web implementados
1. Consultar todos los usuarios.
2. Consultar todos los productos.
3. Consultar todas las ventas.
4. Consultar la información del dashboard.
5. Consultar un usuario por su identificador.
6. Consultar un producto por su identificador.
7. Consultar una venta por su identificador.
8. Consultar el stock disponible de los productos.
## Tecnología utilizada
* Go (Golang)
* Gorilla Mux
* MySQL
* MySQL Workbench
* HTML
* JSON
* Git y GitHub
## Estructura general del proyecto

* `db`: conexión con la base de datos.
* `handlers`: lógica del sistema y servicios web.
* `models`: estructuras de datos.
* `templates`: vistas HTML.
* `static`: archivos estáticos.
* `main.go`: punto de inicio de la aplicación.
## Conclusión
Con este proyecto fue posible aplicar los conocimientos adquiridos durante la asignatura, integrando programación orientada a objetos, desarrollo web, manejo de bases de datos y servicios web dentro de una misma aplicación funcional.
