package main

import (
	"fmt"

	"github.com/Higor-Edgar/booktime.git/config"
)

func main() {
	_ = config.InitDatabase()

	fmt.Println("Hello, World!")
}
