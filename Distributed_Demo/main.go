package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// http.HandleFunc的作用是注册HTTP处理函数Handler，如果由客户端访问本机的HTTP服务
	// 且以“/objects/”开头，那么请求将由Handler负责处理。
	http.HandleFunc("/objects/", Handler)
	println("server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8006", nil))
}

// Handler函数，HTTP中最重要的请求和响应Response，Request为参数，
// 根据客户端不同的请求方式，执行不同的处理函数：Put函数与Get函数。
func Handler(w http.ResponseWriter, r *http.Request) {
	println(r)
	m := r.Method
	if m == http.MethodPut {
		Put(w, r)
		return
	}
	if m == http.MethodGet {
		Get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

// Put函数，r.URL变量记录HTTP请求的URL，EscapedPath方法用于获取结果转义以后的路径部分，
// 该路径形式是：/objects/<object_name>，然后strings.Split函数功能是分割/objects/<object_name>，
// 分割为"“、”objects"、<object_name>,去数组的第三个元素就是<object_name>，os.Create在本地文件系统的根存储目录创建同名文件f，
// 创建成功将r.Body用io.Copy写入文件f
func Put(w http.ResponseWriter, r *http.Request) {
	//C:\Users\Administrator\go\src\awesomeProject\test_file
	f, e := os.Create(("C:/Users/Administrator/go/src/awesomeProject/test_file" + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2]))

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}

// Get 函数负责处理HTTP的Get请求，从本地硬盘上读取内容并将其作为HTTP的响应输出
func Get(w http.ResponseWriter, r *http.Request) {

	f, e := os.Open(("C:/Users/Administrator/go/src/awesomeProject/test_file" + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2]))

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
