package main

import (
	"log"
	"net/http"

	"github.com/manejador_tareas_Golang/handler"
	"github.com/manejador_tareas_Golang/storage"
)

func main() {
	storageIntance := storage.NewMemory()

	mux := http.NewServeMux()

	handler.RouterPerson(mux, storageIntance)
	log.Println("Servidor corriendo en el puesto: 3000")
	http.ListenAndServe(":3000", mux)
	
}