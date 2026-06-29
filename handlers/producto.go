package handlers

import (
	"html/template"
	"net/http"

	"proyecto_poo/db"
	"proyecto_poo/models"
)

func ProductosHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(
		"SELECT id,nombre,precio,stock FROM productos",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var productos []models.Producto

	for rows.Next() {

		var p models.Producto

		err := rows.Scan(
			&p.ID,
			&p.Nombre,
			&p.Precio,
			&p.Stock,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		productos = append(productos, p)
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/productos.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", productos)
}
