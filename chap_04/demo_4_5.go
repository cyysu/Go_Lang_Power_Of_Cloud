
package chap_04

import (
	"fmt"	
)

func Demo_4_5(){	
	add := func(base int) func(int) int {
		return func(n int) int {
			return base + n
		}
	}	
	add5 := add(5)	
	fmt.Println(add5(10))
}
