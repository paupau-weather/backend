package save

import (
	"fmt"
)

//Вывод таблицы умножения
func MultiplicationTable(){
	for i:=1; i<10; i++{
		for j:=1; j<10; j++{
			fmt.Print("\t", i * j)
		}
		fmt.Println()
	}
}

//Числа Фибоначчи
func Fibonacci(){
	var limit int
	fmt.Print("Введите предел: ")
	fmt.Scan(&limit)
	Sum := 1 
	k := 0
	var term int
	fmt.Println(k)
	for Sum <= limit{
		fmt.Println(Sum)
		term = Sum
		Sum = k + Sum
		k = term
		if Sum > limit {
			break
		 }
	}
}

//НОД двех чисел
func NOD(){
	var a int
	var b int
	fmt.Print("Введите число а: ")
	fmt.Scan(&a)
	fmt.Print("Введите число b: ")
	fmt.Scan(&b)
	for b != 0{
		var term int
		term = b
		b = a % b
		a = term
	}
	fmt.Print(a)
}

//Поиск простого числа в диапазоне
func Search(a, b int){
	for i:=a; i<=b; i++{
		if(i % 2 != 0){fmt.Println("Простое число: ", i)}
	}
}