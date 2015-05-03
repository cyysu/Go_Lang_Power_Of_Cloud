
package chap_01

import "fmt"

var (
	memory [30000] byte 	// 内存
	p_data int				// 数据指针
	p_code int				// 代码指针
	program = "++++++++++[>++++++++++<-]>++++.+."
)

func loop(inc int){
	for i := inc; i != 0; p_code += inc {
		switch program[p_code + inc]{
		case '[': i++
		case ']': i--
		}
	}
}

func Demo_1_5(){
	for{
		switch program[p_code]{
		case '>': p_data++
		case '<': p_data--
		case '+': memory[p_data]++
		case '-': memory[p_data]--
		case '.': fmt.Println(memory[p_data])
		case '[': if memory[p_data] == 0 { loop(1) }
		case ']': if memory[p_data] != 0 { loop(-1)}
		default:
			fmt.Println("非法指令")
		}
		
		// 移动指令指针
		p_code++
		if p_code == len(program){
			return
		}
	}
}
