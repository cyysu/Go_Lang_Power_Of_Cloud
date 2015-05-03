
package chap_01

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var file_idx = make(chan int)

// 包初始化的时候启动一个去程生成文件序号
func init(){
	go func(){
		for idx := 0; ; idx++{
			file_idx <- idx
		}
	}()
}

func Demo_1_7(){
	
	if err := os.Chdir(os.TempDir()); err != nil {
		log.Fatal(err)
	}

	// 注册HTTP请求处理函数
	http.HandleFunc("/",FrontPage)
	http.HandleFunc("/compile",Compile)
	// 启动Web服务器
	log.Fatal(http.ListenAndServe("127.0.0.1:1234",nil))
}

// 写入首页
func FrontPage(html_writer http.ResponseWriter,_ *http.Request){
	html_writer.Write([]byte(html_page))
}

func err(html_writer http.ResponseWriter,e error) bool {
	if e != nil {
		html_writer.Write([]byte(e.Error()))
		return true
	}
	return false
}

// 运行Go程序
func Compile(html_writer http.ResponseWriter,http_req *http.Request){

	// 创建临时文件
	file_name := "play_" + strconv.Itoa(<-file_idx) + ".go"
	
	file,e := os.Create(file_name)
	
	if err(html_writer,e){
		return
	}
	
	defer os.Remove(file_name)
	defer file.Close()
	
	// 写入用户输入的Go程序到临时文件中
	_,e = io.Copy(file,http_req.Body)
	
	if err(html_writer,e){
		return
	}
	
	file.Close()
	
	// 执行包含Go程序的临时文件
	cmd := exec.Command("go","run",file_name)	
	
	prog_out,e := cmd.CombinedOutput()
	
	if err(html_writer,e){
		return
	}

	// 将程序输出写入到页面中
	html_writer.Write(prog_out)
}

const html_page = `
	<!doctype html>
	<html>
		<head>
			<script>
				function compile(){
					
					var req = new XMLHttpRequest();
					req.onreadystatechange = function(){
						if (req && req.readyState == 4){
							document.getElementById("output").innerHTML = req.responseText;
						}
					}
					
					req.open("POST","/compile",true);
					req.setRequestHeader("Content-Type","text/plain;charset=utf-8");
					
					var prog = document.getElementById("edit").value;
					req.send(prog);
				}
			</script>
		</head>
		<body>
			<textarea rows="25" cols="80" id="edit" spellcheck="false">
package chap_01
import "fmt"
func Demo_1_(){
	fmt.Println("Hello,world")
}
			</textarea>
			<button onclick="compile();">run</button>
			<div id="output"></div>
		</body>
	</html>
`
