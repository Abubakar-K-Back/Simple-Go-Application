package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

type person struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}
type subject struct{
	ID string `json:"id"`
	Subject_Name string `json:"Sub_Name"`
}
var subjects []subject;
var persons []person;

func getAllPerson(w http.ResponseWriter, r *http.Request)  {
	
	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(persons);
}

func getPersonByid(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type","application/json");
	
	params := mux.Vars(r);
	fmt.Println(params);
	for _, item := range persons{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item);
			return
		}
	}

}

func getPersonByName(w http.ResponseWriter, r *http.Request){	
	w.Header().Set("Content-Type","application/json");
	query:=r.URL.Query();
	fmt.Println(query["name"])
	for _, item := range persons{
		if item.Name==query["name"][0]{
			json.NewEncoder(w).Encode(item)
		}
	}

}

func createPerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var Person person
	_ = json.NewDecoder(r.Body).Decode(&Person);
	Person.ID = strconv.Itoa(rand.Intn(10000000000));
	persons = append(persons, Person);
	json.NewEncoder(w).Encode(Person);
	fmt.Println(Person);

}
func main()  {
	r:=mux.NewRouter();
	
	persons = append(persons, person{ID: "1",Name: "Abubakar",Address: "H.no 285"});
	persons = append(persons, person{ID: "2",Name: "Saad",Address: "H.no 282"});

	r.HandleFunc("/getAllperson",getAllPerson).Methods("GET");
	r.HandleFunc("/getpersonbyID/{id}",getPersonByid).Methods("GET");
	r.HandleFunc("/createPerson",createPerson).Methods("Post");
	r.HandleFunc("/byName",getPersonByName).Methods("GET");

	fmt.Print("Server Starting at port: 8000");
	log.Fatal(http.ListenAndServe(":8000",r))
}