package main

import "fmt"

type hitung interface{
	luas() float64
	keliling() float64
}

type persegiPanjang struct{
	panjang float64
	lebar float64
}

func (p persegiPanjang) luas() float64{
	return p.panjang * p.lebar
}
func (p persegiPanjang) keliling() float64{
	return 2*(p.panjang+p.lebar)
}
func main(){
	var bangunDatar  hitung 
	bangunDatar = persegiPanjang{5,6}

	fmt.Println("Keliling : ",bangunDatar.keliling())
	fmt.Println("luas : ",bangunDatar.luas())

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	
	for _, each := range fruits {
		fmt.Println(each)
	}
	
}