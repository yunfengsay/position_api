package models

import (
	"fmt"
)

type Message struct {
	Content string
}

func init() {
	m := Message{"models init ok"}
	fmt.Println(m)
}
