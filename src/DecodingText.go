package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"io/ioutil"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals = "0123456789"
	Special = "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:` "
	Ascii = Alphabet + Numerals + Special
)

var readKeyLine [len(Ascii)]string

func readKeyFile(path string){
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readKeyLine[i] = scanner.Text()
		i++
	}
}


func decodingFile(cryptoPath string, keyLineLong int) {
	file := createFile("/home/bilal/Desktop/decoding.txt")
	defer closeFile(file)

	dat, err := ioutil.ReadFile(cryptoPath)
	if err != nil {
		panic(err)
	}
	//dat convert to string because `.substring` ;)
	data := string(dat)
	var i int = 0
	j := 0
	keyLong := keyLineLong//backup
	for keyLineLong < len(data) {
		for j < len(Ascii) {
			if data[i:keyLineLong] == readKeyLine[j] {
				fmt.Fprintf(file, "%c", Ascii[j]) //unit8 to string and write
				j = 96
			}
			j++
		}
		j = 0
		i = i + keyLong
		keyLineLong = keyLineLong+keyLong
	}
}

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

func main (){
	//Read Key File
	fmt.Println("Key Path ?")
	path := ""
	fmt.Scanf("%s", &path)
	readKeyFile(path)

	var keyLineLong int = int(len(readKeyLine[0]))

	fmt.Println("Cyrpto File Path ?")
	cryptoPath := ""
	fmt.Scanf("%s", &cryptoPath)
	decodingFile(cryptoPath, keyLineLong)
}
