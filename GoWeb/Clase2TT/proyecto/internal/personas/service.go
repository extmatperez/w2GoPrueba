package internal

type Service interface {
	GetAll() ([]Persona, error)
	Store(nombre, apellido string, edad int) (Persona, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Persona, error) {
	personas, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return personas, nil
}

func (ser *service) Store(nombre, apellido string, edad int) (Persona, error) {
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Persona{}, err
	}

	per, err := ser.repository.Store(ultimoId+1, nombre, apellido, edad)

	if err != nil {
		return Persona{}, err
	}

	return per, nil
}
