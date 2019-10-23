package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := http.Cookie{
		Name:   "lucaAndTom",
		Value:  "thedroppedcookiehasgoldinit",
		// need to be turned on once we have https
		HttpOnly: true,
		// to be changed it depends on the user
		SameSite: http.SameSiteNoneMode,
		MaxAge: 3600}

	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)
	return
}


func main() {

	fmt.Println("Setting up server, enabling CORS . . .")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	router := mux.NewRouter()
	// Example handler
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)

	// Bind to port 8000 and pass our router in and pass the cors Handler
	log.Fatal(http.ListenAndServe(":8000", c.Handler(router)))
}
