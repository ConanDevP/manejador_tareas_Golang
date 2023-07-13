package handler

import "net/http"

func RouterPerson(mux *http.ServeMux, storage Storage){
	h:= NewTask(storage)

	mux.HandleFunc("/v1/tasks/create",h.create)
	mux.HandleFunc("/v1/tasks/delete",h.delete)
	mux.HandleFunc("/v1/tasks/update",h.update)
	mux.HandleFunc("/v1/tasks/get-all",h.getAll)
	mux.HandleFunc("/v1/tasks/get-by-id",h.getByID)




}