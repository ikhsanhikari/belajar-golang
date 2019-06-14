package main

import "fmt"

func penjumlahan(numbers ...int ){
	sum := 0
	for _, number := range numbers{
		sum+=number
	}
	fmt.Print(sum)
}

func hobby(name string , hobbies ...string){
	fmt.Println("Name : ",name )
	fmt.Print("Hobbies : ")
	var tampHob = ""
	for _,hob := range hobbies{
		tampHob+=hob+", "
	}
	fmt.Print(tampHob)
}

func main(){
	// penjumlahan(2,3,3,3,877789)
	hobby("Nurikhsan", "Foot Ball","Swimming ")
}