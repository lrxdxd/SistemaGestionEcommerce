package handlers

import (
	"html/template"
	"net/http"

	"proyecto_poo/db"
	"proyecto_poo/models"
)

func UsuariosHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(
		"SELECT id,nombre,email FROM usuarios",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {

		var u models.Usuario

		err := rows.Scan(
			&u.ID,
			&u.Nombre,
			&u.Email,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		usuarios = append(usuarios, u)
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/usuarios.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", usuarios)
}
