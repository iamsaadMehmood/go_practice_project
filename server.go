package main

import (
	"fmt"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
			Hye
		</head>
		<body>
			<h1>Saad Mehmood</h1>
			<a href="https://github.com/iamsaadMehmood">Github Link</a>
			<a href="https://www.linkedin.com/in/iamsaadmehmood/">LinkedIn Link</a>
			
		</body>
	</html>`)

}

func main() {

	http.ListenAndServe("localhost:8080", http.HandlerFunc(myFunc))
}
