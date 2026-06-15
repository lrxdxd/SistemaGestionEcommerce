package main

type Producto struct {
	id     int
	nombre string
	precio float64
	stock  int
}

func (p *Producto) GetId() int {
	return p.id
}

func (p *Producto) SetId(id int) {
	p.id = id
}

func (p *Producto) GetNombre() string {
	return p.nombre
}

func (p *Producto) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Producto) GetPrecio() float64 {
	return p.precio
}

func (p *Producto) SetPrecio(precio float64) {
	p.precio = precio
}

func (p *Producto) GetStock() int {
	return p.stock
}

func (p *Producto) SetStock(stock int) {
	p.stock = stock
}
