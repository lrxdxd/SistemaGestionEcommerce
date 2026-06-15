package main

import (
	"fmt"
)

type Usuario struct {
	id     int
	nombre string
	email  string
}

func (u *Usuario) GetId() int {
	return u.id
}

func (u *Usuario) SetId(id int) {
	u.id = id
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) SetNombre(nombre string) {
	u.nombre = nombre
}

func (u *Usuario) GetEmail() string {
	return u.email
}

func (u *Usuario) SetEmail(email string) {
	u.email = email
}

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

type Venta struct {
	id       int
	usuario  *Usuario
	producto *Producto
	cantidad int
	total    float64
}

type SistemaEcommerce struct {
	usuarios       []*Usuario
	productos      []*Producto
	ventas         []*Venta
	lastIdUsuario  int
	lastIdProducto int
	lastIdVenta    int
}

func (s *SistemaEcommerce) AgregarUsuario(nombre, email string) {

	s.lastIdUsuario++

	usuario := &Usuario{}
	usuario.SetId(s.lastIdUsuario)
	usuario.SetNombre(nombre)
	usuario.SetEmail(email)

	s.usuarios = append(s.usuarios, usuario)

	fmt.Println("Usuario agregado correctamente")
}

func (s *SistemaEcommerce) AgregarProducto(nombre string, precio float64, stock int) {

	s.lastIdProducto++

	producto := &Producto{}
	producto.SetId(s.lastIdProducto)
	producto.SetNombre(nombre)
	producto.SetPrecio(precio)
	producto.SetStock(stock)

	s.productos = append(s.productos, producto)

	fmt.Println("Producto agregado correctamente")
}

func (s *SistemaEcommerce) ListarUsuarios() {

	fmt.Println("Cantidad de usuarios:", len(s.usuarios))
	fmt.Println("\n===== USUARIOS =====")

	for _, usuario := range s.usuarios {
		fmt.Printf(
			"ID:%d | Nombre:%s | Email:%s\n",
			usuario.GetId(),
			usuario.GetNombre(),
			usuario.GetEmail(),
		)
	}
}

func (s *SistemaEcommerce) ListarProductos() {

	fmt.Println("\n===== PRODUCTOS =====")

	for _, producto := range s.productos {

		fmt.Printf(
			"ID:%d | Nombre:%s | Precio:%.2f | Stock:%d\n",
			producto.GetId(),
			producto.GetNombre(),
			producto.GetPrecio(),
			producto.GetStock(),
		)
	}
}
func (s *SistemaEcommerce) RegistrarVenta(usuarioId int, productoId int, cantidad int) {

	var usuario *Usuario
	var producto *Producto

	for _, u := range s.usuarios {
		if u.GetId() == usuarioId {
			usuario = u
			break
		}
	}

	for _, p := range s.productos {
		if p.GetId() == productoId {
			producto = p
			break
		}
	}

	if usuario == nil {
		fmt.Println("Usuario no encontrado")
		return
	}

	if producto == nil {
		fmt.Println("Producto no encontrado")
		return
	}

	if producto.GetStock() < cantidad {
		fmt.Println("Stock insuficiente")
		return
	}

	s.lastIdVenta++

	venta := &Venta{
		id:       s.lastIdVenta,
		usuario:  usuario,
		producto: producto,
		cantidad: cantidad,
		total:    float64(cantidad) * producto.GetPrecio(),
	}

	producto.SetStock(producto.GetStock() - cantidad)

	s.ventas = append(s.ventas, venta)

	fmt.Println("Venta registrada correctamente")
}

func (s *SistemaEcommerce) ListarVentas() {

	fmt.Println("\n===== VENTAS =====")

	for _, venta := range s.ventas {

		fmt.Printf(
			"ID:%d | Cliente:%s | Producto:%s | Cantidad:%d | Total: %.2f\n",
			venta.id,
			venta.usuario.GetNombre(),
			venta.producto.GetNombre(),
			venta.cantidad,
			venta.total,
		)
	}
}
func (s *SistemaEcommerce) ReporteGeneral() {

	var ingresos float64

	for _, venta := range s.ventas {
		ingresos += venta.total
	}

	fmt.Println("\n===== REPORTE GENERAL =====")
	fmt.Println("Usuarios:", len(s.usuarios))
	fmt.Println("Productos:", len(s.productos))
	fmt.Println("Ventas:", len(s.ventas))
	fmt.Printf("Ingresos Totales: %.2f\n", ingresos)
}

func main() {

	sistema := &SistemaEcommerce{}

	for {

		fmt.Println("\n==============================")
		fmt.Println(" SISTEMA DE GESTION ECOMMERCE ")
		fmt.Println("==============================")
		fmt.Println("1. Registrar Usuario")
		fmt.Println("2. Registrar Producto")
		fmt.Println("3. Listar Usuarios")
		fmt.Println("4. Listar Productos")
		fmt.Println("5. Registrar Venta")
		fmt.Println("6. Listar Ventas")
		fmt.Println("7. Reporte General")
		fmt.Println("8. Salir")

		var opcion int

		fmt.Print("Seleccione una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {

		case 1:

			var nombre string
			var email string

			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)

			fmt.Print("Email: ")
			fmt.Scanln(&email)

			sistema.AgregarUsuario(nombre, email)

		case 2:

			var nombre string
			var precio float64
			var stock int

			fmt.Print("Nombre producto: ")
			fmt.Scanln(&nombre)

			fmt.Print("Precio: ")
			fmt.Scanln(&precio)

			fmt.Print("Stock: ")
			fmt.Scanln(&stock)

			sistema.AgregarProducto(nombre, precio, stock)

		case 3:

			sistema.ListarUsuarios()

		case 4:

			sistema.ListarProductos()

		case 5:

			var usuarioId int
			var productoId int
			var cantidad int

			fmt.Print("ID Usuario: ")
			fmt.Scanln(&usuarioId)

			fmt.Print("ID Producto: ")
			fmt.Scanln(&productoId)

			fmt.Print("Cantidad: ")
			fmt.Scanln(&cantidad)

			sistema.RegistrarVenta(usuarioId, productoId, cantidad)

		case 6:

			sistema.ListarVentas()

		case 7:

			sistema.ReporteGeneral()

		case 8:

			fmt.Println("Saliendo...")
			return
		}
	}
}
