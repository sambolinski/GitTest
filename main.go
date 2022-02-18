package main

import (
	"fmt"
	"time"
)

func main() {
	pingTime := time.Second * 3
	ticker := time.NewTicker(pingTime)

	for _ = range ticker.C {
		fmt.Println("Hello")
	}
}
