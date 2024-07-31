package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello go"))
// 	})
// 	port := ":9090"
// 	fmt.Printf("Serving on port %s", port)
// 	http.ListenAndServe(port, mux)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// func Main1() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("GET /", c)
// 	port := ":9090"
// 	fmt.Printf("Serving on port %s", port)
// 	http.ListenAndServe(port, mux)
// }
// func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	if err := json.NewEncoder(w).Encode(v); err != nil {
// 		return fmt.Errorf("encode json: %w", err)
// 	}
// 	return nil
// }

// func decode[T any](r *http.Request) (T, error) {
// 	var v T
// 	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
// 		return v, fmt.Errorf("decode json: %w", err)
// 	}
// 	return v, nil
// }
// func c(w http.ResponseWriter, r *http.Request) {
// 	encode(w, r, http.StatusCreated, struct {
// 		Name string `json:"name,omitempty"`
// 	}{
// 		Name: "Hundera",
// 	})
// }
