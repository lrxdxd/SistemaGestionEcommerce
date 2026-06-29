package handlers

import "proyecto_poo/models"

type DatosVenta struct {
	Usuarios  []models.Usuario
	Productos []ProductoVenta
}
type ProductoVenta struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}
