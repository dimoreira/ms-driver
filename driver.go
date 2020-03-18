package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type iDriver struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func homeEndpoint(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"microservice": "driver",
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Add(
		"content-type",
		"application/json",
	)
	w.Write([]byte(responseData))
}

func loadDrivers() []iDriver {
	jsonFile, err := os.Open("drivers.json")
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error())
	}

	var drivers []iDriver
	json.Unmarshal(data, &drivers)

	return drivers
}

func listDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := loadDrivers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}

func getDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var drivers []iDriver = loadDrivers()

	for _, driver := range drivers {
		if driver.Uuid == params["driverId"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(driver)
			return
		}
	}

	// Return if error
	errorResponse := map[string]string{
		"error": "Driver not found",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(errorResponse)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeEndpoint).Methods("GET")
	r.HandleFunc("/drivers", listDrivers).Methods("GET")
	r.HandleFunc("/drivers/{driverId}", getDriver).Methods("GET")
	fmt.Println("Server started")
	http.ListenAndServe(":8081", r)
}
