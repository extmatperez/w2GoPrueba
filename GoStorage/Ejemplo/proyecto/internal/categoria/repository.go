package internal

import (
	"database/sql"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/internal/models"
)

type RepositoryCategoria interface {
	GetOne(id int) (models.Categoria, error)
}

type repositoryCategoria struct {
	db *sql.DB
}

func NewRepositoryCategoria(baseDeDatos *sql.DB) RepositoryCategoria {
	return &repositoryCategoria{db: baseDeDatos}
}

func (repo *repositoryCategoria) GetOne(id int) (models.Categoria, error) {
	query := "SELECT id, nombre FROM categoria WHERE id = ?"

	var categoriaLeida models.Categoria

	rows, err := repo.db.Query(query, id)

	if err != nil {
		return categoriaLeida, err
	}

	for rows.Next() {
		err = rows.Scan(&categoriaLeida.ID, &categoriaLeida.Nombre)
		if err != nil {
			return categoriaLeida, err
		}
	}
	return categoriaLeida, nil
}
