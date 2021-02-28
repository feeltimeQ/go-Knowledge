package main

import (
	"fmt"
	"net/http"
)


//w 要返还给浏览器的内容响因  r 指针请求  （一个内容对应一个响应）
func sayHello(w http.ResponseWriter, r *http.Request){
	_,_ =fmt.Fprintln(w, "hello qlf")

}



func main(){
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		fmt.Println("failed")
		return
	}
}