package universe

import (
	"fmt"
)

func init() { // Added the function name 'main'
	var input int
	fmt.Scan(&input)
	if input == 42 {
		fmt.Println("Hello Universe")
	} else {
		fmt.Println(input)
	}
}
