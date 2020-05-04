package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Println("Hello", os.Getenv("WHO")+"!")

		time.Sleep(time.Second * 1)
	}
}
