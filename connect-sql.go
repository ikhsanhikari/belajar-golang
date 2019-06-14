package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"


type student struct {
    id    string
    name  string
    age   int
    grade int
}


func connect() (*sql.DB,error){
	db, err := sql.Open("mysql","root:@tcp(localhost:3306)/belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}




func main(){
	connect()
}
