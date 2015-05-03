
package chap_04

import (
	"fmt"	
)

func protect(g func()){
	defer func(){
		fmt.Println("done")
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v\n",x)
		}
	}()
	fmt.Println("start")
	g()
}

func Demo_4_7(){
	/*
	defer func(){
		if err := recover(); err != nil {
			fmt.Printf("main中的recover(): %v\n",err)
		}
	}()
	*/
	var s []byte
	protect(func() { s[0] = 0 })
	protect(func() { panic("主动触发的panic")})
	
	/*
		为什么这两句的效果不同?
		可以这样解释吗?
		recover()会显示出错信息和调用栈,然后继续传播panic(),
		而自己的defer函数不会继续传播
	*/
	/*
	defer func() { 
		if err := recover(); err != nil {
			fmt.Printf("main中的recover(): %v\n",err)
		}
		fmt.Println("==")
	} ()
	*/
	defer func() { recover() }()
	//defer recover()
	s[0] = 42
}
