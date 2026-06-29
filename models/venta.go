package models

type Venta struct {
	ID         int
	UsuarioID  int
	ProductoID int
	Cantidad   int
	Total      float64
}
