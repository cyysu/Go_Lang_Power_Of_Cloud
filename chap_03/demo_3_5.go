
package chap_03

import (
	"fmt"	
)

func Demo_3_5(){
	str := "好好学习,天天向上"
	for i,ch := range str {
		fmt.Printf("%d: %c\n",i,ch);
		i++
	}
	
	array := []int{4,5,6,8}
	for i,v := range array {
		fmt.Printf("%d: %v\n",i,v)
		i = 100
	}
}
