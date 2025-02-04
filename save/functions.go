package save

import (
	"fmt"
)

//Факториал числа
func factorial(n int){
	fact := 1
	for i:=2; i<=n; i++{
		fact = fact * i
	}
	fmt.Print(fact)
}

//Рекурсивный способ
func RecFactorial(n int) int{
	if n == 0{
        return 1
    }
    return n * RecFactorial(n - 1)
}

//Функция, которая возвращает несколько значений
func function(x, y int) (int, int){
	var Sum int
	var Sub int
	Sum = x + y
	Sub = x - y
	return Sum, Sub
}

//Калкулятор
func Calculator(x, y int){
	Sum := func (x, y int) int{return x + y}
	Sub := func (x, y int) int{return x - y}
	Mul := func (x, y int) int{return x * y}
	Div := func (x, y float64) float64{return float64(x) / float64(y)}
	fmt.Println("Сумма = ", Sum(x, y))
	fmt.Println("Разность = ", Sub(x, y))
	fmt.Println("Умножение = ", Mul(x, y))
	fmt.Print("Деление = ", Div(float64(x), float64(y)))
}