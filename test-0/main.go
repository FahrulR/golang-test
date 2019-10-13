package main

import f "fmt"
import "strings"

func main() {
	var input string
	check :=[]bool{false,false,false,false,false,}
	var countV int
	var countC int
	f.Printf("Input any word: ")
	f.Scanln(&input)
	input = strings.ToLower(input)
	for _, str:= range input {
	switch str {
		case 'a':
			if check[0] == false {
			countV++
			check[0] = true
			}
		case 'e':
			if check[1] == false {
			countV++
			check[1] = true
			}
		case 'i':
			if check[2] == false {
			countV++
			check[2] = true
			}
		case 'o':
			if check[3] == false {
			countV++
			check[3] = true
			}
		case 'u':
			if check[4] == false {
			countV++
			check[4] = true
			}
		default: 
			countC++
	}
	}

	f.Printf("Huruf mati: %d\n", countC)
	f.Printf("Huruf hidup: %d\n", countV)
}