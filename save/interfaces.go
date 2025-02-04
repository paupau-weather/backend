package save

import (
	"fmt"
)

type Shape interface{
	Area()
	Perimetr()
}

type Rectangle struct{
	Ar int
	Per int
}

type Circle struct{
	Ar int
	Per int
}

type Triangle struct{
	Ar int
	Per int
}

func (r Rectangle) Area(){
	fmt.Println("Площадь: ", r.Ar)
}

func (r Rectangle) Perimetr(){
	fmt.Println("Периметр: ", r.Per)
}

func (c Circle) Area(){
	fmt.Println("Площадь: ", c.Ar)
}

func (c Circle) Perimetr(){
	fmt.Println("Периметр: ", c.Per)
}

func (t Triangle) Area(){
	fmt.Println("Площадь: ", t.Ar)
}

func (t Triangle) Perimetr(){
	fmt.Println("Периметр: ", t.Per)
}

func print(shape []Shape){
	for _, shape := range shape {
		shape.Area()
		shape.Perimetr()
	   }
}
/*
func main() {
	var triangle Shape = Triangle{23, 40}
	var circle Shape = Circle{18, 20}
	var rectangle Shape = Rectangle{15, 30}

	Shape := []Shape{triangle, circle, rectangle}
	print(Shape)
}*/