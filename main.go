package main


import (
	//"fmt"
	"net/http"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Monster model
type Monster struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Power int `json:"power"`
	Block int `json:"block"`
	Owner *Owner `json:"owner"`
}

// Owner model 
type Owner struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
}

// init monsters variable as a slice

var monsters []Monster

//methods

func getMonsters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(monsters)
}

func getMonster(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // takes params from request
    for _, item := range monsters {
		if item.ID == params["id"] {
		json.NewEncoder(w).Encode(item)
		return
	}
  }
    json.NewEncoder(w).Encode(&Monster{})
}

func addMonster(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var monster Monster
	_ = json.NewDecoder(r.Body).Decode(&monster)
	monster.ID = strconv.Itoa(rand.Intn(10000000))
	monsters = append(monsters, monster)
	json.NewEncoder(w).Encode(monster)
}

func updateMonster(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range monsters {
		if item.ID == params["id"] {
		monsters = append(monsters[:index], monsters[index+1:]...)
	    var monster Monster
		_ = json.NewDecoder(r.Body).Decode(&monster)
		monster.ID = params["id"]
		monsters = append(monsters, monster)
		json.NewEncoder(w).Encode(monster)
		return
		}
	}
}

func deleteMonster(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range monsters {
		if item.ID == params["id"] {
		monsters = append(monsters[:index], monsters[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(monsters)
}


func main() {
	
 // Init router
 r := mux.NewRouter()

 // data

monsters = append(monsters, Monster{ID: "1", Name:"Devox", Power: 150, Block: 88, Owner: &Owner{Name: "Marco", Surname: "Cronos"}})
monsters = append(monsters, Monster{ID: "2", Name:"Siruoh", Power: 110,Block: 56, Owner: &Owner{Name: "Nathan", Surname: "Pure"}})
monsters = append(monsters, Monster{ID: "3", Name:"Jenir", Power: 240, Block: 6, Owner: &Owner{Name: "Mina", Surname: "Murphy"}})

 // route handlers
 r.HandleFunc("/api/monsters", getMonsters).Methods("GET")
 r.HandleFunc("/api/monsters/{id}", getMonster).Methods("GET")
 r.HandleFunc("/api/monsters", addMonster).Methods("POST")
 r.HandleFunc("/api/monsters/{id}", updateMonster).Methods("PUT")
 r.HandleFunc("/api/monsters/{id}", deleteMonster).Methods("DELETE")

 //server listening

 log.Fatal(http.ListenAndServe(":5000", r))

}