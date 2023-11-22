package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	l := NewLinkedList[int]()
	l.Append(1)
	l.Append(4)
	l.Append(3)
	l.Append(20)
	l.Append(10)

	for x := 0; x <= 10; x++ {
		lb, _ := json.Marshal(l)
		fmt.Printf("%v\n", string(lb))
		var t *LinkedList[int]
		_ = json.Unmarshal(lb, &t)
		fmt.Printf("%v\n", t)
	}
}
