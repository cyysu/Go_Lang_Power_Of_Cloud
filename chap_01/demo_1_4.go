
package chap_01

import (
	"fmt"
	"sort"
)

// 二分查找法来猜数
func binary_search(){

	min,max := 0,100
	
	fmt.Printf("请想一个[%d,%d)之间的整数。\n",min,max)
	
	for min < max {
		
		mid := (min + max) / 2
		fmt.Printf("该数小于或者等于%d吗？(y/n)",mid)
		
		var s string
		fmt.Scanln(&s)
		
		if s != "" && s[0] == 'y'{
			max = mid
		}else{
			min = mid + 1
		}
	}
	
	fmt.Printf("该数是: %d\n",max)
}

// 通过sort包中的函数来猜数
func guess_by_sort(){
	
	min,max := 0,100
	
	fmt.Printf("请想一个[%d,%d)之间的整数。\n",min,max)
	
	fmt.Printf(
		"该数是: %d\n",
		sort.Search(
			100,
			func(i int) bool {
				
				fmt.Printf("该数小于或者等于%d吗？(y/n)",i)
				
				var s string
				fmt.Scanln(&s)
				
				return s != "" && s[0] == 'y'
	}))
}

func Demo_1_4(){
	guess_by_sort()
}
