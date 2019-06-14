package main

import "encoding/json"
import "fmt"

type Data struct {
	FullName string `json:"Name"`
	Grade string
}

func main()  {
	var jsonData =  `{"Name":"Nurikhsan","Grade":"A"}`
	fmt.Println(jsonData)
	
	var temData Data
	
	var err = json.Unmarshal([]byte(jsonData),&temData)
	if err!=nil{
		fmt.Print(err.Error())
		return
	}

	fmt.Println("Name : ",temData.FullName)
	fmt.Println("Grade : ",temData.Grade)


	var dataStatic = Data{"Nurikhsan","A"}
	dataRes ,err := json.Marshal(dataStatic)
	fmt.Print(string(dataRes))

}