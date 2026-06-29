package handlers

import (
	"html/template"
	"net/http"

	"proyecto_poo/db"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var totalUsuarios int
	var totalProductos int
	var totalVentas int
	var ingresos float64

	db.DB.QueryRow("SELECT COUNT(*) FROM usuarios").Scan(&totalUsuarios)
	db.DB.QueryRow("SELECT COUNT(*) FROM productos").Scan(&totalProductos)
	db.DB.QueryRow("SELECT COUNT(*) FROM ventas").Scan(&totalVentas)
	db.DB.QueryRow("SELECT IFNULL(SUM(total),0) FROM ventas").Scan(&ingresos)

	dashboard := Dashboard{
		Usuarios:  totalUsuarios,
		Productos: totalProductos,
		Ventas:    totalVentas,
		Ingresos:  ingresos,
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/inicio.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", dashboard)
}
