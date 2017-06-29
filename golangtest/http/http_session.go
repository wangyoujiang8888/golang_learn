package main

import (
	"fmt"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	"html/template"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                    //解析url传递的参数，对于POST则解析响应包的主体（request body）
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		username := sess.Get("username")

		fmt.Println(username)
		t.Execute(w, sess.Get("username"))
	} else {
		r.ParseForm()
		fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	globalSessions.SessionDestroy(w, r)
}

func main() {

	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)   //设置访问的路由
	http.HandleFunc("/logout", logout) //设置访问的路由

	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

var globalSessions *session.Manager

func init() {
	var structdata session.ManagerConfig
	structdata.CookieName = "gosessionid"
	structdata.EnableSetCookie = true
	structdata.Gclifetime = 3600
	structdata.Maxlifetime = 3600
	structdata.CookieLifeTime = 3600
	structdata.ProviderConfig = "10.10.3.246:6379"
	globalSessions, _ = session.NewManager("redis", &structdata)
	go globalSessions.GC()
}
