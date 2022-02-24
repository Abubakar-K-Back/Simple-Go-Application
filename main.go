package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type subject struct {
	ID        string `json:"id"`
	SubName   string `json:"Sname"`
	Deparment string `json:"Dept"`
}

type person struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Subject []subject `json:"subj"`
}

// var subjects []subject;

var persons []person

func getAllPerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func getPersonByid(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func getPersonByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	fmt.Println(query["name"])
	for _, item := range persons {
		if item.Name == query["name"][0] {
			json.NewEncoder(w).Encode(item)
		}
	}

}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Person person
	_ = json.NewDecoder(r.Body).Decode(&Person)
	Person.ID = strconv.Itoa(rand.Intn(10000000000))
	persons = append(persons, Person)
	json.NewEncoder(w).Encode(Person)
	fmt.Println(Person)

}
func getSubjectofStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item.Subject)
		}
	}
}
func AddSubject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, items := range persons {
		if items.ID == params["id"] {
			items.Subject = append(items.Subject, subject{ID: strconv.Itoa(rand.Intn(10000000000)),
				Deparment: "Physics",
				SubName:   "Geo-Physics",
			})
		}
	}
	json.NewEncoder(w).Encode(persons)

}

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")

	persons = append(persons, person{ID: strconv.Itoa(rand.Intn(10000000000)), Name: "Abubakar", Address: "H.no 285", Subject: []subject{
		{ID: strconv.Itoa(rand.Intn(10000000000)), Deparment: "Physics", SubName: "Analytical Physics"},
		{ID: strconv.Itoa(rand.Intn(10000000000)), Deparment: "Religion", SubName: "Islamiat"},
		{ID: strconv.Itoa(rand.Intn(10000000000)), Deparment: "Software", SubName: "OOP"},
		{ID: strconv.Itoa(rand.Intn(10000000000)), Deparment: "Chemistry", SubName: "Elements"},
	}})

	persons = append(persons, person{ID: strconv.Itoa(rand.Intn(10000000000)), Name: "Saad", Address: "H.no 281", Subject: []subject{
		{ID: strconv.Itoa(rand.Intn(10000000000)), Deparment: "Physics", SubName: "Analytical Physics"},
		{ID: strconv.Itoa(rand.Intn(10000000000)), Deparment: "Chemistry", SubName: "Elements"},
	}})

	r.HandleFunc("/getAllperson", getAllPerson).Methods("GET")
	r.HandleFunc("/getpersonbyID/{id}", getPersonByid).Methods("GET")
	r.HandleFunc("/createPerson", createPerson).Methods("Post")
	r.HandleFunc("/byName", getPersonByName).Methods("GET")
	r.HandleFunc("/seeSubjects/{id}", getSubjectofStudent).Methods("GET")
	r.HandleFunc("/AddSubject/{id}", AddSubject).Methods("POST")

	fmt.Print("Server Starting at port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
