package storage

import (
	"fmt"

	"github.com/manejador_tareas_Golang/model"
)

type Memory struct {
	CurrensID uint
	Tasks     map[uint]model.Tasks
}


func NewMemory()*Memory{
	task := make(map[uint]model.Tasks)

	return &Memory{0,task}
}


func (m *Memory)Create(task *model.Tasks)error{
	if task == nil{
		return fmt.Errorf("Error: %v\n", model.TaskIsNil)
	}

	m.CurrensID ++
	m.Tasks[m.CurrensID] = *task
	return nil
}

func (m *Memory)Update(ID uint, task *model.Tasks)error{
	if task == nil{
		return fmt.Errorf("Error: %v\n", model.TaskIsNil)
	}

	if _,ok := m.Tasks[ID]; !ok{
		return fmt.Errorf("Error: %v\n", model.IDNotFound)
	}

	m.Tasks[ID] = *task
	return nil
}

func (m *Memory)Delete(ID uint)error{
	_, ok := m.Tasks[ID]

	if !ok {
		return fmt.Errorf("Error: %v\n", model.IDNotFound)
	}

	delete(m.Tasks, ID)
	return nil
}

func (m *Memory)GetByID(ID uint)(*model.Tasks, error){
	task, ok := m.Tasks[ID]

	if !ok {
		return &model.Tasks{}, fmt.Errorf("Error: %v\n", model.IDNotFound)
	}

	return &task, nil
}

func (m *Memory)GetAll()([]model.Tasks, error){
	if len(m.Tasks)< 1{
		return nil,model.TasksNotFound
	}
	var tasks []model.Tasks

	for _, task:=range m.Tasks{
		tasks = append(tasks, task)
	}

	return tasks, nil
}

