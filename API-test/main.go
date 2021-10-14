package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

/**
 * Web API to add, update, and remove people from a data store
 * @author Jonathan
 * @version 1.0
 * @since 2021-10-14
 */

//user defined type to represent person data
type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Balance float32 `json:"balance"`
	Email   string  `json:"email"`
	Address string  `json:"address"`
}

//array of user defined type Person
var People []Person

//array type to sort array members by name
type SortByName []Person

//sorting algorithm
func (a SortByName) Len() int           { return len(a) }
func (a SortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

/**
* allPeople method to populate and sort array by names
*
* @param input   interface, struct (response writer and request input)
 */
func allPeople(w http.ResponseWriter, r *http.Request) {
	People := []Person{
		{Name: "Zeb", Age: 25, Balance: 1800.33, Email: "mzeb@internet.com", Address: "155 Bay Road"},
		{Name: "Jon", Age: 50, Balance: 200.20, Email: "jon@internet.com", Address: "1 Fox Avenue"},
		{Name: "Andy", Age: 35, Balance: 900.20, Email: "and123@internet.com", Address: "58 Derby Street"},
		{Name: "Mary", Age: 39, Balance: 1900.70, Email: "mary@internet.com", Address: "58 Derby Street"},
	}
	sort.Sort(SortByName(People))
	fmt.Println("All people endpoint hit")
	json.NewEncoder(w).Encode(People)
}

/**
* createNewPerson method to create a person record
*
* @param input   interface, struct (response writer and request input)
 */
func createNewPerson(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newPerson Person
	json.Unmarshal(reqBody, &newPerson)
	People = append(People, newPerson)
	json.NewEncoder(w).Encode(newPerson)
}

/**
* deletePerson method to delete a person record
*
* @param input   interface, struct (response writer and request input)
 */
func deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	for index, article := range People {
		if article.Name == name {
			People = append(People[:index], People[index+1:]...)
		}
	}
}

/**
* handleRequests method with HandleFunc calls
 */
func handleRequests() {
	http.HandleFunc("/app/people", allPeople)
	http.HandleFunc("/app/newperson", createNewPerson)
	http.HandleFunc("/app/newperson/{name}", deletePerson)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

/**
 * main method
 */
func main() {
	handleRequests()
}
