package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var readLine [len(Ascii)+1]string

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals = "0123456789"
	Special = "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
	Ascii = Alphabet + Numerals + Special

)

func createFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, randString string) {
	fmt.Fprintln(f, randString)
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func randomKeyGenerator(characterNumber int,filterValue int) string {
	b := make([]byte, characterNumber)
	if filterValue == 1 {
		for i := range b {
			b[i] = Alphabet[rand.Intn(len(Alphabet))]
		}
	} else if filterValue == 2{
		for i := range b {
			b[i] = Numerals[rand.Intn(len(Numerals))]
		}
	} else if filterValue == 3{
		z := 0
		for i := range b {
			if z%2 == 0 {
				b[i] = Numerals[rand.Intn(len(Numerals))]
			} else if z % 2 == 1 {
				b[i] = Alphabet[rand.Intn(len(Alphabet))]
			}
			z++
		}
	} else if filterValue == 4{
		for i := range b {
			b[i] = Special[rand.Intn(len(Special))]
		}
	} else if filterValue == 5{
		z := 0
		for i := range b {
			if z%2 == 0 {
				b[i] = Special[rand.Intn(len(Special))]
			} else if z % 2 == 1 {
				b[i] = Alphabet[rand.Intn(len(Alphabet))]
			}
			z++
		}
	} else if filterValue == 6 {
		z := 0
		for i := range b {
			if z%2 == 0 {
				b[i] = Special[rand.Intn(len(Special))]
			} else if z % 2 == 1 {
				b[i] = Numerals[rand.Intn(len(Numerals))]
			}
			z++
		}
	} else if filterValue == 7 {
		for i := range b {
			b[i] = Ascii[rand.Intn(len(Ascii))]
		}
	}
	return string(b)
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
	characterNumber := 2
	fmt.Println("1 Karakter kac karaktere denk dussun ? e.g  w = 1asd1  1 e 4 gibi  Default = 2")
	fmt.Scanf("%d", &characterNumber)

	// File Name
	t := time.Now()
	formatedTime := t.Format(time.RFC1123)//Time to String

	// Create File
	file := createFile("/home/bilal/Desktop/"+formatedTime+".key")
	defer closeFile(file)

	// Write File
	j:=0
	for j < len(Ascii)+1 {
		randString := randomKeyGenerator(characterNumber,filterValue)
		writeFile(file, randString)
		j++
	}
}
