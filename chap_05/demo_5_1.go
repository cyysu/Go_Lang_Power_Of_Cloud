
package chap_05
//package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Lockmap struct{
	sync.Mutex
	m map[string]int	
}

var lockmap = make([]Lockmap,26)

// 包初始化
func init(){
	for i := range lockmap {
		lockmap[i].m = make(map[string]int)
	}
}

// 生成随机小写字母字符串
func gen_rand_string() (retval string) {
	count := rand.Uint32() % 20 + 1;
	for i := uint32(0); i < count; i++ {
		retval += fmt.Sprintf("%c",(rand.Uint32() % 26) + 'a')
	}
	return
}

// 字符串计数
func counter(s string){
	
	idx := int(s[0] - 'a')
	time.Sleep(time.Duration(idx) * time.Millisecond)

	lockmap[idx].Lock()
	defer lockmap[idx].Unlock()
	
	// 引用的元素不存在的时候会自动创建零值元素
	lockmap[idx].m[s]++
}

// 入口
func Demo_5_1(){
	
	for i := 0; i < 20; i++{
		go counter(gen_rand_string())
	}
	
	time.Sleep(time.Second)
	
	for idx := range lockmap {
		lockmap[idx].Lock()
		defer lockmap[idx].Unlock()
		for k,v := range lockmap[idx].m {
			fmt.Printf("[%d] %s = %d\n",idx,k,v)
		}
	}	
}

/*
func main(){
	Demo_5_1()
}
*/
