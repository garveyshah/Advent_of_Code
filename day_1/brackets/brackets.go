package brackets

import (
	"fmt"
)

func SortBrackets(s string) uint {
	var move uint
	for _, char := range s {
		switch {
		case char == '(' :
			move += 1
			fmt.Println("moves + 1", move)
		case char == ')':
			move -= 1
			fmt.Println("move - 1", move)
		default:
		}
	}
	
	return move
}
