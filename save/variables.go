package save

import (
	"fmt"
)

//Нахождение периметра прямоугольника
func Perimeter(width, height int) int{
	return (2 * width) + (2 * height)
}

//Конвертирование градусов Фарингейта в градусы Цельсия
func Convert(Fahrenheit float64) float64{
	var celsius float64
	celsius = (Fahrenheit - 32) * (5.0/9.0)
	return celsius
}

//Нахождение средного арифметического
func MiddleCount(array []int) (n float64){
	var Sum int
	for i := 0; i < len(array); i++{
		Sum += array[i]
	}
	n = float64(Sum) / float64(len(array))
	return n
}

//Четность числа
func Parity(){
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	if n % 2 == 0{fmt.Print("Число четное")}else{fmt.Print("Число не чётное")}
}