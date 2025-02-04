package save

import (
	"fmt"
)

//Поиск минимального и максимального значений в массиве
func MaxAndMin(arr []int){
	for i:=0; i < len(arr); i++{
		for j:=0; j < len(arr) - 1; j++{
			if arr[j] > arr[j + 1]{
				var term int
				term = arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = term
			}
		}
	}
	fmt.Print("Минимум: ", arr[0], " ", "Максимум: ", arr[len(arr) - 1])
}


func Max(arr []int){
	max := arr[0]
	min := arr[0] 
	for i:=1; i < len(arr); i++{
		if arr[i] > max{
			max = arr[i]
		}
		if arr[i] < min{
			min = arr[i]
		}
	}
	fmt.Print(max, min)
}

//Переворот массива
func Reverse(arr []int){
	middle := len(arr) / 2
	for i:=0; i< middle; i++{
		var term int
		term = arr[i]
		arr[i] = arr[len(arr) - 1 - i]
		arr[len(arr) - 1 - i] = term
	}
	fmt.Print(arr)
}

//Удаление дубликатов в массиве
func RemoveDuplicates(arr []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, num := range arr {
        if !seen[num] {
            seen[num] = true
            result = append(result, num)
        }
    }
    return result
}