package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, precio float64) (Producto, error)
	Update(id int, nombre string, precio float64) (Producto, error)
	Delete(id int) error
	Average() (float64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll() ([]Producto, error) {
	Productos, err := ser.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return Productos, nil
}

func (ser *service) Store(nombre string, precio float64) (Producto, error) {
	ultimoId, err := ser.repository.LastId()

	if err != nil {
		return Producto{}, err
	}
	per, err := ser.repository.Store(ultimoId+1, nombre, precio)

	if err != nil {
		return Producto{}, err
	}

	return per, nil
}

func (ser *service) Update(id int, nombre string, precio float64) (Producto, error) {
	return ser.repository.Update(id, nombre, precio)
}

func (ser *service) Delete(id int) error {
	return ser.repository.Delete(id)
}
func (ser *service) Average() (float64, error) {
	return ser.repository.Average()
}
