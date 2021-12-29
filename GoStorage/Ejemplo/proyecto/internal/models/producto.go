package models

type Producto struct {
	ID        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Precio    float64   `json:"precio"`
	Categoria Categoria `json:"categoria"`
}
