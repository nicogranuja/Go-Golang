package main

import(
	"net/http"
	"io"

)

func main() {

	http.HandleFunc("/helloWorld", func (w http.ResponseWriter,r *http.Request ) {
		io.WriteString(w, "Hi")
	})
	http.HandleFunc("/", handler)

	//instead of null pass object with routing info
	http.ListenAndServe(":8000", nil)
}

//response writer is the structure for the uri and 
//request is the pointer that points to the information from the request
//the parameters are a signature for handling any web request
func handler(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w, "Hello World")		
}