package main

type Venta struct {
	id       int
	usuario  *Usuario
	producto *Producto
	cantidad int
	total    float64
}
