package save

type Student struct{
	name string
	age int
	grade int
}

func Sort(arr []Student){
	for i:=0; i<len(arr) - 1; i++{
		if arr[i].age > arr[i + 1].age{
			term := arr[i]
			arr[i] = arr[i + 1]
			arr[i + 1] = term
		}
	}
}

/*Bob := Student{"Боб", 18, 5}
Tom := Student{"Том", 19, 3}
Mark := Student{"Марк", 20, 4}
arr := []Student{Bob, Mark, Tom}
fmt.Println(arr)
Sort(arr)
fmt.Print(arr)*/


/*type Point struct{
	X int
	Y int
}

func Distance(p, p_1 Point){
	dis := math.Sqrt((float64(p_1.X)-float64(p.X)) + (float64(p_1.Y)-float64(p.Y)))
	fmt.Print(dis)
}

dot := Point{5, 6}
	dot_1 := Point{10, 7}
	Distance(dot, dot_1)*/