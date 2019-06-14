package main

import (
    "fmt"
	"math/rand"
	"math"
	"time"
	"strings"
)

func main() {
	var name string = "hikari"
	fmt.Println("hello world ")
	fmt.Println("test " + name)

	

	var chicken = map[string]int {
		"januari":50, 
		"februari":40, 
		"maret":34, 
		"april":67, 
	}
	
	fmt.Println("januari", chicken["januari"])

	cekPositiveNumber(-9)

	rand.Seed(time.Now().Unix())
    var randomValue int

    randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	
	var diameter float64 = 15
    var area, circumference = calculate(diameter)

    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
	


	yourHobbies("Nurikhsan", "Foot Ball", "Tennis", "aaaaaa")

	
}

func cekPositiveNumber(point int) {
	if point > 0 {
		fmt.Println("lebih besar dari 0")
	}else if point < 0 {	
		fmt.Println("lebih kecil dari 0")

	}else {
		fmt.Println("0")

	}
}

func randomWithRange(min, max int)int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}
func calculate(d float64)(float64, float64) {
    // hitung luas
    var area = math.Pi * math.Pow(d / 2, 2)
    // hitung keliling
    var circumference = math.Pi * d

    // kembalikan 2 nilai
    return area, circumference
}

// func variadic(numbers ...int)float64 {
// 	var total int = 0
// 	for _, number: = range numbers {
// 		total + = number
// 	}
// 	var avg = float64(total)/float64(len(number))
// 	return avg
// }

func yourHobbies(name string, hobbies ...string) {
    var hobbiesAsString = strings.Join(hobbies, ", ")

    fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

// func sum(nums ...int) {
//     fmt.Print(nums, " ")
//     total: = 0
//     for _, num: = range nums {
//         total +  = num
//     }
//     fmt.Println(total)
// }