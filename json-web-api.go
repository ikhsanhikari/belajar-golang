package main

import "fmt"
import "encoding/json"
import "net/http"


type Student struct {
	Id string
	Name string
	Phone string
	Email string 
}

func getUser(w http.ResponseWriter, r *http.Request){
	var data = []Student{
		{"1","Nurikhsan1","089656541471","ikhsanhikari29@gmail.com"},
		{"2","Nurikhsan2","089656541471","ikhsanhikari29@gmail.com"},
		{"3","Nurikhsan3","089656541471","ikhsanhikari29@gmail.com"},
	}
	if r.Method == "GET"{
		var jsonData , err = json.Marshal(data)
		if err!=nil{
			fmt.Print(err.Error())
			return
		}
		w.Write(jsonData)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func getUserById(w http.ResponseWriter, r *http.Request){
	var data = []Student{
		{"1","Nurikhsan1","089656541471","ikhsanhikari29@gmail.com"},
		{"2","Nurikhsan2","089656541471","ikhsanhikari29@gmail.com"},
		{"3","Nurikhsan3","089656541471","ikhsanhikari29@gmail.com"},
	}
	
	if r.Method == "GET"{
		var id = r.FormValue("id")
		for _, each := range data{
			if each.Id == id{
				var jsonData , err = json.Marshal(each)
				if err!=nil{
					fmt.Print(err.Error())
					return
				}
				w.Write(jsonData)
				return
			}
		}
		
	}
	http.Error(w, "", http.StatusBadRequest)
}

func saveUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		w.Header().Set("Content-Type","application/json")
		studentRequest := []Student{}
		error:= json.NewDecoder(r.Body).Decode(&studentRequest)
		if error!=nil{
			fmt.Print(error.Error())
		}
		studentResponse , err:= json.Marshal(studentRequest)
		if err!=nil{
			fmt.Print(err.Error())
		}
		w.WriteHeader(http.StatusOK)
		w.Write(studentResponse)
		return
	}

	http.Error(w,"ini pesan",http.StatusBadRequest)
	
}


func main(){
	http.HandleFunc("/",getUser)
	http.HandleFunc("/user",getUserById)
	http.HandleFunc("/user/save",saveUser)	
	fmt.Print("localhost:8989")

	http.ListenAndServe(":8989",nil)
}