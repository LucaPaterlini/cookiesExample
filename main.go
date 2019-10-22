package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PROD = false
)


func SayHelloWorld(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "null")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	html := "Hello World! "
	w.Write([]byte(html))
}

func ReadCookie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "null")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Cache, Content-Type, Accept")
	c, err := r.Cookie("ithinkidroppedacookie")
	if err != nil {
		w.Write([]byte("error in reading cookie : "+err.Error() + "\n"))
	}else{
		value := c.Value
		w.Write([]byte("cookie has : " + value + "\n"))
	}
}

func DeleteCookie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "null")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Cache, Content-Type, Accept")
	c := http.Cookie{
		Name:   "ithinkidroppedacookie",
		MaxAge: -1,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &c)
	w.Write([]byte("old cookie deleted!\n"))
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "null")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Cache, Content-Type, Accept")

	fmt.Println(r.Header.Get("Origin"))
	c := http.Cookie{
		Name:   "ithinkidroppedacookie",
		Value:  "thedroppedcookiehasgoldinit",
		// need to be turned on once we have https
		Secure:   PROD,
		HttpOnly: true,
		// to be changed it depends on the user
		SameSite: http.SameSiteNoneMode,
		MaxAge: 3600}

	http.SetCookie(w, &c)
	fmt.Println("create",w.Header())

	w.Write([]byte("new cookie created!\n"))
}




func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHelloWorld)
	mux.HandleFunc("/readcookie", ReadCookie)
	mux.HandleFunc("/deletecookie", DeleteCookie)
	mux.HandleFunc("/createcookie", CreateCookie)
	switch PROD{
	case true:
		err := http.ListenAndServeTLS(":443",
			"/etc/letsencrypt/live/lucatesting.sportabletech.org/cert.pem",
			"/etc/letsencrypt/live/lucatesting.sportabletech.org/privkey.pem",
			mux)
		if err != nil {
			log.Fatal("HTTPS server: ", err)
		}
	default:
		err := http.ListenAndServe(":8080",mux)
		if err != nil {
			log.Fatal("HTTPS server: ", err)
		}

	}
}