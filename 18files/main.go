package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GoLang")
	content := "This needs to go in a file - LearnCodeOnline.in"
	fileName := "./myfile.txt"
	file, err := os.Create(fileName)
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length of the content is ", length)
	defer file.Close()
	ReadFile(fileName)
}

func ReadFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is\n", string(dataByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
