package save

import (
	"fmt"
	"strings"
)

//Полиндром
func isPalindrome(x string) bool {
    if(x == ""){return false}
	for i:=0;i<len(x)/2;i++{
		if(x[i] != x[len(x) - 1 - i]){
			return false
		}
	}
	return true
}

//Подсчет гласных в строке
func VowelsAndConsonants(str string){
	count := 0
    vowels := []rune{'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я', 'a', 'e', 'i', 'o', 'u'}
    lowerStr := strings.ToLower(str)
    for _, r := range lowerStr {
        for _, v := range vowels {
            if r == v {
                count++
                break
            }
        }
    }
    fmt.Println(count)
}

//Разбиение строки на слова
func FilteredWord(str string){
	delimiters := []string{",", ";", ".", "?", "!"}
    commonDelimiter := " "

    // Заменяем все разделители на общий разделитель
    for _, delimiter := range delimiters {
        str = strings.ReplaceAll(str, delimiter, commonDelimiter)
    }

	words := strings.Split(str, commonDelimiter)

	var filteredWords []string
       for _, word := range words {
           if word != "" {
               filteredWords = append(filteredWords, word)
           }
       }
	
	fmt.Println(filteredWords)
}

//Поиск подстроки
func Substring(str, substr string){
	index := strings.Index(str, substr)
	fmt.Print(index)
}