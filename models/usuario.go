package main

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
