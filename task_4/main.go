package main

// import (
// 	"net/http"
// 	"regexp"
// 	"sync"
// )

// var (
// 	listUsersRe   = regexp.MustCompile(`^\/users[\/]*$`)
// 	getUsersRe    = regexp.MustCompile(`^\/users\/(\d+)$`)
// 	createUsersRe = regexp.MustCompile(`^\/users[\/]*$`)
// )

// type user struct {
// 	Id   string `json:"id"`
// 	Name string `json:"name"`
// }
// type dataStore struct {
// 	m map[string]user
// 	*sync.RWMutex
// }
// type userHandler struct {
// 	store *dataStore
// }

// func (h *userHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	switch {
// 	case r.Method == http.MethodGet && listUsersRe.MatchString(r.URL.Path):
// 		h.List(w, r)
// 		return
// 	case r.Method == http.MethodGet && getUsersRe.MatchString(r.URL.Path):
// 		h.Get(w, r)
// 		return
// 	case r.Method == http.MethodPost && createUsersRe.MatchString(r.URL.Path):
// 		h.Create(w, r)
// 		return
// 	default:
// 		notFound(w, r)
// 		return
// 	}
// }
// func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
// 		users := make([]user,0,)
// }
// func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {

// }
// func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {

// }

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusNotFound)
// 	w.Write([]byte(`{"error":"not found"}`))
// }
// func Main() {
// 	mux := http.ServeMux()
// 	userH := &userHandler{
// 		store: &dataStore{
// 			m: map[string]user{
// 				"1": user{
// 					Id:   "1",
// 					Name: "yene",
// 				},
// 			},
// 			RWMutex: &sync.RWMutex{},
// 		},
// 	}
// 	mux.Handle("/users/", userH)
// 	mux.Handle("/users", userH)
// 	http.ListenAndServe("localhost:8080", mux)
// }
