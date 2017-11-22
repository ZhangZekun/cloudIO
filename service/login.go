package service

import (
	"fmt"
//	"html/template"
//	"log"
	"net/http"
	"github.com/unrolled/render"
)

/*
func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		fmt.Println("login get")
		http.FileServer(http.Dir("/home/os/gopath/src/webTest/cloudIO/assets/")).ServeHTTP(w,r)
		// t, err := template.ParseFiles("../assets/login.gtpl")
		// if err != nil {
		// 	fmt.Println("dangerous!")
		// 	return
		// }
		//log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

*/


func loginHandler(formatter *render.Render) http.HandlerFunc {
	
		return func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm();
			fmt.Println("method:", r.Method) //获取请求的方法
			if r.Method == "GET" {
				fmt.Println("login get")
				formatter.HTML(w, http.StatusOK, "login", "")
			}
			if r.Method == "POST" {
				//请求的是登录数据，那么执行登录的逻辑判断
				fmt.Println("username:", r.Form["username"])
				fmt.Println("password:", r.Form["password"])
				
				formatter.HTML(w, http.StatusOK, "table", struct {
					Username      string
					Username1	  string
					Phone1		  string
					Username2	  string
					Phone2		  string
				}{Username: r.Form.Get("username"), Username1: "zhangZekun", Phone1: "123456",
				Username2:"zhangBangtian", Phone2:"666666"})
			}
		}
}
