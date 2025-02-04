package save

import(
	"fmt"
)

//Замена значений
func Swap(n, m int){
	point := &n
	_point := &m
	term := n
	*point = m
	*_point = term
	fmt.Println(n)
	fmt.Print(m)
}

//Увеличение возраста объекта структуры Students
/*type Students struct{
	name string
	age int
}

func (p *Students) AgeInc(){
	p.age ++
}

func main() {
	Bob := Students{"Bob", 21}
	fmt.Println(Bob.age)
	Bob.AgeInc()
	fmt.Print(Bob.age)
}*/