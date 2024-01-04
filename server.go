package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
			Hye
		</head>
		<body>
			<h1>Login</h1>
			
		</body>
	</html>`)

}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<body>
			<h1>Welcome</h1>
			
		</body>
	</html>`)

}
func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/welcome", welcome)
	// http.Handle("/loginHandle", http.HandlerFunc(login))
	// http.Handle("/welcomeHandle", http.HandlerFunc(welcome))
	fmt.Println("Listening to port 8080")
	http.ListenAndServe("localhost:8080", nil)
}
