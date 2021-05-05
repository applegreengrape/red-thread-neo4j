package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/applegreengrape/red-thread-neo4j/loader"
	//"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "red thread neo4j api endpoint")
}

func createNode(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "request body error")
	}

	var new Node
	json.Unmarshal(reqBody, &new)
	loader.CreateNode(new.ID, new.Name, new.NodeType)
	json.NewEncoder(w).Encode(new)
}

func CreateRel(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "request body error")
	}

	var rel Rel
	json.Unmarshal(reqBody, &rel)
	loader.CreateRel(rel.PID, rel.ID)
	json.NewEncoder(w).Encode(rel)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/createNode", createNode).Methods("POST")
	router.HandleFunc("/createRel", CreateRel).Methods("POST")
	http.ListenAndServe(":8080", router)
}
