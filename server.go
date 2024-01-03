package main

import (
	"fmt"
	"net/http"
)

type TestWebServerType bool

func (m TestWebServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
			Hye
		</head>
		<body>
			<h1>Saad Mehmood</h1>
			<a href="https://github.com/iamsaadMehmood">Github Link</a>
		</body>


	</html>`)

}
func main() {
	var k TestWebServerType
	http.ListenAndServe("localhost:8080", k)
}
