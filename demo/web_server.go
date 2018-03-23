package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

func main() {
	http.HandleFunc("/", sayHello) //设置访问的路由
	http.HandleFunc("/login",login)
	http.HandleFunc("/login1",login1)

	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.ParseForm())
	fmt.Println(r.Header)
	fmt.Println(r.URL)
	fmt.Println(r.ContentLength)
	for k,v := range r.Form {
		fmt.Println("key is ",k)
		fmt.Println("value is ", v)
	}

	fmt.Fprint(w,"hello ")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("/Users/finup123/IdeaProjects/goDemo/demo/login.gtpl")
		if err != nil{
			fmt.Println("error: ",err)
		} else {
			log.Println(t.Execute(w, nil))
		}
	} else {
		r.ParseForm()
		fmt.Println(template.HTMLEscapeString(r.Form.Get("username")))


		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
	//if r.Method == "GET" {
	//	files, _ := template.ParseFiles("login.gtpl")
	//	log.Println(files.Execute(w, nil))
	//} else {
	//	r.ParseForm()
	//	username := r.Form["username"]
	//	fmt.Println(username)
	//	password := r.Form["password"]
	//	fmt.Println(password)
	//}
}

func login1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET"{
		current_time := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(current_time, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("/Users/finup123/IdeaProjects/goDemo/demo/login.gtpl")
		t.Execute(w, token)
	}
}
