package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}


func sayHello(w http.ResponseWriter,r *http.Request){
	//定义模板

	//解析模板
	t,err := template.ParseFiles("./hellp.tmpl")
	if err != nil{
		fmt.Println("parse temple failed, err:%v",err)
		return
	}
	u1 := User{
		Name: "qlf",
		Gender: "nain",
		Age: 128,
	}
	mi := map[string]interface{}{
		"name":"qdp",
		"gander":"nv",
		"age":"4",
	}

	hoboylist := []string{
		"wan",
		"chi",
		"shui",
	}
	//渲染模板
	t.Execute(w,map[string]interface{}{
		"u1":u1,
		"mi":mi,
		"hobbylist":hoboylist,
	})

}



func main() {
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil{
		fmt.Println("http server start is failed: %v",err)
		return
	}
}