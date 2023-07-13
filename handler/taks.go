package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/manejador_tareas_Golang/model"
)

type task struct {
	storage Storage
}

func NewTask(storage Storage)  task{
	return task{storage}
}


func (p *task)create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "Método no valido"}`))
		return 
	}

	data := model.Tasks{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type":"error", "message": "Estructura no valida"}`))
		return 

	}

	err = p.storage.Create(&data)

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"type-masage": "error", "message": "No se pudo guardar la tareas"}`))
		return 

	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"type-message": "ok", "message": "Todo bien"}`))

}

func (t *task)delete(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "Método no valido"}`))
		return
	}

	ID:= r.URL.Query().Get("id")

	intID, err := strconv.Atoi(ID)

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "La consulta salio mal"}`))
		return
	}

	err = t.storage.Delete(uint(intID))

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"messaje-type":"error", "messaje" : "Algo salio mal y no se elimino la tarea"}`))
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message-type": "ok", "message": "Tarea eliminada"}`))
}

func (t *task)update(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "El método no es valido"}`))
		return
	}

	ID := r.URL.Query().Get("id")

	IDInteger, err:= strconv.Atoi(ID)

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "Algo esta mal con la consulta"}`))
		return

	}

	data := model.Tasks{}

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "Estructura enviada en el body, esta mal estructurada"}`))
		return 

	}

	err = t.storage.Update(uint(IDInteger), &data)

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message-type": "error", "message": "Algo salio mal en el servidor"}`))
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message-type": "ok", "message": "Se actualizo de manera correcta"}`))
}


func (t *task)getAll(w http.ResponseWriter, r * http.Request){
	if r.Method != http.MethodGet{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "Método no permitido"}`))
		return
	}

	tasksList, err := t.storage.GetAll()

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message-type": "error", "message": "El servidor fallo"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tasksList)

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message-type": "error", "message": "El servidor falló al codificar los datos en formato JSON"}`))
		return
	}
	
}

func ( t *task)getByID(w http.ResponseWriter,r *http.Request ){
	 if r.Method != http.MethodGet{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "Método no permitido"}`))
		return
	 }

	 ID := r.URL.Query().Get("id")
	 IDInteger,err := strconv.Atoi(ID)

	 if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message-type": "error", "message": "Algo salio mal con la consulta"}`))
		return
	 }

	 task, err := t.storage.GetByID(uint(IDInteger))

	 if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message-type": "error", "message": "Algo salio mal con el servidor"}`))
		return
	 }
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
	 err = json.NewEncoder(w).Encode(task)

	 if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message-type": "error", "message": "El servidor falló al codificar los datos en formato JSON"}`))
		return
	 }
}