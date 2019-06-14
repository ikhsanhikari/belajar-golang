package main

import "fmt"


type user struct {
	Username string
	Password string
	Email string
}


func (s user ) getName(){
	fmt.Println(s.Username)
}

func main(){
	user1 := []user {
		user {Username: "Nurikhsan", Password: "hikari29", Email:"ikhsanhikari29@gmail.com"},
		user {Username: "Hikari", Password: "hikari29", Email:"ikhsanhikari29@gmail.com"},
		user {Username: "Hikari29", Password: "hikari29", Email:"ikhsanhikari29@gmail.com"},
	}
	userTambahan := user {Username: "Tambahan Lst ", Password: "hikari29", Email:"ikhsanhikari29@gmail.com"} 
	sudahDiAppend:= append(user1,userTambahan)
	
	
	
	for _, data:= range sudahDiAppend{
		fmt.Println(data.Username)
	}

	user2 := user{"Murikhsan jjjjjj","Hikari","hikari@gmail.com"}
	user2.getName()
}