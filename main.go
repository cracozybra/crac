package main

import (
	"fmt"
	"math/rand"
)

func main() {
	memes()

}
func meme() {
	p := 0
	c := 0
	for (c != 3) && (p != 3) {

		fmt.Println("Введите: камень, ножницы или бумага")
		a := ""
		fmt.Scanf("%s\n", &a)
		arr := [3]string{"бумага", "ножницы", "камень"}
		x := arr[rand.Intn(len(arr))]
		fmt.Println("Компьютер выбрал ", x)

		if a == x {
			fmt.Println("Ничья :(")
		} else {

			if (a == "камень" && x == "ножницы") || (a == "бумага" && x == "камень") || (a == "ножницы" && x == "бумага") {

				p++
			} else {

				c++
			}

		}
		fmt.Println("Ваш счёт:", p, "\nСчёт компьютера: ", c)

	}
	if p > c {
		fmt.Println("Вы выйграли!!!")
	} else {
		fmt.Println("Я выйграл :)")
	}
}
func memes() {

}
