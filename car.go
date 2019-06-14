package main

import (
		"fmt"
		"encoding/json"
		"net/http")

type car struct {
	Vin string
	Name string
}

var data = []car {
	car {"E109R", "Avanza"}, 
	car {"E1009R", "Avanza R GGGGGG"}, 
}

func cars(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
	
}


func main() {
	http.HandleFunc("/cars", cars)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}