package main

import "fmt"

var contohMap = []map[string]string{
	{"nama": "Nurikhsan","alamat": "cirebon","test": "cirebon"},
	{"nama": "Nurikhsan2","alamat": "cirebon2"},
}





func main(){
	for _,data := range contohMap{
		fmt.Println(data["nama"],":",data["alamat"],":",data["test"])
	}
}