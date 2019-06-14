package main

import "fmt"

func main(){
	var getLuasKeliling = func(panjang int, lebar int ) (int, int) {
		return panjang*lebar, 2*(panjang+lebar)
	}

	var getMinMax = func(n []int) (int , int){
		var min , max int 
		for i, number := range n{
			if i == 0 {
				min , max = number, number
			}else if number>max{
				max = number
			}else if number<min{
				min = number
			}
		}
		return min, max
	}

	var luas , keliling = getLuasKeliling(4,5)
	fmt.Print("Luas: ",luas,"  \nKeliling : ",keliling)

	var a = []int {4,4,54,54,54,1}
	var min , max = getMinMax(a)
	fmt.Println("\nMin : ",min,"    Max:   ",max)
}
