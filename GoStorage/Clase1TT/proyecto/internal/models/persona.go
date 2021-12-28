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

type DTOAvgAge struct {
	Nombre  string  `json:"nombre"`
	Average float64 `json:"average"`
}
