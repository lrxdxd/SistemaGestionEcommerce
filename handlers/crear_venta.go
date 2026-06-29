package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"proyecto_poo/db"
	"proyecto_poo/models"
)

func CrearVentaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		usuarioID, _ := strconv.Atoi(r.FormValue("usuario"))
		productoID, _ := strconv.Atoi(r.FormValue("producto"))
		cantidad, _ := strconv.Atoi(r.FormValue("cantidad"))

		var precio float64
		var stock int

		err := db.DB.QueryRow(
			"SELECT precio, stock FROM productos WHERE id=?",
			productoID,
		).Scan(&precio, &stock)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if stock < cantidad {
			http.Error(w, "Stock insuficiente", http.StatusBadRequest)
			return
		}

		total := precio * float64(cantidad)

		_, err = db.DB.Exec(
			"INSERT INTO ventas(usuario_id, producto_id, cantidad, total) VALUES(?,?,?,?)",
			usuarioID,
			productoID,
			cantidad,
			total,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = db.DB.Exec(
			"UPDATE productos SET stock = stock - ? WHERE id=?",
			cantidad,
			productoID,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/ventas", http.StatusSeeOther)
		return
	}

	rowsUsuarios, _ := db.DB.Query(
		"SELECT id,nombre,email FROM usuarios",
	)

	defer rowsUsuarios.Close()

	var usuarios []models.Usuario

	for rowsUsuarios.Next() {

		var u models.Usuario

		rowsUsuarios.Scan(
			&u.ID,
			&u.Nombre,
			&u.Email,
		)

		usuarios = append(usuarios, u)
	}

	rowsProductos, _ := db.DB.Query(
		"SELECT id,nombre,precio,stock FROM productos",
	)

	defer rowsProductos.Close()

	var productos []ProductoVenta

	for rowsProductos.Next() {

		var p ProductoVenta

		rowsProductos.Scan(
			&p.ID,
			&p.Nombre,
			&p.Precio,
			&p.Stock,
		)

		productos = append(productos, p)
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/crear_venta.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datos := DatosVenta{
		Usuarios:  usuarios,
		Productos: productos,
	}

	tmpl.ExecuteTemplate(w, "base", datos)
}
