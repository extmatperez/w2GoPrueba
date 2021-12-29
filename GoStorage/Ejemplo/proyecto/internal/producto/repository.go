package internal

import (
	"database/sql"

	internal "github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/internal/categoria"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/internal/models"
)

type RepositoryProducto interface {
	GetOne(id int) (models.Producto, error)
	SumaPorCategoria() ([]models.SumaPorCategoria, error)
}

type repositoryProducto struct {
	db *sql.DB
}

func NewRepositoryProducto(baseDeDatos *sql.DB) RepositoryProducto {
	return &repositoryProducto{db: baseDeDatos}
}

func (repo *repositoryProducto) GetOne(id int) (models.Producto, error) {
	repoCat := internal.NewRepositoryCategoria(repo.db)
	// query := "select p.id, p.nombre, p.precio, c.id, c.nombre from producto p inner join categoria c on p.idcategoria = c.id WHERE p.id = ?"
	query := "SELECT id, nombre, precio, idcategoria FROM Producto WHERE id = ?"

	var ProductoLeido models.Producto

	rows, err := repo.db.Query(query, id)

	if err != nil {
		return ProductoLeido, err
	}

	for rows.Next() {
		err = rows.Scan(&ProductoLeido.ID, &ProductoLeido.Nombre, &ProductoLeido.Precio, &ProductoLeido.Categoria.ID)
		// err = rows.Scan(&ProductoLeido.ID, &ProductoLeido.Nombre, &ProductoLeido.Precio, &ProductoLeido.Categoria.ID, &ProductoLeido.Categoria.Nombre)
		if err != nil {
			return ProductoLeido, err
		}
		ProductoLeido.Categoria, err = repoCat.GetOne(ProductoLeido.Categoria.ID)
		if err != nil {
			return ProductoLeido, err
		}
	}
	return ProductoLeido, nil
}

func (repo *repositoryProducto) SumaPorCategoria() ([]models.SumaPorCategoria, error) {
	var tablaSumasPorCategorias []models.SumaPorCategoria
	// query := "select p.id, p.nombre, p.precio, c.id, c.nombre from producto p inner join categoria c on p.idcategoria = c.id WHERE p.id = ?"
	query := "select c.nombre, sum(p.precio) suma from producto p inner join categoria c on p.idcategoria = c.id group by c.nombre order by suma"

	rows, err := repo.db.Query(query)

	if err != nil {
		return tablaSumasPorCategorias, err
	}

	for rows.Next() {
		var sumaAgregar models.SumaPorCategoria
		err = rows.Scan(&sumaAgregar.NombreCategoria, &sumaAgregar.Suma)
		// err = rows.Scan(&ProductoLeido.ID, &ProductoLeido.Nombre, &ProductoLeido.Precio, &ProductoLeido.Categoria.ID, &ProductoLeido.Categoria.Nombre)
		if err != nil {
			return tablaSumasPorCategorias, err
		}

		tablaSumasPorCategorias = append(tablaSumasPorCategorias, sumaAgregar)
	}

	return tablaSumasPorCategorias, nil
}
