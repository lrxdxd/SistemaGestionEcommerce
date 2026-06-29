package handlers

import (
	"html/template"
	"net/http"

	"proyecto_poo/db"
)

func VentasHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(`
		SELECT
			v.id,
			u.nombre,
			p.nombre,
			v.cantidad,
			v.total
		FROM ventas v
		INNER JOIN usuarios u ON v.usuario_id = u.id
		INNER JOIN productos p ON v.producto_id = p.id
		ORDER BY v.id
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var ventas []VentaView

	for rows.Next() {

		var v VentaView

		err := rows.Scan(
			&v.ID,
			&v.Usuario,
			&v.Producto,
			&v.Cantidad,
			&v.Total,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ventas = append(ventas, v)
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/ventas.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", ventas)
}
