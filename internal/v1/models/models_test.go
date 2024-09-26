package models_test

import (
	"encoding/json"
	"testing"

	"seguridad-cicd/internal/v1/models"

	"github.com/stretchr/testify/assert"
)

func TestAlumno(t *testing.T) {
	alumno := models.Alumno{
		Nombre:   "John",
		Apellido: "Doe",
	}

	// Convert alumno to JSON
	jsonData, err := json.Marshal(alumno)
	assert.NoError(t, err)

	// Convert JSON back to alumno
	var newAlumno models.Alumno
	err = json.Unmarshal(jsonData, &newAlumno)
	assert.NoError(t, err)

	// Check if the values are the same
	assert.Equal(t, alumno.Nombre, newAlumno.Nombre)
	assert.Equal(t, alumno.Apellido, newAlumno.Apellido)
}
