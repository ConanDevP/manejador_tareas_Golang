package model

import "errors"

var (
	
	TaskIsNil = errors.New("La tarea no puede estar vacia")
	IDNotFound = errors.New("ID no encontrado")
	TasksNotFound = errors.New("Tareas no encontradas")

)