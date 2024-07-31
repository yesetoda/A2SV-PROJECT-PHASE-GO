package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type user struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var (
	nextId    = 1
	dataStore = map[int]user{
		0: {
			Name: "yene",
			Id:   0,
		},
	}
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /yene", handleGet)
	mux.HandleFunc("GET /yene/{id}", handleGetById)
	mux.HandleFunc("POST /yene", handleCreate)
	mux.HandleFunc("PUT /yene", handleUpdate)
	mux.HandleFunc("DELETE /yene", handleDelete)
	fmt.Println("serivng and port 9090")
	http.ListenAndServe(":9090", mux)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(dataStore)
}

func handleGetById(w http.ResponseWriter, r *http.Request) {
	str_id := r.PathValue("id") //use the PathValue to get the path in the url: example /yene/{id} here if the url is /yene/1 this means the id = 1
	id, err := strconv.Atoi(str_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("this is the key", str_id)
		w.Write([]byte("invalid key"))
		return
	}
	usr, ufound := dataStore[id]
	if !ufound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not user found by this id"))
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(usr)
	// w.Write([]byte("this is the update function"))

}
func handleCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	new_user := user{
		Id:   nextId,
		Name: name,
	}
	nextId += 1
	dataStore[new_user.Id] = new_user
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("this is the create function"))
}
func handleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	new_name := r.FormValue("name")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid key"))
		return
	}
	usr, ufound := dataStore[id]
	if !ufound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not user found by this id"))
		return
	}
	usr.Name = new_name
	dataStore[id] = usr
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(usr)
	// w.Write([]byte("this is the update function"))
}
func handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	fmt.Println(id, err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid key"))
		return
	}
	_, ufound := dataStore[id]
	if !ufound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`no user with the id: ` + strconv.Itoa(id)))
		return
	}
	delete(dataStore, id)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("user with id: " + strconv.Itoa(id) + " is removed"))

}
