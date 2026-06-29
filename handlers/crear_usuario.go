package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"proyecto_poo/db"
)

func CrearUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nombre := r.FormValue("nombre")
		email := r.FormValue("email")

		_, err := db.DB.Exec(
			"INSERT INTO usuarios(nombre,email) VALUES(?,?)",
			nombre,
			email,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Usuario registrado:", nombre)

		http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/crear_usuario.html",
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	tmpl.ExecuteTemplate(w, "base", nil)
}
