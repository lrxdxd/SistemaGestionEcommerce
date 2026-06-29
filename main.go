package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"proyecto_poo/db"
	"proyecto_poo/handlers"
)

func main() {

	// Conectar a MySQL
	db.Conectar()

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)

	r.HandleFunc("/usuarios", handlers.UsuariosHandler)
	r.HandleFunc("/usuarios/crear", handlers.CrearUsuarioHandler)

	r.HandleFunc("/productos", handlers.ProductosHandler)
	r.HandleFunc("/productos/crear", handlers.CrearProductoHandler)

	r.HandleFunc("/ventas", handlers.VentasHandler)
	r.HandleFunc("/ventas/crear", handlers.CrearVentaHandler)

	r.HandleFunc("/reportes", handlers.ReportesHandler)
	r.HandleFunc("/api/usuarios", handlers.APIUsuariosHandler)
	r.HandleFunc("/api/productos", handlers.APIProductosHandler)
	r.HandleFunc("/api/ventas", handlers.APIVentasHandler)
	r.HandleFunc("/api/dashboard", handlers.APIDashboardHandler)
	r.HandleFunc("/api/usuario/{id}", handlers.APIUsuarioHandler)
	r.HandleFunc("/api/producto/{id}", handlers.APIProductoHandler)
	r.HandleFunc("/api/venta/{id}", handlers.APIVentaHandler)
	r.HandleFunc("/api/stock", handlers.APIStockHandler)

	log.Println("Servidor iniciado en http://localhost:8081")

	err := http.ListenAndServe(":8081", r)

	if err != nil {
		log.Fatal(err)
	}
}
