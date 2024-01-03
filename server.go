package main

import (
	"fmt"
	"net/http"
)

type TestWebServerType bool

func (m TestWebServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there !!!")
	fmt.Fprintf(w, "Request is: %+v", r)

}
func main() {
	var k TestWebServerType
	http.ListenAndServe("localhost:8080", k)
}
