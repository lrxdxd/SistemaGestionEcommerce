package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"proyecto_poo/db"
)

func APIUsuariosHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT id,nombre,email FROM usuarios")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Usuario struct {
		ID     int
		Nombre string
		Email  string
	}

	var usuarios []Usuario

	for rows.Next() {
		var u Usuario

		rows.Scan(
			&u.ID,
			&u.Nombre,
			&u.Email,
		)

		usuarios = append(usuarios, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func APIProductosHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT id,nombre,precio,stock FROM productos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Producto struct {
		ID     int
		Nombre string
		Precio float64
		Stock  int
	}

	var productos []Producto

	for rows.Next() {

		var p Producto

		rows.Scan(
			&p.ID,
			&p.Nombre,
			&p.Precio,
			&p.Stock,
		)

		productos = append(productos, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}

func APIVentasHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(`
		SELECT
			v.id,
			u.nombre,
			p.nombre,
			v.cantidad,
			v.total
		FROM ventas v
		INNER JOIN usuarios u ON v.usuario_id=u.id
		INNER JOIN productos p ON v.producto_id=p.id
		ORDER BY v.id
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	type Venta struct {
		ID       int
		Usuario  string
		Producto string
		Cantidad int
		Total    float64
	}

	var ventas []Venta

	for rows.Next() {

		var v Venta

		rows.Scan(
			&v.ID,
			&v.Usuario,
			&v.Producto,
			&v.Cantidad,
			&v.Total,
		)

		ventas = append(ventas, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ventas)
}

func APIDashboardHandler(w http.ResponseWriter, r *http.Request) {

	var dashboard Dashboard

	db.DB.QueryRow("SELECT COUNT(*) FROM usuarios").Scan(&dashboard.Usuarios)
	db.DB.QueryRow("SELECT COUNT(*) FROM productos").Scan(&dashboard.Productos)
	db.DB.QueryRow("SELECT COUNT(*) FROM ventas").Scan(&dashboard.Ventas)
	db.DB.QueryRow("SELECT IFNULL(SUM(total),0) FROM ventas").Scan(&dashboard.Ingresos)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)
}

func APIUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	type Usuario struct {
		ID     int
		Nombre string
		Email  string
	}

	var u Usuario

	err := db.DB.QueryRow(
		"SELECT id,nombre,email FROM usuarios WHERE id=?",
		params["id"],
	).Scan(
		&u.ID,
		&u.Nombre,
		&u.Email,
	)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func APIProductoHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	type Producto struct {
		ID     int
		Nombre string
		Precio float64
		Stock  int
	}

	var p Producto

	err := db.DB.QueryRow(
		"SELECT id,nombre,precio,stock FROM productos WHERE id=?",
		params["id"],
	).Scan(
		&p.ID,
		&p.Nombre,
		&p.Precio,
		&p.Stock,
	)

	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func APIVentaHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	type Venta struct {
		ID       int
		Usuario  string
		Producto string
		Cantidad int
		Total    float64
	}

	var v Venta

	err := db.DB.QueryRow(`
		SELECT
			v.id,
			u.nombre,
			p.nombre,
			v.cantidad,
			v.total
		FROM ventas v
		INNER JOIN usuarios u ON v.usuario_id=u.id
		INNER JOIN productos p ON v.producto_id=p.id
		WHERE v.id=?
	`, params["id"]).Scan(
		&v.ID,
		&v.Usuario,
		&v.Producto,
		&v.Cantidad,
		&v.Total,
	)

	if err != nil {
		http.Error(w, "Venta no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func APIStockHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(
		"SELECT nombre,stock FROM productos",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	type Stock struct {
		Producto string
		Stock    int
	}

	var lista []Stock

	for rows.Next() {

		var s Stock

		rows.Scan(
			&s.Producto,
			&s.Stock,
		)

		lista = append(lista, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}
