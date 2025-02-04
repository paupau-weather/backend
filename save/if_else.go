package save

import (
	"fmt"
)

//Високосный год
func Leap_year(){
	fmt.Print("Введите год: ")
	var year int
	fmt.Scan(&year)
	if year % 4 == 0 && year / 100 != 0{
		fmt.Println("Год висакосный")
	}else if year % 400 == 0{
		fmt.Println("Год висакосный")
	}else {fmt.Println("Год не висакосный")}
}

//Определние в каком диапазоне находитс введённый час - утро, день, вечер, или ночь.
func Determinant(){
	var hour int
	fmt.Print("Введите час: ")
	fmt.Scan(&hour)
	if hour == 0 || hour < 0 {panic("Час должен находиться в диапазоне 0 < hour <= 24")}
	if hour > 4 && hour < 12{fmt.Print("Сейчас утро")}
	if hour > 12 &&  hour < 15{fmt.Print("Сейчас день")}
	if hour > 15 && hour < 24 {fmt.Print("Сейчас вечер")}
	if hour < 4 {fmt.Print("Сейчас ночь")}
}

//Простой калькулятор
func calculator(){
	var a, b, n int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Println("Выберите действие:\n 1)Сложение \n 2)Вычитание \n 3)Умножение \n 4)Деление")
	fmt.Scan(&n)
	if n == 1 {
		var Sum int 
		Sum = a + b
		fmt.Print(Sum)
	}
	if n == 2 {
		var Subtruction int 
		Subtruction = a - b
		fmt.Print(Subtruction)
	}
	if n == 3 {
		var Multiply int 
		Multiply = a * b
		fmt.Print(Multiply)
	}
	if n == 4 {
		var Division float64 
		Division = float64(a) / float64(b)
		fmt.Print(Division)
	}
}