package main

import (
	"fmt"
)
func randomKeyGenerator(){

}

func filter(filterValue int) int {
	temp := 0
	i := 0
	var question [3]string
	fmt.Println("Olsun diyorsan 1, olmasin diyorsan 0 gir")
	question[0] = "Harf olsun mu ?"
	question[1] = "Sayi olsun mu ?"
	question[2] = "Ozel Karakter olsun mu ?"

	for i<3 {
		fmt.Println(question[i])
		fmt.Scanf("%d", &temp)
		if temp == 1 {
			if i == 0 {
				filterValue = filterValue+1
			} else if i == 1 {
				filterValue = filterValue+2
			} else if i == 2 {
				filterValue = filterValue+4
			}
		}
		i++
	}

	return filterValue
}

func main(){
	filterValue := 0
	filterValue = filter(filterValue)

	characterNumber := 0
	fmt.Println("1 Karakter kac karaktere denk dussun ? e.g  w = 1asd1  1 e 4 gibi")
	fmt.Scanf("%d", &characterNumber)

	//-------TEST------
	fmt.Println(filterValue)
	fmt.Println(characterNumber)


}
