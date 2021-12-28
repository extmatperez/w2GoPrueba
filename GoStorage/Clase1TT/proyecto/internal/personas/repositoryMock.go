package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/models"
)

type RepositorySQLMock interface {
	Store(persona models.Persona) (models.Persona, error)
	GetOne(id int) models.Persona
	Update(persona models.Persona) (models.Persona, error)
	GetAll() ([]models.Persona, error)
	Delete(id int) error
	GetFullData() ([]models.Persona, error)

	GetOneWithContext(ctx context.Context, id int) (models.Persona, error)
}

type repositorySQLMock struct {
	db *sql.DB
}

func NewRepositorySQLMock(db *sql.DB) RepositorySQLMock {
	return &repositorySQLMock{db}
}

func (r *repositorySQLMock) Store(persona models.Persona) (models.Persona, error) {

	stmt, err := r.db.Prepare("INSERT INTO personas(nombre, apellido, edad) VALUES( ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad)
	if err != nil {
		return models.Persona{}, err
	}
	idCreado, _ := result.LastInsertId()
	persona.ID = int(idCreado)

	return persona, nil
}

func (r *repositorySQLMock) GetOne(id int) models.Persona {

	var personaLeida models.Persona
	rows, err := r.db.Query("SELECT id, nombre,apellido, edad FROM personas WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return personaLeida
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return personaLeida
		}

	}
	return personaLeida
}
func (r *repositorySQLMock) GetAll() ([]models.Persona, error) {
	var misPersonas []models.Persona

	var personaLeida models.Persona
	rows, err := r.db.Query("SELECT id, nombre, apellido, edad FROM personas")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misPersonas = append(misPersonas, personaLeida)
	}
	return misPersonas, nil
}

func (r *repositorySQLMock) Update(persona models.Persona) (models.Persona, error) {

	stmt, err := r.db.Prepare("UPDATE personas SET nombre = ?, apellido = ?, edad = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(persona.Nombre, persona.Apellido, persona.Edad, persona.ID)
	if err != nil {
		return models.Persona{}, err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return models.Persona{}, errors.New("No se encontro la persona")
	}

	return persona, nil
}

func (r *repositorySQLMock) Delete(id int) error {

	stmt, err := r.db.Prepare("DELETE FROM personas WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	filasActualizadas, _ := result.RowsAffected()
	if filasActualizadas == 0 {
		return errors.New("No se encontro la persona")
	}
	return nil
}

func (r *repositorySQLMock) GetFullData() ([]models.Persona, error) {
	var misPersonas []models.Persona

	var personaLeida models.Persona
	rows, err := r.db.Query("select p.id,p.nombre, p.apellido, p.edad, c.nombre, c.nombrepais from personas p inner join ciudad c on p.idciudad = c.id")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad, &personaLeida.Domicilio.Nombre, &personaLeida.Domicilio.NombrePais)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misPersonas = append(misPersonas, personaLeida)
	}
	return misPersonas, nil
}

func (r *repositorySQLMock) GetOneWithContext(ctx context.Context, id int) (models.Persona, error) {

	var personaLeida models.Persona
	// rows, err := db.QueryContext(ctx, "select sleep(30) from dual")
	rows, err := r.db.QueryContext(ctx, "SELECT id, nombre,apellido, edad FROM personas WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return personaLeida, err
	}

	for rows.Next() {
		err = rows.Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad)
		if err != nil {
			log.Fatal(err)
			return personaLeida, err
		}

	}
	return personaLeida, nil
}
