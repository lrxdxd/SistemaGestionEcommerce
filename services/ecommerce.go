package main

func (s *SistemaEcommerce) TotalUsuarios() int {
	return len(s.usuarios)
}

func (s *SistemaEcommerce) TotalProductos() int {
	return len(s.productos)
}

func (s *SistemaEcommerce) TotalVentas() int {
	return len(s.ventas)
}
