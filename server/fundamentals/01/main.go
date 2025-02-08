package main

import (
	"net/http"
)

type api struct{
	addr string
}

func(s *api) ServeHTTP(w http.ResponseWriter,r *http.Request){
	// w.Write(([]byte("Hello from the server")))

	switch r.Method{
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write(([]byte("Hello from the server")))	
		case "/users":
			w.Write(([]byte("Hello from the users page")))
		case "/about":
			w.Write(([]byte("Hello from the about page")))
		}


	}
}

func main(){

	s:= &server{addr:":8080"}

	http.ListenAndServe(s.addr,s)
}