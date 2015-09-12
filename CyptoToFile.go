package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"io/ioutil"
	"io"
)
const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals = "0123456789"
	Special = "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:` "
	Ascii = Alphabet + Numerals + Special
)

var readKeyLine [len(Ascii)+1]string

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

func convertLine(txtpath string){

	file := createFile("/home/bilal/Desktop/encrypted.txt")
	defer closeFile(file)

	dat, err := ioutil.ReadFile(txtpath)
	if err != nil {
		panic(err)
	}
	i := 0
	j := 0
	for i < len(string(dat)) {
		for j < len(Ascii){
			if Ascii[j] == string(dat)[i] {
				io.WriteString(file, readKeyLine[j] )
				//writeFile(file, readKeyLine[j], )
				j = len(Ascii)+1
			}
			j++
		}
		j = 0
		i++
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

func main(){
	//Read Key File
	fmt.Println("Key Path ?")
	path := ""
	fmt.Scanf("%s", &path)
	readKeyFile(path)
	fmt.Println("Convert to Txt Path ?")
	txtpath:=""
	fmt.Scanf("%s", &txtpath)
	convertLine(txtpath)
}
