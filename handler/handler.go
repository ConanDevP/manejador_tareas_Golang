package handler

import "github.com/manejador_tareas_Golang/model"

type Storage interface {
	Create(*model.Tasks)error
	Update(uint, *model.Tasks)error
	Delete(uint)error
	GetByID(uint)(*model.Tasks, error)
	GetAll()([]model.Tasks, error)
}