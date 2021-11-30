package internal

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var personas []Persona
var lastID int

type Repository interface {
	GetAll() ([]Persona, error)
	Store(id int, nombre string, apellido string, edad int) (Persona, error)
	//Store2(nuevaPersona Persona)(Persona, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Persona, error) {
	return personas, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, edad int) (Persona, error) {
	per := Persona{id, nombre, apellido, edad}
	lastID = id
	personas = append(personas, per)
	return per, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
