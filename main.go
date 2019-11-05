package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"./fs"
)

type portStruct struct {
	Port string
}

func main() {
	r := mux.NewRouter()
	port := ":8080"

	//get example
	r.HandleFunc("/{potato}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		potato := vars["potato"]

		fmt.Fprintf(w, "You get the get ! %s", potato)
	})

	//post example
	r.HandleFunc("/ChangeMyPort/", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var t portStruct
		err := decoder.Decode(&t)

		if err != nil {
			panic(err)
		}

		fmt.Println(t.Port)

		fs.NewPort(t.Port, "./test/try.txt")

		fmt.Fprintf(w, "The port has potatoed ! %s", t.Port)
	}).Methods("POST")

	fmt.Printf("Server running on %s \n", port)
	http.ListenAndServe(port, r)
}
