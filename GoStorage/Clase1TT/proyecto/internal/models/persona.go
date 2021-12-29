package models

type Ciudad struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	NombrePais string `json:"nombrepais"`
}
type Persona struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Domicilio Ciudad `json:"domicilio"`
}
type PersonaDynamo struct {
	ID       string `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}
