package main

import "fmt"

func main() {
	/*imia := 10

	fmt.Print("popka \n ianka \t kakaska", " \n axaxa \n")
	if imia == 10 {
		imia += 10
	} else {
		imia -= 10
	}
	fmt.Print("imia=", imia)

	//(цикл) for if(условие)
	for i := 0; i != 10; i++ {
		if i%2 == 0 {
			fmt.Print("\n", i)
		}


	}
	*/
	//массивы()
	var faf [16]int

	for a := 0; a < len(faf); a++ {
		faf[a] = a
	}
	fmt.Println(faf)

	/*
		for a := 0; a < len(faf); a++ {
			if faf[a]%2 == 0 {
				fmt.Print(faf[a], " ")
			} else {
				fmt.Print("* ")
			}

		}*/
	c := 0
	for a := 0; a < len(faf); a++ {
		if faf[a]%2 == 1 {

			c = c + faf[a]
		}
	}
	fmt.Print(c)
}
