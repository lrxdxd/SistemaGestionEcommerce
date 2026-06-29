package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"proyecto_poo/db"
)

func CrearProductoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nombre := r.FormValue("nombre")

		precio, _ := strconv.ParseFloat(r.FormValue("precio"), 64)

		stock, _ := strconv.Atoi(r.FormValue("stock"))

		_, err := db.DB.Exec(
			"INSERT INTO productos(nombre,precio,stock) VALUES(?,?,?)",
			nombre,
			precio,
			stock,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/productos", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/crear_producto.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}
