package main

import "net/http"
import "database/sql"
import _ "modernc.org/sqlite"
import "fmt"



func main(){
	db,err:=sql.Open("sqlite","database.db")
	if err!=nil{
		panic(err)
	}
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		getsessid,err:=r.Cookie("sessid")
		if err!=nil{
			//mengecek apakah cookie masih ada atau expired
			if err==http.ErrNoCookie{
				w.Header().Set("Content-Type","text/html")
				w.Write([]byte("Session expired. Please click <a href='/login'>here</a>"))
			}
		}else{
			//fitur tambahan untuk mengecek sessid dengan database dengan membandingkan sessid
			getsessidcommand:=`SELECT id,username from user where sessid=?`
			var userid string
			var username string
			db.QueryRow(getsessidcommand,getsessid.Value).Scan(&userid,&username)
			if userid==""&&username==""{
				w.Header().Set("Content-Type","text/html")
				w.Write([]byte("Login failed. Please <a href='login'>login</a> again"))
			}else{
				//http.Redirect(w,r,"/home",http.StatusSeeOther)
				w.Header().Set("Content-Type","text/html")
				w.Write([]byte("<h1>Login Successful</h1><br><p>Login as "+username+"</p>"))
			}
		}
	})
	http.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"login.html")
		username:=r.FormValue("username")
		password:=r.FormValue("password")
		fmt.Println(username,password)
	})
	http.ListenAndServe(":1395",nil)
}