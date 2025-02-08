package main

import (
	"fmt"
	"net/http"
)

func main(){
	server := &http.Server{Addr: ":80"}

	// Define routes for the server.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my homepage!")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the about page.")
	})
	server.Handler = mux

	// Start the server and listen for incoming requests.
	server.ListenAndServe()

}


// curl http://localhost:80/


// StatusCode        : 200
// StatusDescription : OK
// Content           : Welcome to my homepage!
// RawContent        : HTTP/1.1 200 OK
//                     Content-Length: 23
//                     Content-Type: text/plain; charset=utf-8
//                     Date: Fri, 07 Feb 2025 07:22:22 GMT

//                     Welcome to my homepage!
// Forms             : {}
// Headers           : {[Content-Length, 23], [Content-Type, text/plain; charset=utf-8], 
//                     [Date, Fri, 07 Feb 2025 07:22:22 GMT]}
// Images            : {}
// InputFields       : {}
// Links             : {}
// ParsedHtml        : System.__ComObject
// RawContentLength  : 23

// curl http://localhost:80/about


// StatusCode        : 200
// StatusDescription : OK
// Content           : This is the about page.
// RawContent        : HTTP/1.1 200 OK
//                     Content-Length: 23
//                     Content-Type: text/plain; charset=utf-8
//                     Date: Fri, 07 Feb 2025 07:23:01 GMT

//                     This is the about page.
// Forms             : {}
// Headers           : {[Content-Length, 23], [Content-Type, text/plain; charset=utf-8],    
//                     [Date, Fri, 07 Feb 2025 07:23:01 GMT]}
// Images            : {}
// InputFields       : {}
// Links             : {}
// ParsedHtml        : System.__ComObject
// RawContentLength  : 23
