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

	//массивы()
	var faf [16]int

	for a := 0; a < len(faf); a++ {
		faf[a] = a
	}
	fmt.Println(faf)


				for a := 0; a < len(faf); a++ {
					if faf[a]%2 == 0 {
						fmt.Print(faf[a], " ")
					} else {
						fmt.Print("* ")
					}

				}
			c := 0
			for a := 0; a < len(faf); a++ {
				if faf[a]%2 == 1 {

					c = c + faf[a]
				}
			}
			fmt.Print(c)

		s := make(map[string]int)
		s["Яна кака"] = 2
		fmt.Println(s["Яна кака"]) */
	for {

		h := menu()
		fmt.Println("Введите первое число")
		a := 0
		b := 0
		fmt.Scanf("%d\n", &a)
		fmt.Println("Введите второе число")
		fmt.Scanf("%d\n", &b)

		//fmt.Scanf("%d\n", &s)
		switch h {
		case 1:
			fmt.Println(sum(a, b))
		case 2:
			fmt.Println(vich(a, b))
		case 3:
			fmt.Println(umn(a, b))
		case 4:
			fmt.Println(del(a, b))
		case 5:
			fmt.Println(voz(a, b))
		case 6:
			return

		}
	}

}
func sum(a int, b int) int {

	return a + b

}
func vich(a int, b int) int {

	return a - b

}
func umn(a int, b int) int {

	return a * b

}
func del(a int, b int) int {

	return a / b

}
func voz(a int, b int) int {

	if b == 1 {
		return a
	} else {
		return a * voz(a, b-1)
	}

}
func menu() int {
	fmt.Println("Выберите пункт меню:\n1 - сложение\n2- вычитание\n3- умножение\n4-деление\n5-возведение в степень\n6-выход")
	var s int

	fmt.Scanf("%d\n", &s)
	return s
}
