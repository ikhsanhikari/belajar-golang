package main


import "fmt"

func getMap(param []map[string]string) {
	for _, data := range param{
		fmt.Println(data["name"]," - ",data["email"])
	}
} 


func persegiPanjang(p int , l int) (keliling int , luas int){
	luas  = p * l
	keliling = 2 *(p + l)
	return
}



func main(){
	arg := []map[string]string{
		{"name":"Nurikhsan","email":"ikhsanhikari29@gmail.com"},
		{"name":"Nurikhsan2","email":"ikhsanhikari29@gmail.com2"},
	}

	getMap(arg)
	luas,keliling := persegiPanjang(3,4)
	fmt.Println("Luas : ",luas)
	fmt.Println("Kliling : ",keliling)
}