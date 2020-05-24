package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(fn string) []byte {
	dat, err := ioutil.ReadFile(fn)
	check(err)
	return dat
}

func main() {
	fmt.Println(string(read(os.Args[1])))
}
