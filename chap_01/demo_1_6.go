
package chap_01

import (
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"strings"
)

var (
	esc = html.EscapeString	// 重命名函数
	html_opt *bool = flag.Bool("html",true,"html output")
)

func Demo_1_6(){
	
	flag.Parse()	// 解析命令行
	
	// 读入标志文件,进行解析(结果是字符串数组),输出结果
	in,_ := ioutil.ReadAll(os.Stdin)
	out := parse(string(in))
	
	for i := range out {
		fmt.Println(out[i])
	}
}

func parse(in string) []string {
	
	segs := strings.Split(in,"\n\n")
	
	for i := 0; i < len(segs); i++ {
		
		seg := segs[i]
		
		if seg == "" {
			continue
		}
		if seg[0] == '\n' {
			seg = seg[1:]
		}
		// 普通的短段落
		if len(seg) < 4 {
			segs[i] = para(seg)
			continue
		}
		
		// 前两个字符
		switch seg[:2] {
		case "01": segs[i] = header(seg)
		case "02": segs[i] = importFile(seg)
		default: segs[i] = para(seg)
		}
	}
	
	return segs
}

// 处理段落
func para(seg string) string{
	// 普通格式输出
	if !*html_opt {
		return seg
	// HTML格式输出
	}else{
		seg = esc(seg)	// 处理转义字符
		// 规则2: 以空白字符开头则作为HTML的pre
		if seg[0] == ' ' || seg[0] == '\t' {
			seg = strings.Replace(seg,"\t","    ",-1)
			return "<pre>" + seg + "</pre>"
		// 规则3: 简单段落对应HTML的p
		}else{
			return "<p>" + seg + "</p>"
		}		
	}
	//==== Go解释器/编译器似乎有bug,没有下面这个语句则会提示出错
	// function ends without a return statement	
	return seg
}

// 处理标题
func header(seg string) string {
	// 非HTML输出
	if !*html_opt {
		return "\t" + seg[4:]
	}else{
		// 01<标题级别><空格><标题>
		level := string(seg[2])
		title := esc(seg[4:])
		seg = "<h" + level + ">" + title + "</h" + level + ">"
		return seg
	}
	//==== Go解释器/编译器似乎有bug,没有下面这个语句则会提示出错
	// function ends without a return statement	
	return seg
}

// 导入文件
func importFile(seg string) string {
	// 020<空格><要导入的文件名>
	cont_bytes,err := ioutil.ReadFile(seg[4:])
	var cont_string string
	if err != nil {
		cont_string = fmt.Sprintf("Error: %v",err)
	}else{
		cont_string = string(cont_bytes)
	}
	return para(cont_string)
}
