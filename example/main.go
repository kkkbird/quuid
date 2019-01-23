package main

import (
	"fmt"

	"github.com/kkkbird/quuid"
)

func main() {
	id := quuid.UUID()
	fmt.Println(id)
}
