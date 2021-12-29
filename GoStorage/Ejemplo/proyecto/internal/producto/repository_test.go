package internal

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Ejemplo/proyecto/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestGetOneProducto(t *testing.T) {
	db := db.StorageDB
	repo := NewRepositoryProducto(db)

	producto, err := repo.GetOne(1)

	assert.Nil(t, err)
	assert.Equal(t, "Frescos", producto.Categoria.Nombre)
	//fmt.Printf("\n%+v", producto)
}

func TestGetOneProducto_txdb(t *testing.T) {
	// db := db.StorageDB
	db, err := db.InitDb()
	assert.Nil(t, err)
	repo := NewRepositoryProducto(db)

	producto, err := repo.GetOne(1)

	assert.Nil(t, err)
	assert.Equal(t, "Frescos", producto.Categoria.Nombre)
	//fmt.Printf("\n%+v", producto)
}

func TestGetOneProducto_sqlmock(t *testing.T) {
	// db := db.StorageDB
	// db, err := db.InitDb()

	idBuscado := 1

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	rowsProductos := sqlmock.NewRows([]string{"id", "nombre", "precio", "idcategoria"})
	rowsProductos.AddRow(1, "Manteca", 102, 5)
	mock.ExpectQuery("SELECT id, nombre, precio, idcategoria FROM Producto WHERE id = ?").WithArgs(idBuscado).WillReturnRows(rowsProductos)

	rowsCategoria := sqlmock.NewRows([]string{"id", "nombre"})
	rowsCategoria.AddRow(5, "Frescos")
	mock.ExpectQuery("SELECT id, nombre FROM categoria WHERE id = ?").WithArgs(5).WillReturnRows(rowsCategoria)

	repo := NewRepositoryProducto(db)

	producto, err := repo.GetOne(idBuscado)

	assert.Nil(t, err)
	assert.Equal(t, "Frescos", producto.Categoria.Nombre)
	// fmt.Printf("\n%+v", producto)
}

func TestSumaCategoriasProducto(t *testing.T) {
	db := db.StorageDB
	repo := NewRepositoryProducto(db)

	sumasPorCategorias, err := repo.SumaPorCategoria()

	assert.Nil(t, err)
	assert.Equal(t, 105.0, sumasPorCategorias[0].Suma)
	assert.True(t, len(sumasPorCategorias) == 4)
	//fmt.Printf("\n%+v", producto)
}
