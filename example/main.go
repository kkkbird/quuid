package main

import (
	"fmt"

	"github.com/kkkbird/quuid"
)

func main() {
	id := quuid.UUID()
	fmt.Println(id)

	id2 := quuid.New(quuid.WithHWAddressPrefix)
	fmt.Println(id2.UUID())
}
