package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("/tmp/tmp/tmp", 0777)
	if err != nil {
		panic(err)
	}

	fmt.Print("ok")
}
